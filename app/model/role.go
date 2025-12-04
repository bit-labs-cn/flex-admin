package model

import (
	"bit-labs.cn/owl/provider/db"
	"gorm.io/gorm/schema"
)

var _ schema.Tabler = (*Role)(nil)
var _ schema.Tabler = (*RoleMenu)(nil)

type Role struct {
	db.BaseModel
	Status int    `gorm:"comment:角色状态" json:"status"`
	Name   string `gorm:"comment:角色名称;type:string;size:128" json:"name"`
	Code   string `gorm:"comment:角色编码;type:string;size:64" json:"code"`
	Remark string `gorm:"comment:角色描述" json:"remark"`

	Menus []Menu `gorm:"many2many:admin_role_menu;joinForeignKey:role_id;References:id;JoinReferences:menu_id" json:"menus"`
}

func (i *Role) TableName() string {
	return "admin_role"
}
func (i *Role) SetMenus(menus []Menu) {
	i.Menus = menus
}

// RoleMenu 用户和角色的菜单
type RoleMenu struct {
	RoleID uint   `json:"roleID" gorm:"comment:角色id;index"`
	MenuID string `json:"menuID" gorm:"comment:菜单id;index"`
}

func (i *RoleMenu) TableName() string {
	return "admin_role_menu"
}
