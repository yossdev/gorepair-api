package dto

import "gorepair-rest-api/src/w-services/entities"

type ServiceDetailsResp struct {
	ID          uint64 `json:"id"`
	WorkshopID  uint64 `json:"workshop_id"`
	Vehicle     string `json:"vehicle"`
	VehicleType string `json:"vehicle_type"`
	Services    string `json:"services"`
	Price       int    `json:"price"`
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