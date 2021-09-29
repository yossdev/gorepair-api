package handlers

import (
	"gorepair-rest-api/internal/utils/helper"
	"gorepair-rest-api/internal/web"
	"gorepair-rest-api/src/users/dto"
	"gorepair-rest-api/src/users/entities"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type UserHandlers interface {
	GetUser(ctx *fiber.Ctx) error
	Register(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
	// Logout(ctx *fiber.Ctx) error
	// Account(ctx *fiber.Ctx) error
	// Address(ctx *fiber.Ctx) error
}

type userHandlers struct {
	UserService entities.UserService
}

func NewHttpHandler(userService entities.UserService) UserHandlers {
	return &userHandlers{
		UserService: userService,
	}
}

func (service *userHandlers) GetUser(ctx *fiber.Ctx) error {
	user, err := service.UserService.GetUser(ctx.Params("username"))
	if err != nil {
		return web.JsonResponse(ctx, http.StatusOK, "User is not exist", nil)
	}

	return web.JsonResponse(ctx, http.StatusOK, "Success", dto.FromDomain(user))
}

func (service *userHandlers) Register(ctx *fiber.Ctx) error {
	payload := new(dto.UserRequestRegisterBody)
	if err := ctx.BodyParser(payload); err != nil {
		return web.JsonResponse(ctx, http.StatusBadRequest, "something is not right", nil)
	}

	if ok, _ := helper.ValidateInputs(*payload); !ok {
		return web.JsonResponse(ctx, http.StatusBadRequest, "field cannot be empty", nil)
	}

	user, err := service.UserService.Register(payload.ToDomain())
	if err != nil {
		return web.JsonResponse(ctx, http.StatusInternalServerError, "User already exist", nil)
	}

	return web.JsonResponse(ctx, http.StatusCreated, "Account Created", dto.FromDomain(user))
}

func (service *userHandlers) Login(ctx *fiber.Ctx) error {
	payload := new(dto.UserRequestLoginBody)
	if err := ctx.BodyParser(payload); err != nil {
		log.Fatal(err)
	}

	res, err := service.UserService.Login(payload.ToDomain())
	if err != nil {
		return web.JsonResponse(ctx, http.StatusUnauthorized, "email or password is wrong!", nil)
	}

	return web.JsonResponse(ctx, http.StatusOK, "welcome!", res)
}

// func (service *userHandlers) Logout(ctx *fiber.Ctx) error {

// }

// func (service *userHandlers) Account(ctx *fiber.Ctx) error {

// }

// func (service *userHandlers) Address(ctx *fiber.Ctx) error {

// }