package model

// ModelRoleMenu 用户和角色的菜单
type ModelUserMenu struct {
	UserID uint   `json:"userID" gorm:"comment:角色id;index"`
	MenuID string `json:"menuID" gorm:"comment:菜单id;index"`
}

func (i *ModelUserMenu) TableName() string {
	return "admin_user_menu"
}
