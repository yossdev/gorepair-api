package handlers

import (
	"gorepair-rest-api/internal/utils/helper"
	"gorepair-rest-api/internal/web"
	"gorepair-rest-api/src/users/dto"
	"gorepair-rest-api/src/users/entities"
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
		return web.JsonResponse(ctx, http.StatusBadRequest, web.BadRequest, nil)
	}

	res, err := service.UserService.Login(payload.ToDomain())
	if err != nil {
		return web.JsonResponse(ctx, http.StatusUnauthorized, web.UsernamePasswordWrong, nil)
	}

	return web.JsonResponse(ctx, http.StatusOK, web.Welcome, res)
}

func (service *userHandlers) Logout(ctx *fiber.Ctx) error {
	// err := service.UserService.FindByID(ctx.Get("id"))
	// if err != nil {
	// 	return web.JsonResponse(ctx, http.StatusForbidden, web.Forbidden, nil)
	// }

	// user, err := service.UserService.GetUser(ctx.Params("username"))
	// if err != nil {
	// 	return web.JsonResponse(ctx, http.StatusForbidden, web.Forbidden, nil)
	// }

	e := service.UserService.Logout(ctx.Get("id"), ctx.Params("username"))
	if e != nil {
		return web.JsonResponse(ctx, http.StatusForbidden, web.Forbidden, nil)
	}

	return web.JsonResponse(ctx, http.StatusOK, web.SuccessLogOut, nil)
}

func (service *userHandlers) Register(ctx *fiber.Ctx) error {
	payload := new(dto.UserRequestRegisterBody)
	if err := ctx.BodyParser(payload); err != nil {
		return web.JsonResponse(ctx, http.StatusBadRequest, web.BadRequest, nil)
	}

	if ok, _ := helper.ValidateInputs(*payload); !ok {
		return web.JsonResponse(ctx, http.StatusBadRequest, web.CannotEmpty, nil)
	}

	user, err := service.UserService.Register(payload.ToDomain(), payload.Street)
	if err != nil {
		return web.JsonResponse(ctx, http.StatusInternalServerError, web.UserExist, nil)
	}

	return web.JsonResponse(ctx, http.StatusCreated, web.AccountCreated, dto.FromDomain(user))
}

func (service *userHandlers) GetUser(ctx *fiber.Ctx) error {
	user, err := service.UserService.GetUser(ctx.Params("username"))
	if err != nil {
		return web.JsonResponse(ctx, http.StatusForbidden, web.Forbidden, nil)
	}

	// ok := service.UserService.FindByID(fmt.Sprintf("%d", user.ID))
	// if ok != nil {
	// 	return web.JsonResponse(ctx, http.StatusForbidden, web.Forbidden, nil)
	// }

	return web.JsonResponse(ctx, http.StatusOK, web.Success, dto.FromDomain(user))
}

func (service *userHandlers) UpdateAccount(ctx *fiber.Ctx) error {
	// rec, err := service.UserService.GetUser(ctx.Params("username"))
	// if err != nil {
	// 	return web.JsonResponse(ctx, http.StatusForbidden, web.Forbidden, nil)
	// }

	// ok := service.UserService.FindByID(fmt.Sprintf("%d", rec.ID))
	// if ok != nil {
	// 	return web.JsonResponse(ctx, http.StatusForbidden, web.Forbidden, nil)
	// }

	account := new(dto.UserAccountUpdateBody)
	e := ctx.BodyParser(account)
	if e != nil {
		return web.JsonResponse(ctx, http.StatusBadRequest, web.BadRequest, nil)
	}

	if ok, _ := helper.ValidateInputs(*account); !ok {
		return web.JsonResponse(ctx, http.StatusBadRequest, web.CannotEmpty, nil)
	}

	res, err := service.UserService.UpdateAccount(account.ToDomain(), ctx.Params("username"))
	if err != nil {
		return web.JsonResponse(ctx, http.StatusForbidden, web.Forbidden, nil)
	}

	return web.JsonResponse(ctx, http.StatusOK, web.UpdateSuccess, dto.FromDomainUpdate(res))
}

func (service *userHandlers) UpdateAddress(ctx *fiber.Ctx) error {
	// rec, err := service.UserService.GetUser(ctx.Params("username"))
	// if err != nil {
	// 	return web.JsonResponse(ctx, http.StatusForbidden, web.Forbidden, nil)
	// }

	// ok := service.UserService.FindByID(fmt.Sprintf("%d", rec.ID))
	// if ok != nil {
	// 	return web.JsonResponse(ctx, http.StatusForbidden, web.Forbidden, nil)
	// }

	address := new(dto.UserAddressUpdateBody)
	e := ctx.BodyParser(address)
	if e != nil {
		return web.JsonResponse(ctx, http.StatusBadRequest, web.BadRequest, nil)
	}

	if ok, _ := helper.ValidateInputs(*address); !ok {
		return web.JsonResponse(ctx, http.StatusBadRequest, web.CannotEmpty, nil)
	}

	res, err := service.UserService.UpdateAddress(address.ToDomain(), ctx.Params("username"))
	if err != nil {
		return web.JsonResponse(ctx, http.StatusForbidden, web.Forbidden, nil)
	}

	return web.JsonResponse(ctx, http.StatusOK, web.UpdateSuccess, dto.FromDomainAddress(res)) //TODO
}

func (service *userHandlers) GetAddress(ctx *fiber.Ctx) error {
	// user, err := service.UserService.GetUser(ctx.Params("username"))
	// if err != nil {
	// 	return web.JsonResponse(ctx, http.StatusOK, web.UserNotExist, nil)
	// }

	// ok := service.UserService.FindByID(fmt.Sprintf("%d", user.ID))
	// if ok != nil {
	// 	return web.JsonResponse(ctx, http.StatusForbidden, web.Forbidden, nil)
	// }

	address, err := service.UserService.GetAddress(ctx.Params("username"))
	if err != nil {
		return web.JsonResponse(ctx, http.StatusForbidden, web.Forbidden, nil)
	}
	return web.JsonResponse(ctx, http.StatusOK, web.Success, dto.FromDomainAddress(address))
}