package models

type ApiResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Login struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type SignUp struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Phone    string `json:"phone" form:"phone"`
}