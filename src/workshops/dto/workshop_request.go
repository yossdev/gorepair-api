package dto

import (
	"gorepair-rest-api/src/workshops/entities"
)

type WorkshopRequestLoginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type WorkshopRequestRegisterBody struct {
	Username         string `json:"username" validate:"required"`
	Name             string `json:"name" validate:"required"`
	Email            string `json:"email" validate:"required"`
	Password         string `json:"password" validate:"required"`
	Phone            string `json:"phone" validate:"required"`
	OperationalStart string `json:"operational_start" validate:"required"`
	OperationalEnd   string `json:"operational_end" validate:"required"`
	Street           string `json:"street" validate:"required"`
	Description 	 string	`json:"description" validate:"required"`
}

type WorkshopAddressUpdateBody struct {
	BuildingNumber string `json:"building_number" validate:"required"`
	Street         string `json:"street" validate:"required"`
	City           string `json:"city" validate:"required"`
	Country        string `json:"country" validate:"required"`
	PostalCode     string `json:"postal_code" validate:"required"`
	Province       string `json:"province" validate:"required"`
}

type WorkshopAccountUpdateBody struct {
	Username         string `json:"username" validate:"required"`
	Name             string `json:"name" validate:"required"`
	Email            string `json:"email" validate:"required"`
	Password         string `json:"password" validate:"required"`
	Phone            string `json:"phone" validate:"required"`
	OperationalStart string `json:"operational_start" validate:"required"`
	OperationalEnd   string `json:"operational_end" validate:"required"`
}

type WorkshopDescriptionUpdateBody struct {
	Description	string `json:"description" validate:"required"`
}

type ServicesNewReq struct {
	Vehicle     string	`json:"vehicle" validate:"required"`
	VehicleType string	`json:"vehicle_type" validate:"required"`
	Services    string	`json:"services" validate:"required"`
	Price       int		`json:"price" validate:"required"`
}

func (req *WorkshopRequestLoginBody) ToDomain() *entities.Workshops {
	return &entities.Workshops{
		Email:    req.Email,
		Password: req.Password,
	}
}

func (req *WorkshopRequestRegisterBody) ToDomain() *entities.Workshops {
	return &entities.Workshops{
		Username:         req.Username,
		Name:             req.Name,
		Email:            req.Email,
		Password:         req.Password,
		Phone:            req.Phone,
		OperationalStart: req.OperationalStart,
		OperationalEnd:   req.OperationalEnd,
	}
}

func (req *WorkshopAddressUpdateBody) ToDomain() *entities.WorkshopAddress {
	return &entities.WorkshopAddress{
		BuildingNumber: req.BuildingNumber,
		Street:         req.Street,
		City:           req.City,
		Country:        req.Country,
		PostalCode:     req.PostalCode,
		Province:       req.Province,
	}
}

func (req *WorkshopAccountUpdateBody) ToDomain() *entities.Workshops {
	return &entities.Workshops{
		Username:         req.Username,
		Name:             req.Name,
		Email:            req.Email,
		Password:         req.Password,
		Phone:            req.Phone,
		OperationalStart: req.OperationalStart,
		OperationalEnd:   req.OperationalEnd,
	}
}

func (req *WorkshopDescriptionUpdateBody) ToDomain() *entities.Descriptions {
	return &entities.Descriptions{
		Description: req.Description,
	}
}

func (req *ServicesNewReq) ToDomain() *entities.Services {
	return &entities.Services{
		Vehicle: 		req.Vehicle,
		VehicleType: 	req.VehicleType,
		Services: 		req.Services,
		Price: 			req.Price,
	}
}