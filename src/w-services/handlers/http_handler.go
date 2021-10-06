package handlers

import (
	"gorepair-rest-api/infrastructures/third-party/freegeoapi"
	"gorepair-rest-api/internal/web"
	"gorepair-rest-api/src/w-services/dto"
	"gorepair-rest-api/src/w-services/entities"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type WServicesHandlers interface {
	GetAll(ctx *fiber.Ctx) error
	GetDetails(ctx *fiber.Ctx) error
	GetAllWorkshop(ctx *fiber.Ctx) error
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
	res, err := s.WServicesService.GetAll()
	if err != nil {
		return web.JsonResponse(ctx, http.StatusOK, web.ServicesNotExist, nil)
	}

	return web.JsonResponse(ctx, http.StatusOK, web.Success, dto.FromDomainGetServicesSlice(res))
}

func (s *wservicesHandlers) GetDetails(ctx *fiber.Ctx) error {
	res, err := s.WServicesService.GetDetails(ctx.Params("serviceId"))
	if err != nil {
		return web.JsonResponse(ctx, http.StatusOK, web.ServicesNotExist, nil)
	}
	// ! CHANGE THISS
	ipgeo, _ := freegeoapi.NewIpAPI().GetLocationByIP()
	return web.JsonResponse(ctx, http.StatusOK, ipgeo.City, dto.FromDomainGetServices(res))
}

func (s *wservicesHandlers) GetAllWorkshop(ctx *fiber.Ctx) error {
	res, err := s.WServicesService.GetAllWorkshop()
	if err != nil {
		return web.JsonResponse(ctx, http.StatusOK, web.DataNotFound, nil)
	}

	return web.JsonResponse(ctx, http.StatusOK, web.Success, dto.FromDomainWS(res))
}