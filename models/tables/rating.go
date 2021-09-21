package tables

import "time"

type Rating struct {
	UserID      uint64    `json:"userId"`
	WorkshopID  uint64    `json:"workshopId"`
	Rating      float32   `gorm:"not null" json:"rating"`
	Description string    `json:"description" form:"description"`
	CreatedAt   time.Time `json:"createdAt"`
}