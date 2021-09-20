package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type User struct {
	ID        uint64         `gorm:"primaryKey; autoIncrement" json:"id"`
	Email     string         `gorm:"size:255; unique; not null" json:"email" form:"email"`
	Password  string         `gorm:"size:255; not null" json:"password" form:"password"`
	Name      string         `gorm:"size:125; not null" json:"name" form:"name"`
	Gender    string         `gorm:"size:1" json:"gender" form:"gender"`
	DOB       datatypes.Date `json:"dob" form:"dob"`
	Phone     string         `gorm:"size:13; not null" json:"phone" form:"phone"`
	Address   UserAddress    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"address"`
	Orders    []Order        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"orders"`
	Ratings   []Rating   	 `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"ratings"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
}

type UserAddress struct {
	ID             uint64         `gorm:"primaryKey; autoIncrement" json:"id"`
	UserID         uint64         `json:"userId"`
	BuildingNumber uint16         `json:"buildingNumber" form:"buildingNumber"`
	Street         string         `gorm:"size:255" json:"street" form:"street"`
	City           string         `gorm:"size:50" json:"city" form:"city"`
	CountryCode    string         `gorm:"size:5" json:"countryCode" form:"countryCode"`
	PostalCode     string         `gorm:"size:10" json:"postalCode" form:"postalCode"`
	Province       string         `gorm:"size:50" json:"province" form:"province"`
	Lat 		   float64		  `json:"lat"`
	Lng 		   float64		  `json:"lng"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}