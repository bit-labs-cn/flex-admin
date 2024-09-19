package model

import (
	"gorm.io/gorm"
)

type UserRoleModel struct {
	gorm.Model
	UserId uint
	RoleId uint
}

func (i *UserRoleModel) TableName() string {
	return "admin_user_role"
}
