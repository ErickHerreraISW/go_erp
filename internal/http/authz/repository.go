package authz

import (
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	WithTx(tx *gorm.DB) Repository
	UserHasPermission(ctx context.Context, userID uint, resource, action string) (bool, error)
	CreateRole(role *Role) error
	CreatePermission(role *Permission) error
	CreateRolePermission(role *RolePermission) error
	CreateUserRole(role *UserRole) error
}

type repo struct{ db *gorm.DB }

func NewRepository(db *gorm.DB) Repository { return &repo{db: db} }

func (r *repo) WithTx(tx *gorm.DB) Repository { return &repo{db: tx} }

func (r *repo) UserHasPermission(ctx context.Context, userID uint, resource, action string) (bool, error) {
	var cnt int64
	err := r.db.WithContext(ctx).
		Table("user_roles AS ur").
		Joins("JOIN role_permissions rp ON rp.role_id = ur.role_id").
		Joins("JOIN permissions p ON p.id = rp.permission_id").
		Where("ur.user_id = ? AND p.resource = ? AND p.action = ?", userID, resource, action).
		Count(&cnt).Error
	if err != nil {
		return false, err
	}
	return cnt > 0, nil
}

func (r *repo) CreateRole(role *Role) error { return r.db.Create(role).Error }

func (r *repo) CreatePermission(role *Permission) error { return r.db.Create(role).Error }

func (r *repo) CreateRolePermission(role *RolePermission) error { return r.db.Create(role).Error }

func (r *repo) CreateUserRole(role *UserRole) error { return r.db.Create(role).Error }
