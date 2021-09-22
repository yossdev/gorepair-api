package middlewares

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

type JwtClaims struct {
	UserId  int 	`json:"id"`
	Name 	string  `json:"name"`
	jwt.StandardClaims
}

func CreateToken(userId int, name string) (string, error) {
	// Set custom claims
	claims := &JwtClaims{
		userId,
		name,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	jwtToken, err := token.SignedString([]byte(viper.GetString(`SECRET_JWT`)))
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func Restricted(c echo.Context) error {
	var claims JwtClaims
	user := c.Get("user").(*jwt.Token)
	tmp, _ := json.Marshal(user.Claims)
	_ = json.Unmarshal(tmp, &claims)
	name := claims.Name
	return c.String(http.StatusOK, "Welcome "+name+"!")
}
