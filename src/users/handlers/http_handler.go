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
	Register(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
	GetUser(ctx *fiber.Ctx) error
	Refresh(ctx *fiber.Ctx) error
}

type userHandlers struct {
	UserService entities.Service
}

func NewHttpHandler(userService entities.Service) UserHandlers {
	return &userHandlers{
		UserService: userService,
	}
}

func (service *userHandlers) Register(ctx *fiber.Ctx) error {
	userData := new(dto.UserRequestRegisterBody)
	if err := ctx.BodyParser(userData); err != nil {
		return web.JsonResponse(ctx, http.StatusBadRequest, "something is not right", nil)
	}

	if ok, _ := helper.ValidateInputs(*userData); !ok {
		return web.JsonResponse(ctx, http.StatusBadRequest, "field cannot be empty", nil)
	}

	user, err := service.UserService.Register(userData.ToDomain())

	if err != nil {
		return web.JsonResponse(ctx, http.StatusInternalServerError, "User already exist", nil)
	}

	return web.JsonResponse(ctx, http.StatusCreated, "Account Created", dto.FromDomain(*user))
}

func (service *userHandlers) Login(ctx *fiber.Ctx) error {
	userData := new(dto.UserRequestLoginBody)
	if err := ctx.BodyParser(userData); err != nil {
		log.Fatal(err)
	}

	res, err := service.UserService.Login(userData.ToDomain())
	if err != nil {
		return web.JsonResponse(ctx, http.StatusUnauthorized, "email or password is wrong!", nil)
	}

	return web.JsonResponse(ctx, http.StatusOK, "welcome!", dto.FromAuth(res))
}

func (service *userHandlers) GetUser(ctx *fiber.Ctx) error {
	user, err := service.UserService.GetUser(ctx.Params("username"))
	if err != nil {
		return web.JsonResponse(ctx, http.StatusOK, "User is not exist", nil)
	}

	return web.JsonResponse(ctx, http.StatusOK, "Success", dto.FromDomain(*user))
}

func (service *userHandlers) Refresh(ctx *fiber.Ctx) error {
	id := ctx.Get("id")

	res, err := service.UserService.RefreshToken(id)
	if err != nil {
		return web.JsonResponse(ctx, http.StatusUnauthorized, err.Error(), nil)
	}

	return web.JsonResponse(ctx, http.StatusOK, "welcome back!", dto.FromAuth(res))
}
