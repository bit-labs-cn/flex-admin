package model

import "bit-labs.cn/owl"

type Menu struct {
	ID       string       `gorm:"primarykey" json:"id"`
	Name     string       `gorm:"comment:菜单名称;" json:"name"` // 唯一键
	ParentID string       `gorm:"comment:父级菜单;" json:"parentId"`
	MenuType owl.MenuType `gorm:"comment:菜单类型;" json:"menuType"`
	Path     string       `gorm:"comment:菜单路径;" json:"path"`
	Rank     int          `gorm:"comment:菜单排序;" json:"sort"`
	Meta     string       `gorm:"comment:菜单描述;" json:"meta"`
	Level    int          `gorm:"comment:菜单层级;" json:"level"`
}

func (i *Menu) TableName() string {
	return "admin_menu"
}
