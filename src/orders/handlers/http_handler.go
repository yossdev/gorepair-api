package handlers

import (
	"gorepair-rest-api/internal/utils/helper"
	"gorepair-rest-api/internal/web"
	"gorepair-rest-api/src/orders/dto"
	"gorepair-rest-api/src/orders/entities"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type OrderHandlers interface {
	OrderNew(ctx *fiber.Ctx) error
	GetUserOrderDetails(ctx *fiber.Ctx) error
	GetWorkshopOrderDetails(ctx *fiber.Ctx) error
	UserCancelOrder(ctx *fiber.Ctx) error
}

type orderHandlers struct {
	OrderService entities.OrderService
}

func NewHttpHandler(orderService entities.OrderService) OrderHandlers {
	return &orderHandlers{
		OrderService: orderService,
	}
}

func (service *orderHandlers) OrderNew(ctx *fiber.Ctx) error {
	new := new(dto.OrderRequestBody)
	e := ctx.BodyParser(new)
	if e != nil {
		return web.JsonResponse(ctx, http.StatusBadRequest, "something is not right with your request", nil)
	}

	if ok, _ := helper.ValidateInputs(*new); !ok {
		return web.JsonResponse(ctx, http.StatusBadRequest, "field cannot be empty", nil)
	}

	res, err := service.OrderService.OrderNew(new.ToDomain(), ctx.Get("id"))
	if err != nil {
		return web.JsonResponse(ctx, http.StatusBadRequest, "problem with db", nil)
	}

	return web.JsonResponse(ctx, http.StatusCreated, "your order is placed!", dto.FromDomainOrder(res))
}

func (service *orderHandlers) GetUserOrderDetails(ctx *fiber.Ctx) error {
	res, err := service.OrderService.GetUserOrderDetails(ctx.Params("orderId"), ctx.Get("id"))
	if err != nil {
		return web.JsonResponse(ctx, http.StatusBadRequest, "order is not exist", nil)
	}

	return web.JsonResponse(ctx, http.StatusOK, "success", dto.FromDomainOrderGet(res))
}

func (service *orderHandlers) GetWorkshopOrderDetails(ctx *fiber.Ctx) error {
	res, err := service.OrderService.GetWorkshopOrderDetails(ctx.Params("orderId"), ctx.Get("id"))
	if err != nil {
		return web.JsonResponse(ctx, http.StatusBadRequest, "order is not exist", nil)
	}

	return web.JsonResponse(ctx, http.StatusOK, "success", dto.FromDomainOrderGet(res))
}

func (service *orderHandlers) UserCancelOrder(ctx *fiber.Ctx) error {
	err := service.OrderService.UserCancelOrder(ctx.Params("orderId"), ctx.Get("id"), ctx.Params("username"))
	if err != nil {
		return web.JsonResponse(ctx, http.StatusBadRequest, "order is not exist", nil)
	}

	return web.JsonResponse(ctx, http.StatusOK, "order canceled!", nil)
}

