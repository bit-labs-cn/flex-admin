package model

import (
	"gorm.io/gorm"
	"time"
)

type UserRoleModel struct {
	gorm.Model
	Name        string     `json:"name"`
	Account     string     `json:"account"`
	Password    string     `gorm:"comment:api路径"       json:"password"`
	Description *string    `gorm:"comment:api中文描述" json:"description"`
	Phone       *string    `gorm:"comment:api组" json:"phone"`
	Email       *string    `gorm:"comment:api组" json:"email"`
	VerifiedAt  *time.Time `gorm:"comment:api组" json:"verified_at"`
}

func (i *UserRoleModel) TableName() string {
	return "admin_user"
}
