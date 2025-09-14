package erpinstanceuser

import (
	"github.com/ErickHerreraISW/go_erp/internal/feature/erpinstance"
	"github.com/ErickHerreraISW/go_erp/internal/feature/users"
)

type ErpInstanceUser struct {
	ID            uint `gorm:"primaryKey" json:"id"`
	ErpInstanceID uint `gorm:"index"`
	ErpInstance   erpinstance.ErpInstance
	UserID        uint `gorm:"index"`
	User          users.User
}
