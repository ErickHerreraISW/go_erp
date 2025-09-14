package users

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:200;not null" json:"name"`
	Email     string    `gorm:"size:200;uniqueIndex;not null" json:"email"`
	Password  string    `gorm:"size:200;not null" json:"-"`
	Role      string    `gorm:"size:50;default:user" json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
