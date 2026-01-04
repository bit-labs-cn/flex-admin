package model

type Area struct {
	AreaID   uint   `gorm:"column:area_id;comment:区域ID" json:"areaId,string"`       // 区域ID
	ParentID uint   `gorm:"column:parent_id;comment:父级区域ID" json:"parentId,string"` // 父级区域ID
	Name     string `gorm:"column:name;comment:区域名称;size:128" json:"name"`          // 区域名称
}

func (Area) TableName() string {
	return "admin_area"
}
