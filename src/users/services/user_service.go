package services

import (
	"database/sql"
	"errors"
	"fmt"
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

func (c *userService) GetUser(username string) (*entities.Users, error) {
	user, e := c.userMysqlRepository.GetUser(username)
	if e != nil {
		return nil, e
	}
	if err := c.userScribleRepository.FindUserRefreshToken(fmt.Sprintf("%d", user.ID)); err != nil {
		return nil, err
	}
	return user, nil
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

func (c *userService) Logout(ctxId, username string) error {
	if err := c.userScribleRepository.FindUserRefreshToken(ctxId); err != nil {
		return err
	}

	user, err := c.userMysqlRepository.GetUser(username)
	if err != nil {
		return err
	}

	id := fmt.Sprintf("%d", user.ID)

	if id != ctxId {
		return errors.New("")
	}

	if err := c.userScribleRepository.DeleteUserRefreshToken(id); err != nil {
		return err
	}
	return nil
}

func (c *userService) UpdateAccount(payload *entities.Users, username string) (*entities.Users, error) {
	u, e := c.userMysqlRepository.GetUser(username)
	if e != nil {
		return nil, e
	}

	id := fmt.Sprintf("%d", u.ID)

	if err := c.userScribleRepository.FindUserRefreshToken(id); err != nil {
		return nil, err
	}

	payload.Password, _ = helper.Hash(payload.Password)
	user, err := c.userMysqlRepository.UpdateAccount(payload, u.ID)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (c *userService) UpdateAddress(payload *entities.UserAddress, username string) (*entities.UserAddress, error) {
	u, e := c.userMysqlRepository.GetUser(username)
	if e != nil {
		return nil, e
	}

	id := fmt.Sprintf("%d", u.ID)

	if err := c.userScribleRepository.FindUserRefreshToken(id); err != nil {
		return nil, err
	}
	
	res, err := c.userMysqlRepository.UpdateAddress(payload, u.ID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *userService) GetAddress(username string) (*entities.UserAddress, error)  {
	u, e := c.userMysqlRepository.GetUser(username)
	if e != nil {
		return nil, e
	}

	id := fmt.Sprintf("%d", u.ID)

	if err := c.userScribleRepository.FindUserRefreshToken(id); err != nil {
		return nil, err
	}
	address, err := c.userMysqlRepository.GetAddress(u.ID)
	return address, err
}