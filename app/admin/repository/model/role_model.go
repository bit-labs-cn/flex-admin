package model

import (
	"gorm.io/gorm"
)

type RoleModel struct {
	gorm.Model
	Name string `json:"name"`
}

func (i *RoleModel) TableName() string {
	return "admin_role"
}
