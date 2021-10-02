package handlers

import (
	"fmt"
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
	Logout(ctx *fiber.Ctx) error
	UpdateAccount(ctx *fiber.Ctx) error
	UpdateAddress(ctx *fiber.Ctx) error
	GetAddress(ctx *fiber.Ctx) error
}

type userHandlers struct {
	UserService entities.UserService
}

func NewHttpHandler(userService entities.UserService) UserHandlers {
	return &userHandlers{
		UserService: userService,
	}
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

func (service *userHandlers) Logout(ctx *fiber.Ctx) error {
	err := service.UserService.FindByID(ctx.Get("id"))
	if err != nil {
		return web.JsonResponse(ctx, http.StatusForbidden, "forbidden", nil)
	}

	user, err := service.UserService.GetUser(ctx.Params("username"))
	if err != nil {
		return web.JsonResponse(ctx, http.StatusForbidden, "forbidden", nil)
	}

	e := service.UserService.Logout(fmt.Sprintf("%d", user.ID))
	if e != nil {
		return web.JsonResponse(ctx, http.StatusForbidden, "forbidden", nil)
	}

	return web.JsonResponse(ctx, http.StatusOK, "successfully logged out", nil)
}

func (service *userHandlers) Register(ctx *fiber.Ctx) error {
	payload := new(dto.UserRequestRegisterBody)
	if err := ctx.BodyParser(payload); err != nil {
		return web.JsonResponse(ctx, http.StatusBadRequest, "something is not right", nil)
	}

	if ok, _ := helper.ValidateInputs(*payload); !ok {
		return web.JsonResponse(ctx, http.StatusBadRequest, "field cannot be empty", nil)
	}

	user, err := service.UserService.Register(payload.ToDomain(), payload.Street)
	if err != nil {
		return web.JsonResponse(ctx, http.StatusInternalServerError, "user already exist", nil)
	}

	return web.JsonResponse(ctx, http.StatusCreated, "account created!", dto.FromDomain(user))
}

func (service *userHandlers) GetUser(ctx *fiber.Ctx) error {
	err := service.UserService.FindByID(ctx.Get("id"))
	if err != nil {
		return web.JsonResponse(ctx, http.StatusForbidden, "forbidden", nil)
	}

	user, err := service.UserService.GetUser(ctx.Params("username"))
	if err != nil {
		return web.JsonResponse(ctx, http.StatusOK, "user is not exist", nil)
	}

	return web.JsonResponse(ctx, http.StatusOK, "success", dto.FromDomain(user))
}

func (service *userHandlers) UpdateAccount(ctx *fiber.Ctx) error {
	rec, err := service.UserService.GetUser(ctx.Params("username"))
	if err != nil {
		return web.JsonResponse(ctx, http.StatusForbidden, "forbidden", nil)
	}

	ok := service.UserService.FindByID(fmt.Sprintf("%d", rec.ID))
	if ok != nil {
		return web.JsonResponse(ctx, http.StatusForbidden, "forbidden", nil)
	}

	account := new(dto.UserAccountUpdateBody)
	e := ctx.BodyParser(account)
	if e != nil {
		return web.JsonResponse(ctx, http.StatusBadRequest, "something is not right with your request", nil)
	}

	if ok, _ := helper.ValidateInputs(*account); !ok {
		return web.JsonResponse(ctx, http.StatusBadRequest, "field cannot be empty", nil)
	}

	res, err := service.UserService.UpdateAccount(account.ToDomain(), rec.ID)
	if err != nil {
		return web.JsonResponse(ctx, http.StatusBadRequest, "problem with db", nil)
	}

	return web.JsonResponse(ctx, http.StatusOK, "successfully updated!", dto.FromDomainUpdate(res))
}

func (service *userHandlers) UpdateAddress(ctx *fiber.Ctx) error {
	rec, err := service.UserService.GetUser(ctx.Params("username"))
	if err != nil {
		return web.JsonResponse(ctx, http.StatusForbidden, "forbidden", nil)
	}

	ok := service.UserService.FindByID(fmt.Sprintf("%d", rec.ID))
	if ok != nil {
		return web.JsonResponse(ctx, http.StatusForbidden, "forbidden", nil)
	}

	address := new(dto.UserAddressUpdateBody)
	e := ctx.BodyParser(address)
	if e != nil {
		return web.JsonResponse(ctx, http.StatusBadRequest, "something is not right with your request", nil)
	}

	if ok, _ := helper.ValidateInputs(*address); !ok {
		return web.JsonResponse(ctx, http.StatusBadRequest, "field cannot be empty", nil)
	}

	res, err := service.UserService.UpdateAddress(address.ToDomain(), rec.ID)
	if err != nil {
		return web.JsonResponse(ctx, http.StatusBadRequest, "problem with db", nil)
	}

	return web.JsonResponse(ctx, http.StatusOK, "successfully updated!", dto.FromDomainAddress(res)) //TODO
}

func (service *userHandlers) GetAddress(ctx *fiber.Ctx) error {
	user, err := service.UserService.GetUser(ctx.Params("username"))
	if err != nil {
		return web.JsonResponse(ctx, http.StatusOK, "user is not exist", nil)
	}

	ok := service.UserService.FindByID(fmt.Sprintf("%d", user.ID))
	if ok != nil {
		return web.JsonResponse(ctx, http.StatusForbidden, "forbidden", nil)
	}

	address, _ := service.UserService.GetAddress(user.ID)

	return web.JsonResponse(ctx, http.StatusOK, "success", dto.FromDomainAddress(address))
}