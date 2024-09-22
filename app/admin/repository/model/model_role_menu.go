package model

// ModelRoleMenu 角色与菜单关联表, 菜单映射到权限
type ModelRoleMenu struct {
	ModelBase
	RoleId string `json:"roleId"`
	MenuId string `json:"menuId"`
}

func (i *ModelRoleMenu) TableName() string {
	return "admin_role_menu"
}
