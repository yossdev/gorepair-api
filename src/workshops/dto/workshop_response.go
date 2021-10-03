package dto

import (
	"gorepair-rest-api/src/workshops/entities"
)

type WorkshopResponseBody struct {
	ID               uint64 `json:"id"`
	Username         string `json:"username"`
	Email            string `json:"email"`
	Name             string `json:"name"`
	Phone            string `json:"phone"`
	OperationalStart string `json:"operational_start"`
	OperationalEnd   string `json:"operational_end"`
	Description 	 DescriptionUpdateRes `json:"description"`
}

type DescriptionUpdateRes struct {
	ID			uint64	`json:"id"`
	WorkshopID	uint64	`json:"workshop_id"`
	Description	string	`json:"description"`
}

type ServicesNewRes struct {
	ID          uint64	`json:"id"`
	WorkshopID  uint64	`json:"workshop_id"`
	Vehicle     string	`json:"vehicle"`
	VehicleType string	`json:"vehicle_type"`
	Services    string	`json:"services"`
	Price       int		`json:"price"`
}

type WorkshopUpdateResponseBody struct {
	ID               uint64 `json:"id"`
	Username         string `json:"username"`
	Email            string `json:"email"`
	Name             string `json:"name"`
	Phone            string `json:"phone"`
	OperationalStart string `json:"operational_start"`
	OperationalEnd   string `json:"operational_end"`
}

type WorkshopResponseAddressBody struct {
	ID             uint64 `json:"id"`
	WorkshopID     uint64 `json:"workshop_id"`
	BuildingNumber string `json:"building_number"`
	Street         string `json:"street"`
	City           string `json:"city"`
	Country        string `json:"country"`
	PostalCode     string `json:"postal_code"`
	Province       string `json:"province"`
}

func FromDomain(domain *entities.Workshops) WorkshopResponseBody {
	return WorkshopResponseBody{
		ID:       			domain.ID,
		Username: 			domain.Username,
		Email:    			domain.Email,
		Name:     			domain.Name,
		Phone:    			domain.Phone,
		OperationalStart: 	domain.OperationalStart,
		OperationalEnd:   	domain.OperationalEnd,
		Description: 		DescriptionUpdateRes{
			ID: 			domain.Description.ID,
			WorkshopID: 	domain.Description.WorkshopID,
			Description: 	domain.Description.Description,
		},
	}
}

func FromDomainUpdate(domain *entities.Workshops) WorkshopUpdateResponseBody {
	return WorkshopUpdateResponseBody{
		ID: 				domain.ID,
		Username: 			domain.Username,
		Email: 				domain.Email,
		Name: 				domain.Name,
		Phone: 				domain.Phone,
		OperationalStart: 	domain.OperationalStart,
		OperationalEnd:   	domain.OperationalEnd,
	}
}

func FromDomainAddress(domain *entities.WorkshopAddress) WorkshopResponseAddressBody {
	return WorkshopResponseAddressBody{
		ID: 			domain.ID,
		WorkshopID: 	domain.WorkshopID,
		BuildingNumber: domain.BuildingNumber,
		Street: 		domain.Street,
		City: 			domain.City,
		Country: 		domain.Country,
		PostalCode: 	domain.PostalCode,
		Province: 		domain.Province,
	}
}

func FromDomainDescription(domain *entities.Descriptions) DescriptionUpdateRes {
	return DescriptionUpdateRes{
		ID: 			domain.ID,
		WorkshopID: 	domain.WorkshopID,
		Description: 	domain.Description,
	}
}

func FromDomainServices(domain *entities.Services) ServicesNewRes {
	return ServicesNewRes{
		ID: 			domain.ID,
		WorkshopID: 	domain.WorkshopID,
		Vehicle: 		domain.Vehicle,
		VehicleType: 	domain.VehicleType,
		Services: 		domain.Services,
		Price: 			domain.Price,
	}
}