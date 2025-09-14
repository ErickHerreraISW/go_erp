package products

import "time"

type Product struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"size:200;not null" json:"name"`
	SKU         string    `gorm:"size:100;uniqueIndex" json:"sku"`
	Price       float64   `gorm:"not null" json:"price"`
	Stock       int       `gorm:"not null;default:0" json:"stock"`
	Description string    `gorm:"type:text" json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
