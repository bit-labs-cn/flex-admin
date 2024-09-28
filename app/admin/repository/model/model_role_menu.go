package model

// ModelRoleMenu 用户和角色的菜单
type ModelRoleMenu struct {
	RoleID uint   `json:"roleID" gorm:"comment:角色id;index"`
	MenuID string `json:"menuID" gorm:"comment:菜单id;index"`
}

func (i *ModelRoleMenu) TableName() string {
	return "admin_role_menu"
}
