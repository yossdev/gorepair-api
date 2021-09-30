package dto

import (
	"fmt"
	"gorepair-rest-api/src/users/entities"
	"time"
)

type UserResponseBody struct {
	ID       string 	`json:"id"`
	Username string 	`json:"username"`
	Email	 string 	`json:"email"`
	Name     string 	`json:"name"`
	Gender	 string 	`json:"gender"`
	DOB 	 time.Time 	`json:"dob"`
	Phone    string 	`json:"phone"`
}

type UserResponseAddressBody struct {
	ID             string `json:"id"`
	UserID         string `json:"user_id"`
	BuildingNumber uint16 `json:"building_number"`
	Street         string `json:"street"`
	City           string `json:"city"`
	Country        string `json:"country"`
	PostalCode     string `json:"postal_code"`
	Province       string `json:"province"`
}

type UserUpdateResponseBody struct {
	ID       string 	`json:"id"`
	Username string 	`json:"username"`
	Email    string 	`json:"email"`
	Name     string 	`json:"name"`
	Gender   string 	`json:"gender"`
	DOB 	 time.Time 	`json:"dob"`
	Phone    string 	`json:"phone"`
}

func FromDomain(domain *entities.Users) UserResponseBody {
	return UserResponseBody{
		ID: 		fmt.Sprintf("%d", domain.ID),
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
		ID: 		fmt.Sprintf("%d", domain.ID),
		Username: 	domain.Username,
		Email: 		domain.Email,
		Name: 		domain.Name,
		Gender: 	domain.Gender,
		DOB: 		domain.DOB,
		Phone: 		domain.Phone,
	}
}

func FromDomainAddressUpdate(domain *entities.UserAddress) UserResponseAddressBody {
	return UserResponseAddressBody{
		ID: 			fmt.Sprintf("%d", domain.ID),
		UserID: 		fmt.Sprintf("%d", domain.UserID),
		BuildingNumber: domain.BuildingNumber,
		Street: 		domain.Street,
		City: 			domain.City,
		Country: 		domain.Country,
		PostalCode: 	domain.PostalCode,
		Province: 		domain.Province,
		
	}
}