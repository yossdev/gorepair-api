package handlers

import (
	"gorepair-rest-api/internal/utils/helper"
	"gorepair-rest-api/internal/web"
	"gorepair-rest-api/src/workshops/dto"
	"gorepair-rest-api/src/workshops/entities"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type WorkshopHandlers interface {
	GetWorkshop(ctx *fiber.Ctx) error
	Register(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
	Logout(ctx *fiber.Ctx) error
	UpdateAccount(ctx *fiber.Ctx) error
	UpdateAddress(ctx *fiber.Ctx) error
	GetAddress(ctx *fiber.Ctx) error
	UpdateDescription(ctx *fiber.Ctx) error
	ServicesNew(ctx *fiber.Ctx) error
	UpdateServices(ctx *fiber.Ctx) error
	DeleteServices(ctx *fiber.Ctx) error
}

type workshopHandlers struct {
	WorkshopService entities.WorkshopService
}

func NewHttpHandler(workshopService entities.WorkshopService) WorkshopHandlers {
	return &workshopHandlers{
		WorkshopService: workshopService,
	}
}

func (service *workshopHandlers) Login(ctx *fiber.Ctx) error {
	payload := new(dto.WorkshopRequestLoginBody)
	if err := ctx.BodyParser(payload); err != nil {
		return web.JsonResponse(ctx, http.StatusBadRequest, web.BadRequest, nil)
	}

	res, err := service.WorkshopService.Login(payload.ToDomain())
	if err != nil {
		return web.JsonResponse(ctx, http.StatusUnauthorized, web.UsernamePasswordWrong, nil)
	}

	return web.JsonResponse(ctx, http.StatusOK, web.Welcome, res)
}

func (service *workshopHandlers) Logout(ctx *fiber.Ctx) error {
	e := service.WorkshopService.Logout(ctx.Get("id"), ctx.Params("username"))
	if e != nil {
		return web.JsonResponse(ctx, http.StatusForbidden, web.Forbidden, nil)
	}

	return web.JsonResponse(ctx, http.StatusOK, web.SuccessLogOut, nil)
}

func (service *workshopHandlers) Register(ctx *fiber.Ctx) error {
	payload := new(dto.WorkshopRequestRegisterBody)
	if err := ctx.BodyParser(payload); err != nil {
		return web.JsonResponse(ctx, http.StatusBadRequest, web.BadRequest, nil)
	}

	if ok, _ := helper.ValidateInputs(*payload); !ok {
		return web.JsonResponse(ctx, http.StatusBadRequest, web.CannotEmpty, nil)
	}

	workshop, err := service.WorkshopService.Register(payload.ToDomain(), payload.Street, payload.Description)
	if err != nil {
		return web.JsonResponse(ctx, http.StatusInternalServerError, web.WorkshopExist, nil)
	}

	return web.JsonResponse(ctx, http.StatusCreated, web.AccountCreated, dto.FromDomain(workshop))
}

func (service *workshopHandlers) GetWorkshop(ctx *fiber.Ctx) error {
	workshop, err := service.WorkshopService.GetWorkshop(ctx.Params("username"))
	if err != nil {
		return web.JsonResponse(ctx, http.StatusOK, web.WorkshopNotExist, nil)
	}

	return web.JsonResponse(ctx, http.StatusOK, web.Success, dto.FromDomain(workshop))
}

func (service *workshopHandlers) UpdateAccount(ctx *fiber.Ctx) error {
	account := new(dto.WorkshopAccountUpdateBody)
	e := ctx.BodyParser(account)
	if e != nil {
		return web.JsonResponse(ctx, http.StatusBadRequest, web.BadRequest, nil)
	}

	if ok, _ := helper.ValidateInputs(*account); !ok {
		return web.JsonResponse(ctx, http.StatusBadRequest, web.CannotEmpty, nil)
	}

	res, err := service.WorkshopService.UpdateAccount(account.ToDomain(), ctx.Params("username"))
	if err != nil {
		return web.JsonResponse(ctx, http.StatusForbidden, web.Forbidden, nil)
	}

	return web.JsonResponse(ctx, http.StatusOK, web.UpdateSuccess, dto.FromDomainUpdate(res))
}

func (service *workshopHandlers) UpdateAddress(ctx *fiber.Ctx) error {
	address := new(dto.WorkshopAddressUpdateBody)
	e := ctx.BodyParser(address)
	if e != nil {
		return web.JsonResponse(ctx, http.StatusBadRequest, web.BadRequest, nil)
	}

	if ok, _ := helper.ValidateInputs(*address); !ok {
		return web.JsonResponse(ctx, http.StatusBadRequest, web.CannotEmpty, nil)
	}

	res, err := service.WorkshopService.UpdateAddress(address.ToDomain(), ctx.Params("username"))
	if err != nil {
		return web.JsonResponse(ctx, http.StatusForbidden, web.Forbidden, nil)
	}

	return web.JsonResponse(ctx, http.StatusOK, web.UpdateSuccess, dto.FromDomainAddress(res)) //TODO
}

func (service *workshopHandlers) GetAddress(ctx *fiber.Ctx) error {
	address, err := service.WorkshopService.GetAddress(ctx.Params("username"))
	if err != nil {
		return web.JsonResponse(ctx, http.StatusForbidden, web.Forbidden, nil)
	}

	return web.JsonResponse(ctx, http.StatusOK, web.Success, dto.FromDomainAddress(address))
}

func (service *workshopHandlers) UpdateDescription(ctx *fiber.Ctx) error {
	desc := new(dto.WorkshopDescriptionUpdateBody)
	e := ctx.BodyParser(desc)
	if e != nil {
		return web.JsonResponse(ctx, http.StatusBadRequest, web.BadRequest, nil)
	}

	if ok, _ := helper.ValidateInputs(*desc); !ok {
		return web.JsonResponse(ctx, http.StatusBadRequest, web.CannotEmpty, nil)
	}

	res, err := service.WorkshopService.UpdateDescription(desc.ToDomain(), ctx.Params("username"))
	if err != nil {
		return web.JsonResponse(ctx, http.StatusForbidden, web.Forbidden, nil)
	}

	return web.JsonResponse(ctx, http.StatusOK, web.UpdateSuccess, dto.FromDomainDescription(res))
}

func (service *workshopHandlers) ServicesNew(ctx *fiber.Ctx) error {
	new := new(dto.ServicesNewReq)
	e := ctx.BodyParser(new)
	if e != nil {
		return web.JsonResponse(ctx, http.StatusBadRequest, web.BadRequest, nil)
	}

	if ok, _ := helper.ValidateInputs(*new); !ok {
		return web.JsonResponse(ctx, http.StatusBadRequest, web.CannotEmpty, nil)
	}

	res, err := service.WorkshopService.ServicesNew(new.ToDomain(), ctx.Params("username"))
	if err != nil {
		return web.JsonResponse(ctx, http.StatusForbidden, web.Forbidden, nil)
	}

	return web.JsonResponse(ctx, http.StatusCreated, web.ServicesCreated, dto.FromDomainServices(res))
}

func (service *workshopHandlers) UpdateServices(ctx *fiber.Ctx) error {
	s_update := new(dto.ServicesNewReq)
	e := ctx.BodyParser(s_update)
	if e != nil {
		return web.JsonResponse(ctx, http.StatusBadRequest, web.BadRequest, nil)
	}

	if ok, _ := helper.ValidateInputs(*s_update); !ok {
		return web.JsonResponse(ctx, http.StatusBadRequest, web.CannotEmpty, nil)
	}

	res, err := service.WorkshopService.UpdateServices(s_update.ToDomain(), ctx.Params("username"), ctx.Params("serviceId"))
	if err == gorm.ErrRecordNotFound {
		return web.JsonResponse(ctx, http.StatusOK, web.DataNotFound, nil)
	}
	if err != nil {
		return web.JsonResponse(ctx, http.StatusForbidden, web.Forbidden, nil)
	}

	return web.JsonResponse(ctx, http.StatusOK, web.ServicesUpdated, dto.FromDomainServices(res))
}

func (service *workshopHandlers) DeleteServices(ctx *fiber.Ctx) error {
	err := service.WorkshopService.DeleteServices(ctx.Params("username"), ctx.Params("serviceId"))
	if err == gorm.ErrRecordNotFound {
		return web.JsonResponse(ctx, http.StatusOK, web.DataNotFound, nil)
	}
	if err != nil {
		return web.JsonResponse(ctx, http.StatusForbidden, web.Forbidden, nil)
	}

	return web.JsonResponse(ctx, http.StatusOK, web.ServicesDeleted, nil)
}