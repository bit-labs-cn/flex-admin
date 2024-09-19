package model

import (
	"gorm.io/gorm"
)

type RoleModel struct {
	gorm.Model
	Sort        string `json:"sort"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (i *RoleModel) TableName() string {
	return "admin_role"
}
