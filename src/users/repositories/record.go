package repositories

import (
	"gorepair-rest-api/src/users/entities"
	"time"

	"gorm.io/datatypes"
)

type User struct {
	ID       uint64 		`gorm:"primaryKey; autoIncrement"`
	Username string 		`gorm:"size:155; unique; not null"`
	Email    string 		`gorm:"size:255; unique; not null"`
	Password string 		`gorm:"size:255; not null"`
	Name     string 		`gorm:"size:125; not null"`
	Gender   string 		`gorm:"size:1"`
	DOB      datatypes.Date
	Phone    string 		`gorm:"size:13; not null"`
	// Address   UserAddress   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	// Orders    []Order        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	// Ratings   []Rating       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserAddress struct {
	ID             uint64 `gorm:"primaryKey; autoIncrement"`
	UserID         uint64
	BuildingNumber uint16
	Street         string `gorm:"size:255"`
	City           string `gorm:"size:50"`
	Country	       string `gorm:"size:125"`
	PostalCode     string `gorm:"size:10"`
	Province       string `gorm:"size:50"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (rec *User) ToDomain() entities.Users {
	return entities.Users{
		ID: 		rec.ID,
		Username: 	rec.Username,
		Email: 		rec.Email,
		Password: 	rec.Password,
		Name: 		rec.Name,
		Phone: 		rec.Phone,
	}
}

func FromDomain(userDomain entities.Users) *User {
	return &User{
		ID:        userDomain.ID,
		Name:      userDomain.Name,
		Username:  userDomain.Username,
		Password:  userDomain.Password,
	}
}