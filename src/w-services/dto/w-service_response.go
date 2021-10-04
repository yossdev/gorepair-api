package dto

import (
	"gorepair-rest-api/src/w-services/entities"
	_ws "gorepair-rest-api/src/workshops/entities"
)

type ServiceDetailsResp struct {
	ID          uint64 `json:"id"`
	WorkshopID  uint64 `json:"workshop_id"`
	Vehicle     string `json:"vehicle"`
	VehicleType string `json:"vehicle_type"`
	Services    string `json:"services"`
	Price       int    `json:"price"`
}

type WorkshopServicesByIP struct {
	WorkshopID  uint64 `json:"workshop_id"`
}

func FromDomainGetServices(domain entities.WServices) ServiceDetailsResp {
	return ServiceDetailsResp{
		ID:          domain.ID,
		WorkshopID:  domain.WorkshopID,
		Vehicle:     domain.Vehicle,
		VehicleType: domain.VehicleType,
		Services:    domain.Services,
		Price:       domain.Price,
	}
}

func FromDomainGetServicesSlice(domain []entities.WServices) []ServiceDetailsResp {
	res := []ServiceDetailsResp{}

	for _, val := range domain {
		res = append(res, FromDomainGetServices(val))
	}
	return res
}

func FromDomain(rec _ws.WorkshopAddress) WorkshopServicesByIP {
	return WorkshopServicesByIP{
		WorkshopID: rec.WorkshopID,
	}
}

func FromDomainWS(rec []_ws.WorkshopAddress) []WorkshopServicesByIP {
	res := []WorkshopServicesByIP{}

	for _, val := range rec {
		res = append(res, FromDomain(val))
	}
	return res
}