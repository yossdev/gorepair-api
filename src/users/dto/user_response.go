package dto

import (
	"fmt"
	"gorepair-rest-api/internal/utils/auth"
	"gorepair-rest-api/src/users/entities"
)

type UserResponseBody struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type UserTokenResponseBody struct {
	Type         string `json:"type"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

func FromDomain(domain entities.Users) UserResponseBody {
	return UserResponseBody{
		ID: 		fmt.Sprintf("%d", domain.ID),
		Username: 	domain.Username,
		Name: 		domain.Name,
		Email: 		domain.Email,
		Phone: 		domain.Phone,
	}
}

func FromAuth(auth auth.TokenStruct) UserTokenResponseBody {
	return UserTokenResponseBody{
		Type: auth.Type,
		Token: auth.Token,
		RefreshToken: auth.RefreshToken,
	}
}