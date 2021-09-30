package services

import (
	"database/sql"
	"errors"
	"gorepair-rest-api/internal/utils/auth"
	"gorepair-rest-api/internal/utils/helper"
	"gorepair-rest-api/src/users/entities"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userMysqlRepository   entities.UserRepository
	userScribleRepository entities.UserScribleRepositoryInterface
	jwtAuth               auth.JwtTokenInterface
}

func NewUserService(
	userMysqlRepository entities.UserRepository,	
	jwtAuth auth.JwtTokenInterface,
	userScribleRepository entities.UserScribleRepositoryInterface,
) entities.UserService {
	return &userService{
		userMysqlRepository:   userMysqlRepository,
		userScribleRepository:  userScribleRepository,
		jwtAuth:               jwtAuth,
		
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

	token := c.jwtAuth.Sign(jwt.MapClaims{
		"id": user.ID,
		"role": "user",
	})

	return token, nil
}

func (c *userService) Logout(ctx *fiber.Ctx, id string) error {
	role := helper.Restricted(ctx)
	if id == ctx.Get("id") && role == "user" {
		err := c.userScribleRepository.DeleteUserRefreshToken(id)
		if err != nil {
			return err
		}
		return nil
	}

	return errors.New("unauthorized")
}

func (c *userService) UpdateAccount(payload *entities.Users, id string) (*entities.Users, error) {
	payload.Password, _ = helper.Hash(payload.Password)
	payload.UpdatedAt = time.Now()
	user, err := c.userMysqlRepository.UpdateAccount(payload, id)
	if err != nil {
		return user, err
	}
	return user, nil
}