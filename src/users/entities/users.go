package entities

import (
	"time"

	"gorm.io/datatypes"
)

type User struct {
	ID        uint64
	Username  string
	Email     string
	Password  string
	Name      string
	Gender    string
	DOB       datatypes.Date
	Phone     string
	Address   UserAddress
	// Orders    []Order
	// Ratings   []Rating
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserAddress struct {
	ID             uint64
	UserID         uint64
	BuildingNumber uint16
	Street         string
	City           string
	Country	       string
	PostalCode     string
	Province       string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
