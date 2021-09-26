package dto

type UserRequestLoginBody struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UserRequestRegisterBody struct {
	Username string `json:"username" form:"username" validate:"required"`
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Phone    string `json:"phone" form:"phone" validate:"required"`
}