package tables

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID          uint64         `gorm:"primaryKey; autoIncrement" json:"id"`
	UserID      uint64         `json:"userId"`
	WorkshopID  uint64         `json:"workshopId"`
	ServiceID   uint64         `json:"serviceId"`
	OnSite      bool           `json:"onSite"`
	PickUp      bool           `json:"pickUp"`
	Note        string         `json:"note" form:"note"`
	TotalPrice  uint64         `json:"totalPrice" form:"totalPrice"`
	OrderStatus OrderStatus    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"orderStatus"`
	CreatedAt   time.Time      `json:"createdAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

type OrderStatus struct {
	OrderID     uint64    `json:"orderId"`
	Pending     bool      `json:"pending"`
	OnProcess   bool      `json:"onProcess"`
	ReadyToTake bool      `json:"readyToTake"`
	UpdatedAt   time.Time `json:"updatedAt"`
}