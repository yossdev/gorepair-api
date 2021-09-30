package auth

import (
	"fmt"
	"gorepair-rest-api/config"
	"gorepair-rest-api/infrastructures/local_db"
	"gorepair-rest-api/infrastructures/logger"
	"gorepair-rest-api/internal/utils/encryption"
	"time"

	"github.com/golang-jwt/jwt"
)

type TokenStruct struct {
	Type         string `json:"type"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenStruct struct {
	RefreshToken string `json:"refresh_token"`
	Expired      int64  `json:"expired"`
}

type JwtTokenInterface interface {
	Sign(claims jwt.MapClaims) TokenStruct
}

type jwtToken struct {
	cached local_db.ScribleDB
}

func NewJwt(cached local_db.ScribleDB) JwtTokenInterface {
	return &jwtToken{cached: cached}
}

func (o jwtToken) Sign(claims jwt.MapClaims) TokenStruct {
	timeNow := time.Now()
	tokenExpired := timeNow.Add(config.Get().JwtTokenExpired).Unix()

	if claims["id"] == nil {
		return TokenStruct{}
	}

	token := jwt.New(jwt.SigningMethodHS256)
	// setup userdata
	var _, checkExp = claims["exp"]
	var _, checkIat = claims["exp"]

	// if user didn't define claims expired
	if !checkExp {
		claims["exp"] = tokenExpired
	}

	// if user didn't define claims iat
	if !checkIat {
		claims["iat"] = timeNow.Unix()
	}
	claims["token_type"] = "access_token"

	token.Claims = claims

	authToken := new(TokenStruct)
	tokenString, err := token.SignedString([]byte(config.Get().AppKey))

	if err != nil {
		logger.Log.Errorln(err)
		return TokenStruct{}
	}

	authToken.Token = tokenString
	authToken.Type = config.Get().JwtTokenType

	//create refresh token
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	refreshTokenExpired := timeNow.Add(config.Get().JwtRefreshExpired).Unix()

	claims["exp"] = refreshTokenExpired
	claims["token_type"] = "refresh_token"
	refreshToken.Claims = claims

	refreshTokenString, err := refreshToken.SignedString([]byte(config.Get().AppKey))

	if err != nil {
		return TokenStruct{}
	}
	authToken.RefreshToken = refreshTokenString

	//save token to scrible db
	go func() {
		encryptedRefreshToken, err := encryption.AesCFBEncryption(refreshTokenString, config.Get().AppKey)
		if err != nil {
			logger.Log.Errorln(err)
		}
		// check data type of the claims
		switch claims["id"].(type) {
		case uint64:
			claims["id"] = fmt.Sprintf("%d", claims["id"].(uint64))
		case float64:
			claims["id"] = fmt.Sprintf("%d", int(claims["id"].(float64)))
		default:
		}
		if claims["role"].(string) == "user" {
			o.cached.DB().Write("refresh_token_user", claims["id"].(string), RefreshTokenStruct{RefreshToken: encryptedRefreshToken, Expired: refreshTokenExpired})
		} else {
			o.cached.DB().Write("refresh_token_workshop", claims["id"].(string), RefreshTokenStruct{RefreshToken: encryptedRefreshToken, Expired: refreshTokenExpired})
		}
		if err != nil {
			logger.Log.Infoln("Failed to save refresh token to scrible, with err: ", err)
		} else {
			logger.Log.Infoln("Successfully to save refresh token to scrible")
		}
	}()

	return TokenStruct{
		Type:         "Bearer",
		Token:        authToken.Token,
		RefreshToken: authToken.RefreshToken,
	}
}