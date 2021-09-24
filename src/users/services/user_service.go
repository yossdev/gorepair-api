package services

import (
	"database/sql"
	"gorepair-rest-api/helper"
	"gorepair-rest-api/internal/utils/auth"
	"gorepair-rest-api/src/users/dto"
	"gorepair-rest-api/src/users/entities"
	"gorepair-rest-api/src/users/repositories"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	FindByID(id uint) entities.User
	Register(data dto.UserRequestRegisterBody) (*entities.User, error)
	Login(data *dto.UserRequestLoginBody) (dto.UserTokenResponseBody, error)
	RefreshToken(userID string) (dto.UserTokenResponseBody, error)
}

type userService struct {
	userMysqlRepository   repositories.UserMysqlRepositoryInterface
	userScribleRepository repositories.UserScribleRepositoryInterface
	jwtAuth               auth.JwtTokenInterface
}

func NewUserService(
	userMysqlRepository repositories.UserMysqlRepositoryInterface,
	jwtAuth auth.JwtTokenInterface,
	userScribleRepository repositories.UserScribleRepositoryInterface,
) UserService {
	return &userService{
		userScribleRepository: userScribleRepository,
		userMysqlRepository:   userMysqlRepository,
		jwtAuth:               jwtAuth,
	}
}

func (c *userService) FindByID(id uint) entities.User {
	return c.userMysqlRepository.FindByID(id)
}

func (c *userService) Register(data dto.UserRequestRegisterBody) (*entities.User, error) {
	password, _ := helper.Hash(data.Password)

	user, err := c.userMysqlRepository.Register(data.Username, data.Name, data.Email, password, data.Phone)
	return user, err
}

func (c *userService) Login(data *dto.UserRequestLoginBody) (dto.UserTokenResponseBody, error) {

	user := c.userMysqlRepository.FindByEmail(data.Email)
	if user.ID == 0 {
		return dto.UserTokenResponseBody{}, sql.ErrNoRows
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))

	if err != nil {
		return dto.UserTokenResponseBody{}, err
	}

	userToken := c.jwtAuth.Sign(jwt.MapClaims{
		"id": user.ID,
	})

	token := dto.UserTokenResponseBody(userToken)

	return token, nil
}

func (c *userService) RefreshToken(userID string) (dto.UserTokenResponseBody, error) {
	refreshToken, err := c.userScribleRepository.FindUserRefreshToken(userID)
	if err != nil {
		return dto.UserTokenResponseBody{}, err
	}

	if refreshToken.Expired < time.Now().Unix() {
		return dto.UserTokenResponseBody{}, err
	}

	userToken := c.jwtAuth.Sign(jwt.MapClaims{
		"id": userID,
	})

	token := dto.UserTokenResponseBody(userToken)

	return token, nil
}
