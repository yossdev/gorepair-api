package tables

import "time"

type UserLoginLog struct {
	ID        uint64    `gorm:"primaryKey; autoIncrement" json:"id"`
	UserID    uint64    `json:"userId"`
	Ip        string    `gorm:"size:15; not null" json:"ip"`
	City      string    `gorm:"size:125; not null" json:"city"`
	CreatedAt time.Time `json:"createdAt"`
}