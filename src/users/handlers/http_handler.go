package handlers

import (
	"fmt"
	"gorepair-rest-api/internal/web"
	"gorepair-rest-api/src/users/dto"
	"gorepair-rest-api/src/users/services"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type UserHandlers interface {
	Register(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
	Refresh(ctx *fiber.Ctx) error
	GetUser(ctx *fiber.Ctx) error
}

type userHandlers struct {
	UserService services.UserService
}

func NewHttpHandler(
	userService services.UserService,
) UserHandlers {
	return &userHandlers{
		UserService: userService,
	}
}

func (service *userHandlers) Register(ctx *fiber.Ctx) error {
	userData := new(dto.UserRequestRegisterBody)

	if err := ctx.BodyParser(userData); err != nil {
		log.Fatal(err)
	}

	if userData.Username == "" || userData.Name == "" || userData.Email == "" || userData.Password == "" || userData.Phone == "" {
		return web.JsonResponse(ctx, http.StatusBadRequest,  "Bad Request", nil)
	}

	user, err := service.UserService.Register(*userData)
	if err != nil {
		return web.JsonResponse(ctx, http.StatusInternalServerError, "User already exist", nil)
	}

	var res dto.UserResponseBody

	res.ID = fmt.Sprint(user.ID)
	res.Username = user.Username
	res.Name = user.Name
	res.Email = user.Email
	res.Phone = user.Phone

	return web.JsonResponse(ctx, http.StatusCreated, "Account Created", res)
}

func (service *userHandlers) Login(ctx *fiber.Ctx) error {
	userData := new(dto.UserRequestLoginBody)

	if err := ctx.BodyParser(userData); err != nil {
		log.Fatal(err)
	}

	res, err := service.UserService.Login(userData)
	if err != nil {
		return web.JsonResponse(ctx, http.StatusUnauthorized, err.Error(), nil)
	}

	return web.JsonResponse(ctx, http.StatusOK, "", res)
}

func (service *userHandlers) GetUser(ctx *fiber.Ctx) error {
	user, err := service.UserService.GetUser(ctx.Params("username"))
	if err != nil {
		return web.JsonResponse(ctx, http.StatusNotFound, "User is not exist", nil)
	}

	return web.JsonResponse(ctx, http.StatusOK, "Success", user)
}

func (service *userHandlers) Refresh(ctx *fiber.Ctx) error {
	userID := ctx.Get("userID")

	res, err := service.UserService.RefreshToken(userID)

	if err != nil {
		return web.JsonResponse(ctx, http.StatusUnauthorized, err.Error(), nil)
	}

	return web.JsonResponse(ctx, http.StatusOK, "", res)
}
