package entities

import "gorepair-rest-api/internal/utils/auth"

type Users struct {
	ID       uint64
	Username string
	Email    string
	Password string
	Name     string
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

type Service interface {
	FindByID(id uint64) (*Users, error)
	GetUser(username string) (*Users, error)
	Register(data *Users) (*Users, error)
	Login(data *Users) (auth.TokenStruct, error)
	RefreshToken(id string) (auth.TokenStruct, error)
}

type Repository interface {
	// FindAll() []Users
	FindByID(id uint64) (*Users, error)
	FindByEmail(email string) *Users
	GetUser(username string) (*Users, error)
	Register(data *Users) (*Users, error)
}