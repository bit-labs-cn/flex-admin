package model

type ModelRole struct {
	ModelBase
	Status      int    `json:"status"`
	Name        string `gorm:"comment:角色名称" json:"name"`
	Code        string `gorm:"comment:角色编码" json:"code"`
	Sort        string `json:"sort"`
	Description string `gorm:"comment:角色描述" json:"description"`
}

func (i *ModelRole) TableName() string {
	return "admin_role"
}
