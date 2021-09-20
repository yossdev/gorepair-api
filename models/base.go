package models

import "time"

type BaseResponse struct {
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

type Rating struct {
	UserID      uint64    `json:"userId"`
	WorkshopID  uint64    `json:"workshopId"`
	Rating      float32   `gorm:"not null" json:"rating"`
	Description string    `json:"description" form:"description"`
	CreatedAt   time.Time `json:"createdAt"`
}