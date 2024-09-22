package model

type ModelUserRole struct {
	ModelBase
	UserId uint
	RoleId uint
}

func (i *ModelUserRole) TableName() string {
	return "admin_user_role"
}
