package migrations

import (
	"github.com/ErickHerreraISW/go_erp/internal/feature/erpinstance"
	"github.com/ErickHerreraISW/go_erp/internal/feature/erpinstanceuser"
	"github.com/ErickHerreraISW/go_erp/internal/feature/products"
	"github.com/ErickHerreraISW/go_erp/internal/feature/users"
	"github.com/ErickHerreraISW/go_erp/internal/http/authz"
	"gorm.io/gorm"
)

func Run(db *gorm.DB) error {
	return db.AutoMigrate(
		&users.User{},
		&erpinstance.ErpInstance{},
		&erpinstanceuser.ErpInstanceUser{},
		&products.Product{},
		&authz.Role{},
		&authz.Permission{},
		&authz.UserRole{},
		&authz.RolePermission{},
	)
}
