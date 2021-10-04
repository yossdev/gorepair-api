package repositories

import (
	"gorepair-rest-api/src/orders/entities"
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID          uint64 `gorm:"primaryKey; autoIncrement" json:"id"`
	UserID      uint64
	WorkshopID  uint64
	ServiceID   uint64
	OnSite      bool
	PickUp      bool
	Note        string
	TotalPrice  int
	Placed      bool
	CreatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (rec *Order) toDomain() *entities.Orders {
	return &entities.Orders{
		ID: 			rec.ID,
		UserID: 		rec.UserID,
		WorkshopID: 	rec.WorkshopID,
		ServiceID: 		rec.ServiceID,
		OnSite: 		rec.OnSite,
		PickUp: 		rec.PickUp,
		Note: 			rec.Note,
		TotalPrice: 	rec.TotalPrice,
		Placed: 		rec.Placed,
	}
}

func (rec *Order) toDomainGet(street, phone string) *entities.Orders {
	return &entities.Orders{
		ID: 			rec.ID,
		UserID: 		rec.UserID,
		WorkshopID: 	rec.WorkshopID,
		ServiceID: 		rec.ServiceID,
		OnSite: 		rec.OnSite,
		PickUp: 		rec.PickUp,
		Note: 			rec.Note,
		TotalPrice: 	rec.TotalPrice,
		Placed: 		rec.Placed,
		Phone: phone,
		Street: street,
	}
}

func fromDomainOrder(payload *entities.Orders, order *Order) {
	order.WorkshopID = payload.WorkshopID
	order.ServiceID = payload.ServiceID
	order.OnSite = payload.OnSite
	order.PickUp = payload.PickUp
	order.Note = payload.Note
	order.Placed = true
}