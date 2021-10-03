package handlers

import (
	"fmt"
	"gorepair-rest-api/internal/utils/helper"
	"gorepair-rest-api/internal/web"
	"gorepair-rest-api/src/workshops/dto"
	"gorepair-rest-api/src/workshops/entities"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type WorkshopHandlers interface {
	GetWorkshop(ctx *fiber.Ctx) error
	Register(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
	Logout(ctx *fiber.Ctx) error
	UpdateAccount(ctx *fiber.Ctx) error
	UpdateAddress(ctx *fiber.Ctx) error
	GetAddress(ctx *fiber.Ctx) error
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
		log.Fatal(err)
	}

	res, err := service.WorkshopService.Login(payload.ToDomain())
	if err != nil {
		return web.JsonResponse(ctx, http.StatusUnauthorized, "email or password is wrong!", nil)
	}

	return web.JsonResponse(ctx, http.StatusOK, "welcome!", res)
}

func (service *workshopHandlers) Logout(ctx *fiber.Ctx) error {
	err := service.WorkshopService.FindByID(ctx.Get("id"))
	if err != nil {
		return web.JsonResponse(ctx, http.StatusForbidden, "forbidden", nil)
	}

	workshop, err := service.WorkshopService.GetWorkshop(ctx.Params("username"))
	if err != nil {
		return web.JsonResponse(ctx, http.StatusForbidden, "forbidden", nil)
	}

	e := service.WorkshopService.Logout(fmt.Sprintf("%d", workshop.ID))
	if e != nil {
		return web.JsonResponse(ctx, http.StatusForbidden, "forbidden", nil)
	}

	return web.JsonResponse(ctx, http.StatusOK, "successfully logged out", nil)
}

func (service *workshopHandlers) Register(ctx *fiber.Ctx) error {
	payload := new(dto.WorkshopRequestRegisterBody)
	if err := ctx.BodyParser(payload); err != nil {
		return web.JsonResponse(ctx, http.StatusBadRequest, "something is not right", nil)
	}

	if ok, _ := helper.ValidateInputs(*payload); !ok {
		return web.JsonResponse(ctx, http.StatusBadRequest, "field cannot be empty", nil)
	}

	workshop, err := service.WorkshopService.Register(payload.ToDomain(), payload.Street, payload.Description)
	if err != nil {
		return web.JsonResponse(ctx, http.StatusInternalServerError, "workshop already exist", nil)
	}

	return web.JsonResponse(ctx, http.StatusCreated, "account created!", dto.FromDomain(workshop))
}

func (service *workshopHandlers) GetWorkshop(ctx *fiber.Ctx) error {
	// err := service.WorkshopService.FindByID(ctx.Get("id"))
	// if err != nil {
	// 	return web.JsonResponse(ctx, http.StatusForbidden, "forbidden", nil)
	// }

	workshop, err := service.WorkshopService.GetWorkshop(ctx.Params("username"))
	if err != nil {
		return web.JsonResponse(ctx, http.StatusOK, "workshop is not exist", nil)
	}

	return web.JsonResponse(ctx, http.StatusOK, "success", dto.FromDomain(workshop))
}

func (service *workshopHandlers) UpdateAccount(ctx *fiber.Ctx) error {
	rec, err := service.WorkshopService.GetWorkshop(ctx.Params("username"))
	if err != nil {
		return web.JsonResponse(ctx, http.StatusForbidden, "forbidden", nil)
	}

	ok := service.WorkshopService.FindByID(fmt.Sprintf("%d", rec.ID))
	if ok != nil {
		return web.JsonResponse(ctx, http.StatusForbidden, "forbidden", nil)
	}

	account := new(dto.WorkshopAccountUpdateBody)
	e := ctx.BodyParser(account)
	if e != nil {
		return web.JsonResponse(ctx, http.StatusBadRequest, "something is not right with your request", nil)
	}

	if ok, _ := helper.ValidateInputs(*account); !ok {
		return web.JsonResponse(ctx, http.StatusBadRequest, "field cannot be empty", nil)
	}

	res, err := service.WorkshopService.UpdateAccount(account.ToDomain(), rec.ID)
	if err != nil {
		return web.JsonResponse(ctx, http.StatusBadRequest, "problem with db", nil)
	}

	return web.JsonResponse(ctx, http.StatusOK, "successfully updated!", dto.FromDomainUpdate(res))
}

func (service *workshopHandlers) UpdateAddress(ctx *fiber.Ctx) error {
	rec, err := service.WorkshopService.GetWorkshop(ctx.Params("username"))
	if err != nil {
		return web.JsonResponse(ctx, http.StatusForbidden, "forbidden", nil)
	}

	ok := service.WorkshopService.FindByID(fmt.Sprintf("%d", rec.ID))
	if ok != nil {
		return web.JsonResponse(ctx, http.StatusForbidden, "forbidden", nil)
	}

	address := new(dto.WorkshopAddressUpdateBody)
	e := ctx.BodyParser(address)
	if e != nil {
		return web.JsonResponse(ctx, http.StatusBadRequest, "something is not right with your request", nil)
	}

	if ok, _ := helper.ValidateInputs(*address); !ok {
		return web.JsonResponse(ctx, http.StatusBadRequest, "field cannot be empty", nil)
	}

	res, err := service.WorkshopService.UpdateAddress(address.ToDomain(), rec.ID)
	if err != nil {
		return web.JsonResponse(ctx, http.StatusBadRequest, "problem with db", nil)
	}

	return web.JsonResponse(ctx, http.StatusOK, "successfully updated!", dto.FromDomainAddress(res)) //TODO
}

func (service *workshopHandlers) GetAddress(ctx *fiber.Ctx) error {
	workshop, err := service.WorkshopService.GetWorkshop(ctx.Params("username"))
	if err != nil {
		return web.JsonResponse(ctx, http.StatusOK, "workshop is not exist", nil)
	}

	ok := service.WorkshopService.FindByID(fmt.Sprintf("%d", workshop.ID))
	if ok != nil {
		return web.JsonResponse(ctx, http.StatusForbidden, "forbidden", nil)
	}

	address, _ := service.WorkshopService.GetAddress(workshop.ID)

	return web.JsonResponse(ctx, http.StatusOK, "success", dto.FromDomainAddress(address))
}