package authz

type Role struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Code string `gorm:"size:500;uniqueIndex;not null" json:"code"`
}

type Permission struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Resource string `gorm:"size:64;not null;uniqueIndex:ux_perm_res_act" json:"resource"`
	Action   string `gorm:"size:32;not null;uniqueIndex:ux_perm_res_act" json:"action"`
	Scope    uint   `gorm:"not null" json:"scope"`
}

type UserRole struct {
	UserID uint `gorm:"primaryKey;not null;index" json:"user_id"`
	RoleID uint `gorm:"primaryKey;not null;index" json:"role_id"`
}

func (UserRole) TableName() string { return "user_roles" }

type RolePermission struct {
	RoleID       uint `gorm:"primaryKey;not null;index" json:"role_id"`
	PermissionID uint `gorm:"primaryKey;not null;index" json:"permission_id"`

	Role       Role       `gorm:"foreignKey:RoleID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"`
	Permission Permission `gorm:"foreignKey:PermissionID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"`
}

func (RolePermission) TableName() string { return "role_permissions" }
