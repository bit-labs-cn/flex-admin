package model

import (
	"gorm.io/gorm"
)

type MenuModel struct {
	gorm.Model
	Name string `gorm:"comment:菜单名称;" json:"name"`
	Url  string `json:"comment:菜单指向的链接;" json:"url"`
	Sort int    `gorm:"comment:菜单排序;" json:"sort"`
	Icon string `gorm:"comment:菜单图标;" json:"icon"`
	Meta string `gorm:"comment:菜单描述;" json:"meta"`

	ParentId string `gorm:"comment:父级菜单;" json:"parentId"`
	Path     string `gorm:"comment:菜单路径;" json:"path"`
	Level    int    `gorm:"comment:菜单层级;" json:"level"`
}

func (i *MenuModel) TableName() string {
	return "admin_menu"
}
