package model

import "bit-labs.cn/owl/provider/db"

type Dept struct {
	db.BaseModel
	Name        string  `gorm:"comment:部门名称;type:string;size:128" json:"name"`
	ParentId    int     `gorm:"comment:父级部门;type:string;size:64" json:"parentId,string"`
	Level       uint    `gorm:"comment:部门层级" json:"level"`
	Sort        uint    `gorm:"comment:排序" json:"sort"`
	Status      uint    `gorm:"comment:状态" json:"status"`
	Description string  `gorm:"comment:描述" json:"description"`
	Children    []*Dept `gorm:"foreignKey:parent_id" json:"children"`
}

func (Dept) TableName() string {
	return "admin_dept"
}
