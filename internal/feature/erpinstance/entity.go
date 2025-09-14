package erpinstance

import (
	"time"
)

type ErpInstance struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	ErpKey    string    `gorm:"size:500;not null" json:"erpkey"`
	CreatedAt time.Time `json:"created_at"`
}
