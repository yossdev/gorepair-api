package entities

import "gorm.io/datatypes"

type Users struct {
	ID       uint64
	Username string
	Email    string
	Password string
	Name     string
	Gender   string
	DOB      datatypes.Date
	Phone    string
}

type UserAddress struct {
	ID             uint64
	UserID         uint64
	BuildingNumber uint16
	Street         string
	City           string
	Country        string
	PostalCode     string
	Province       string
}

type UserService interface {
	GetUser(username string) (*Users, error)
	Register(payload *Users) (*Users, error)
	Login(payload *Users) (interface{}, error)
	// Logout() error
	// Account(payload *Users) (*Users, error)
	// Address(payload *Users) (*Users, error)
}

type UserRepository interface {
	GetUser(username string) (*Users, error)
	Register(payload *Users) (*Users, error)
	FindByEmail(email string) *Users
	// Account(payload *Users) (*Users, error)
	// Address(payload *Users) (*Users, error)
}