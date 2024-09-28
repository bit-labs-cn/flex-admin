package model

import "gorm.io/gorm/schema"

var _ schema.Tabler = (*ModelRole)(nil)

type ModelRole struct {
	ModelBase
	Status int    `gorm:"comment:角色状态" json:"status"`
	Name   string `gorm:"comment:角色名称" json:"name"`
	Code   string `gorm:"comment:角色编码" json:"code"`
	Remark string `gorm:"comment:角色描述" json:"remark"`

	Users []ModelUser `gorm:"many2many:user_role;joinForeignKey:user_id;JoinReferences:role_id" json:"roles"`
	Menus []ModelMenu `gorm:"many2many:role_menu;joinForeignKey:role_id;References:id;JoinReferences:menu_id" json:"menus"`
}

func (i ModelRole) TableName() string {
	return "admin_role"
}
