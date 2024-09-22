package model

// ModelMenuButton 按钮, 按理来说应该是一张表，但是采用json直接存储到菜单中
type ModelMenuButton struct {
	Name             string   `json:"name"`
	ShowKey          string   `json:"showKey"` // 显示此按钮需要的key
	HasApiPermission []string `json:"hasApiPermission"`
}

type ModelMenu struct {
	ModelBase

	Name       string `gorm:"comment:菜单名称;" json:"id"` // 唯一键
	ParentName string `gorm:"comment:父级菜单;" json:"parentId"`

	Path string `gorm:"comment:菜单路径;" json:"path"`
	Rank int    `gorm:"comment:菜单排序;" json:"sort"`
	Meta string `gorm:"comment:菜单描述;" json:"meta"`

	Ancestors string `gorm:"comment:祖先;" json:"ancestors"`
	Level     int    `gorm:"comment:菜单层级;" json:"level"`
}

func (i *ModelMenu) TableName() string {
	return "admin_menu"
}
