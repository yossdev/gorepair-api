package handlers

import (
	"gorepair-rest-api/internal/web"
	"gorepair-rest-api/src/w-services/dto"
	"gorepair-rest-api/src/w-services/entities"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type WServicesHandlers interface {
	GetAll(ctx *fiber.Ctx) error
	GetDetails(ctx *fiber.Ctx) error
}

type wservicesHandlers struct {
	WServicesService entities.WServicesService
}

func NewHttpHandler(wservicesService entities.WServicesService) WServicesHandlers {
	return &wservicesHandlers{
		WServicesService: wservicesService,
	}
}

func (s *wservicesHandlers) GetAll(ctx *fiber.Ctx) error {
	res, _ := s.WServicesService.GetAll()
	return web.JsonResponse(ctx, http.StatusOK, web.Success, dto.FromDomainGetServicesSlice(res))
}

func (s *wservicesHandlers) GetDetails(ctx *fiber.Ctx) error {
	res, err := s.WServicesService.GetDetails(ctx.Params("serviceId"))
	if err != nil {
		return web.JsonResponse(ctx, http.StatusOK, web.ServicesNotExist, dto.FromDomainGetServices(res))
	}

	return web.JsonResponse(ctx, http.StatusOK, web.Success, dto.FromDomainGetServices(res))
}