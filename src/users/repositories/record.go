package repositories

import (
	"gorepair-rest-api/src/users/entities"
	"time"
)

type User struct {
	ID       uint64 		`gorm:"primaryKey; autoIncrement"`
	Username string 		`gorm:"size:155; unique; not null"`
	Email    string 		`gorm:"size:255; unique; not null"`
	Password string 		`gorm:"size:255; not null"`
	Name     string 		`gorm:"size:125; not null"`
	Gender   string 		`gorm:"size:1"`
	DOB		 time.Time
	Phone    string 		`gorm:"size:13; not null"`
	Address  UserAddress   	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserAddress struct {
	ID             uint64 		`gorm:"primaryKey; autoIncrement"`
	UserID         uint64		`gorm:"unique"`
	BuildingNumber string		`gorm:"size:125"`
	Street         string 		`gorm:"size:255"`
	City           string 		`gorm:"size:50"`
	Country	       string 		`gorm:"size:125"`
	PostalCode     string 		`gorm:"size:10"`
	Province       string 		`gorm:"size:50"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (rec *User) toDomain() *entities.Users {
	return &entities.Users{
		ID: 		rec.ID,
		Username: 	rec.Username,
		Email: 		rec.Email,
		Password: 	rec.Password,
		Name: 		rec.Name,
		Gender:		rec.Gender,
		DOB: 		rec.DOB,
		Phone: 		rec.Phone,
	}
}

func fromDomain(domain entities.Users) *User {
	return &User{
		Username:  	domain.Username,
		Email: 		domain.Email,
		Password:  	domain.Password,
		Name:      	domain.Name,
		Gender: 	domain.Gender,
		DOB: 		domain.DOB,
		Phone: 		domain.Phone,
	}
}

func (rec *UserAddress) toDomain() *entities.UserAddress {
	return &entities.UserAddress{
		ID: 			rec.ID,
		UserID: 		rec.UserID,
		BuildingNumber: rec.BuildingNumber,
		Street: 		rec.Street,
		City: 			rec.City,
		Country: 		rec.Country,
		PostalCode: 	rec.PostalCode,
		Province: 		rec.Province,
	}
}

func fromDomainAddress(payload *entities.UserAddress, address *UserAddress) {
	address.BuildingNumber = payload.BuildingNumber
	address.Street = payload.Street
	address.City = payload.City
	address.Country = payload.Country
	address.PostalCode = payload.PostalCode
	address.Province = payload.Province
}

func fromDomainAccount(payload *entities.Users, account *User) {
	account.Username = payload.Username
	account.Email = payload.Email
	account.Password = payload.Password
	account.Name = payload.Name
	account.Gender = payload.Gender
	account.DOB = payload.DOB
	account.Phone = payload.Phone
}