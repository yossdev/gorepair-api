package services

import (
	"database/sql"
	"errors"
	"gorepair-rest-api/internal/utils/auth"
	"gorepair-rest-api/internal/utils/helper"
	"gorepair-rest-api/src/users/entities"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userMysqlRepository   entities.UserMysqlRepositoryInterface
	userScribleRepository entities.UserScribleRepositoryInterface
	jwtAuth               auth.JwtTokenInterface
}

func NewUserService(
	userMysqlRepository 	entities.UserMysqlRepositoryInterface,	
	userScribleRepository 	entities.UserScribleRepositoryInterface,
	jwtAuth 				auth.JwtTokenInterface,
) entities.UserService {
	return &userService{
		userMysqlRepository:   	userMysqlRepository,
		userScribleRepository:  userScribleRepository,
		jwtAuth: 				jwtAuth,
		
	}
}

func (c *userService) FindByID(id string) error {
	err := c.userScribleRepository.FindUserRefreshToken(id)
	if err != nil {
		return err
	}
	return nil
}

func (c *userService) GetUser(username string) (*entities.Users, error) {
	user, err := c.userMysqlRepository.GetUser(username)
	return user, err
}

func (c *userService) Register(payload *entities.Users , street string) (*entities.Users, error) {
	payload.Password, _ = helper.Hash(payload.Password)
	user, err := c.userMysqlRepository.Register(payload, street)
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

	token := c.jwtAuth.Sign(jwt.MapClaims{
		"id": user.ID,
		"role": "user",
	})

	return token, nil
}

func (c *userService) Logout(id, ctxId string) error {
	if id != ctxId {
		return errors.New("")
	}

	err := c.userScribleRepository.DeleteUserRefreshToken(id)
	if err != nil {
		return err
	}
	return nil
}

func (c *userService) UpdateAccount(payload *entities.Users, id uint64) (*entities.Users, error) {
	payload.Password, _ = helper.Hash(payload.Password)
	user, err := c.userMysqlRepository.UpdateAccount(payload, id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (c *userService) UpdateAddress(payload *entities.UserAddress, id uint64) (*entities.UserAddress, error) {
	res, err := c.userMysqlRepository.UpdateAddress(payload, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *userService) GetAddress(id uint64) (*entities.UserAddress, error)  {
	address, err := c.userMysqlRepository.GetAddress(id)
	return address, err
}