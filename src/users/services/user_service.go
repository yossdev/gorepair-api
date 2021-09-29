package services

import (
	"database/sql"
	"gorepair-rest-api/infrastructures/third-party/freegeoapi"
	"gorepair-rest-api/internal/utils/auth"
	"gorepair-rest-api/internal/utils/helper"
	"gorepair-rest-api/src/users/entities"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userMysqlRepository   entities.UserRepository
	jwtAuth               auth.JwtTokenInterface
}

func NewUserService(
	userMysqlRepository entities.UserRepository,
	jwtAuth auth.JwtTokenInterface,
) entities.UserService {
	return &userService{
		userMysqlRepository:   userMysqlRepository,
		jwtAuth:               jwtAuth,
	}
}

func (c *userService) GetUser(username string) (*entities.Users, error) {
	user, err := c.userMysqlRepository.GetUser(username)
	return user, err
}

func (c *userService) Register(payload *entities.Users) (*entities.Users, error) {
	payload.Password, _ = helper.Hash(payload.Password)
	user, err := c.userMysqlRepository.Register(payload)
	return user, err
}

func (c *userService) Login(payload *entities.Users) (interface{}, error) {
	user := c.userMysqlRepository.FindByEmail(payload.Email)
	if user.ID == 0 {
		return nil, sql.ErrNoRows
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		return nil, err
	}

	loc, _ := freegeoapi.NewIpAPI().GetLocationByIP()

	token := c.jwtAuth.Sign(jwt.MapClaims{
		"id": user.ID,
		"cty": loc.City,
	})

	return token, nil
}

// func (c *userService) Logout() error {
	
// }

// func (c *userService) Account() error {
	
// }

// func (c *userService) Address() error {
	
// }