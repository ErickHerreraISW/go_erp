package erpinstanceuser

import (
	"github.com/ErickHerreraISW/go_erp/internal/feature/erpinstance"
)

type ErpInstanceUser struct {
	ID            uint `gorm:"primaryKey" json:"id"`
	ErpInstanceID uint `gorm:"index"`
	ErpInstance   erpinstance.ErpInstance
	UserID        uint `gorm:"index"`
}
