package services

import (
	"database/sql"
	"errors"
	"gorepair-rest-api/infrastructures/third-party/freegeoapi"
	realemailapi "gorepair-rest-api/infrastructures/third-party/real-email-api"
	"gorepair-rest-api/internal/utils/auth"
	"gorepair-rest-api/internal/utils/helper"
	"gorepair-rest-api/src/users/entities"
	"gorepair-rest-api/src/users/repositories"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userMysqlRepository   entities.Repository
	userScribleRepository repositories.UserScribleRepositoryInterface
	jwtAuth               auth.JwtTokenInterface
}

func NewUserService(
	userMysqlRepository entities.Repository,
	jwtAuth auth.JwtTokenInterface,
	userScribleRepository repositories.UserScribleRepositoryInterface,
) entities.Service {
	return &userService{
		userScribleRepository: userScribleRepository,
		userMysqlRepository:   userMysqlRepository,
		jwtAuth:               jwtAuth,
	}
}

func (c *userService) Register(data *entities.Users) (*entities.Users, error) {
	_, err := strconv.Atoi(data.Phone)
	if err != nil {
		return nil, err
	}
	
	data.Password, _ = helper.Hash(data.Password)
	user, err := c.userMysqlRepository.Register(data)
	return user, err
}

func (c *userService) Login(data *entities.Users) (auth.TokenStruct, error) {
	// will check the real-world email from api
	real := realemailapi.RealEmail(data.Email)
	if real.Status == "invalid" {
		return auth.TokenStruct{}, errors.New("invalid email")
	}

	user := c.userMysqlRepository.FindByEmail(data.Email)
	if user.ID == 0 {
		return auth.TokenStruct{}, sql.ErrNoRows
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if err != nil {
		return auth.TokenStruct{}, err
	}

	loc, _ := freegeoapi.NewIpAPI().GetLocationByIP()

	userToken := c.jwtAuth.Sign(jwt.MapClaims{
		"id": user.ID,
		"cty": loc.City,
	})

	return userToken, nil
}

func (c *userService) GetUser(username string) (*entities.Users, error) {
	user, err := c.userMysqlRepository.GetUser(username)
	return user, err
}

func (c *userService) FindByID(id uint64) (*entities.Users, error) {
	res, err := c.userMysqlRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *userService) RefreshToken(id string) (auth.TokenStruct, error) {
	refreshToken, err := c.userScribleRepository.FindUserRefreshToken(id)
	if err != nil {
		return auth.TokenStruct{}, err
	}

	if refreshToken.Expired < time.Now().Unix() {
		return auth.TokenStruct{}, err
	}

	loc, _ := freegeoapi.NewIpAPI().GetLocationByIP()

	userToken := c.jwtAuth.Sign(jwt.MapClaims{
		"id": id,
		"cty": loc.City,
	})

	return userToken, nil
}
