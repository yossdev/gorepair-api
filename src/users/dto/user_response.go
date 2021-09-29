package dto

import (
	"fmt"
	"gorepair-rest-api/src/users/entities"
)

type UserResponseBody struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

func FromDomain(domain *entities.Users) UserResponseBody {
	return UserResponseBody{
		ID: 		fmt.Sprintf("%d", domain.ID),
		Username: 	domain.Username,
		Name: 		domain.Name,
		Email: 		domain.Email,
		Phone: 		domain.Phone,
	}
}