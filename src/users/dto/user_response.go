package dto

import (
	"gorepair-rest-api/src/users/entities"
	"time"
)

type UserResponseBody struct {
	ID       uint64 	`json:"id"`
	Username string 	`json:"username"`
	Email	 string 	`json:"email"`
	Name     string 	`json:"name"`
	Gender	 string 	`json:"gender"`
	DOB 	 time.Time 	`json:"dob"`
	Phone    string 	`json:"phone"`
}

type UserUpdateResponseBody struct {
	ID       uint64 	`json:"id"`
	Username string 	`json:"username"`
	Email    string 	`json:"email"`
	Name     string 	`json:"name"`
	Gender   string 	`json:"gender"`
	DOB 	 time.Time 	`json:"dob"`
	Phone    string 	`json:"phone"`
}

type UserResponseAddressBody struct {
	ID             uint64 `json:"id"`
	UserID         uint64 `json:"user_id"`
	BuildingNumber string `json:"building_number"`
	Street         string `json:"street"`
	City           string `json:"city"`
	Country        string `json:"country"`
	PostalCode     string `json:"postal_code"`
	Province       string `json:"province"`
}

func FromDomain(domain *entities.Users) UserResponseBody {
	return UserResponseBody{
		ID: 		domain.ID,
		Username: 	domain.Username,
		Email:		domain.Email,
		Name: 		domain.Name,
		Gender: 	domain.Gender,
		DOB: 		domain.DOB,
		Phone: 		domain.Phone,
	}
}

func FromDomainUpdate(domain *entities.Users) UserUpdateResponseBody {
	return UserUpdateResponseBody{
		ID: 		domain.ID,
		Username: 	domain.Username,
		Email: 		domain.Email,
		Name: 		domain.Name,
		Gender: 	domain.Gender,
		DOB: 		domain.DOB,
		Phone: 		domain.Phone,
	}
}

func FromDomainAddress(domain *entities.UserAddress) UserResponseAddressBody {
	return UserResponseAddressBody{
		ID: 			domain.ID,
		UserID: 		domain.UserID,
		BuildingNumber: domain.BuildingNumber,
		Street: 		domain.Street,
		City: 			domain.City,
		Country: 		domain.Country,
		PostalCode: 	domain.PostalCode,
		Province: 		domain.Province,
	}
}