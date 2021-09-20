package models

import (
	"time"

	"gorm.io/gorm"
)

type Workshop struct {
	ID               uint64           `gorm:"primaryKey; autoIncrement" json:"id"`
	Email            string           `gorm:"size:255; unique; not null" json:"email" form:"email"`
	Password         string           `gorm:"size:255; not null" json:"password" form:"password"`
	Name             string           `gorm:"size:125; not null" json:"name" form:"name"`
	Phone            string           `gorm:"size:15; not null" json:"phone" form:"phone"`
	OperationalStart string           `gorm:"size:6; not null" json:"operationalStart" form:"operationalStart"`
	OperationalEnd   string           `gorm:"size:6; not null" json:"operationalEnd" form:"operationalEnd"`
	Description      Description      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"description"`
	Address          WorkshopAddress  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"address"`
	Services         []Service        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"services"`
	Orders           []Order          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"orders"`
	Ratings          []Rating		  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"ratings"`
	CreatedAt        time.Time        `json:"createdAt"`
	UpdatedAt        time.Time        `json:"updatedAt"`
}

type WorkshopAddress struct {
	ID             uint64         `gorm:"primaryKey; autoIncrement" json:"id"`
	WorkshopID     uint64         `json:"workshopId"`
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

type Description struct {
	ID          uint64         `gorm:"primaryKey; autoIncrement" json:"id"`
	WorkshopID  uint64         `json:"workshopId"`
	Description string         `json:"description" form:"description"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

type Service struct {
	ID                 uint64         `gorm:"primaryKey; autoIncrement" json:"id"`
	WorkshopID         uint64         `json:"workshopId"`
	Vehicle            string         `gorm:"size:125" json:"vehicle" form:"vehicle"`
	Type               string         `gorm:"size:45" json:"type" form:"type"`
	Fullservice        bool           `json:"fullservice"`
	FsPrice            uint64         `json:"fsPrice" form:"fsPrice"`
	RoutineMaintenance bool           `json:"routineMaintenance"`
	RmPrice            uint64         `json:"rmPrice" form:"rmPrice"`
	MachineRepair      bool           `json:"machineRepair"`
	MrPrice            uint64         `json:"mrPrice" form:"mrPrice"`
	BodyRepair         bool           `json:"bodyRepair"`
	BrPrice            uint64         `json:"brPrice" form:"brPrice"`
	ElectricalRepair   bool           `json:"electricalRepair"`
	ErPrice            uint64         `json:"erPrice" form:"erPrice"`
	Orders             []Order        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"orders"`
	CreatedAt          time.Time      `json:"createdAt"`
	UpdatedAt          time.Time      `json:"updatedAt"`
	DeletedAt          gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}