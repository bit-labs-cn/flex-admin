package model

// ModelUserMenu 用户与菜单关联，直接给用户赋予权限（一般用于附加权限）
type ModelUserMenu struct {
	ModelBase
	UserId uint
	MenuId string
}

func (i *ModelUserMenu) TableName() string {
	return "admin_user_menu"
}
