package repositories

import (
	"gorepair-rest-api/src/workshops/entities"
	"time"
)

type Workshop struct {
	ID               uint64           `gorm:"primaryKey; autoIncrement"`
	Username 		 string 		  `gorm:"size:155; unique; not null"`
	Email            string           `gorm:"size:255; unique; not null"`
	Password         string           `gorm:"size:255; not null"`
	Name             string           `gorm:"size:125; not null"`
	Phone            string           `gorm:"size:15; not null"`
	OperationalStart string           `gorm:"size:6; not null"`
	OperationalEnd   string           `gorm:"size:6; not null"`
	Description      Description      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Address          WorkshopAddress  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Services         []Service        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	// Orders           []Order          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type WorkshopAddress struct {
	ID             uint64 	`gorm:"primaryKey; autoIncrement"`
	WorkshopID     uint64 	`gorm:"unique"`
	BuildingNumber string 	`gorm:"size:125"`
	Street         string 	`gorm:"size:255"`
	City           string 	`gorm:"size:50"`
	Country 	   string 	`gorm:"size:5"`
	PostalCode     string 	`gorm:"size:10"`
	Province       string	`gorm:"size:50"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Description struct {
	ID          uint64    `gorm:"primaryKey; autoIncrement" json:"id"`
	WorkshopID  uint64    `gorm:"unique"`
	Description string         
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Service struct {
	ID          uint64	`gorm:"primaryKey; autoIncrement" json:"id"`
	WorkshopID  uint64	`gorm:"unique"`
	Vehicle     string 	`gorm:"size:125"`
	VehicleType string 	`gorm:"size:45"`
	Services    string 	`gorm:"size:255" json:"type" form:"type"`
	Price       uint64
	// Orders      []Order	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"orders"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	// DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (rec *Workshop) toDomain() *entities.Workshops {
	return &entities.Workshops{
		ID: 				rec.ID,
		Username: 			rec.Username,
		Email: 				rec.Email,
		Password: 			rec.Password,
		Name: 				rec.Name,
		Phone: 				rec.Phone,
		OperationalStart: 	rec.OperationalStart,
		OperationalEnd: 	rec.OperationalEnd,
	}
}

func fromDomain(domain entities.Workshops) *Workshop {
	return &Workshop{
		Username:  			domain.Username,
		Email: 				domain.Email,
		Password:  			domain.Password,
		Name:      			domain.Name,
		Phone: 				domain.Phone,
		OperationalStart: 	domain.OperationalStart,
		OperationalEnd: 	domain.OperationalEnd,
	}
}

func (rec *WorkshopAddress) toDomain() *entities.WorkshopAddress {
	return &entities.WorkshopAddress{
		ID: 			rec.ID,
		WorkshopID: 	rec.WorkshopID,
		BuildingNumber: rec.BuildingNumber,
		Street: 		rec.Street,
		City: 			rec.City,
		Country: 		rec.Country,
		PostalCode: 	rec.PostalCode,
		Province: 		rec.Province,
	}
}

func fromDomainAddress(payload *entities.WorkshopAddress, address *WorkshopAddress) {
	address.BuildingNumber = payload.BuildingNumber
	address.Street = payload.Street
	address.City = payload.City
	address.Country = payload.Country
	address.PostalCode = payload.PostalCode
	address.Province = payload.Province
}

func fromDomainAccount(payload *entities.Workshops, account *Workshop) {
	account.Username = payload.Username
	account.Email = payload.Email
	account.Password = payload.Password
	account.Name = payload.Name
	account.Phone = payload.Phone
	account.OperationalStart = payload.OperationalStart
	account.OperationalEnd = payload.OperationalEnd
}