package dto

import (
	"fmt"
	"gorepair-rest-api/src/workshops/entities"
)

type WorkshopResponseBody struct {
	ID               string `json:"id"`
	Username         string `json:"username"`
	Email            string `json:"email"`
	Name             string `json:"name"`
	Phone            string `json:"phone"`
	OperationalStart string `json:"operational_start"`
	OperationalEnd   string `json:"operational_end"`
}

type WorkshopUpdateResponseBody struct {
	ID               string `json:"id"`
	Username         string `json:"username"`
	Email            string `json:"email"`
	Name             string `json:"name"`
	Phone            string `json:"phone"`
	OperationalStart string `json:"operational_start"`
	OperationalEnd   string `json:"operational_end"`
}

type WorkshopResponseAddressBody struct {
	ID             string `json:"id"`
	WorkshopID     string `json:"workshop_id"`
	BuildingNumber string `json:"building_number"`
	Street         string `json:"street"`
	City           string `json:"city"`
	Country        string `json:"country"`
	PostalCode     string `json:"postal_code"`
	Province       string `json:"province"`
}

func FromDomain(domain *entities.Workshops) WorkshopResponseBody {
	return WorkshopResponseBody{
		ID:       			fmt.Sprintf("%d", domain.ID),
		Username: 			domain.Username,
		Email:    			domain.Email,
		Name:     			domain.Name,
		Phone:    			domain.Phone,
		OperationalStart: 	domain.OperationalStart,
		OperationalEnd:   	domain.OperationalEnd,
	}
}

func FromDomainUpdate(domain *entities.Workshops) WorkshopUpdateResponseBody {
	return WorkshopUpdateResponseBody{
		ID: 				fmt.Sprintf("%d", domain.ID),
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
		ID: 			fmt.Sprintf("%d", domain.ID),
		WorkshopID: 	fmt.Sprintf("%d", domain.WorkshopID),
		BuildingNumber: domain.BuildingNumber,
		Street: 		domain.Street,
		City: 			domain.City,
		Country: 		domain.Country,
		PostalCode: 	domain.PostalCode,
		Province: 		domain.Province,
	}
}