package model

import "bit-labs.cn/owl/provider/db"

type Dict struct {
	db.BaseModel
	Name   string `gorm:"column:name;comment:字典名（中）" json:"name"`
	Type   string `gorm:"column:type;comment:字典名（英）" json:"type"`
	Status uint8  `gorm:"column:status;comment:状态" json:"status,string"`
	Desc   string `gorm:"column:desc;comment:描述" json:"desc"`
	Sort   uint8  `gorm:"column:sort;comment:排序" json:"sort,string"`
}

func (i Dict) TableName() string {
	return "admin_dict"
}

type DictItem struct {
	db.BaseModel
	Label    string `gorm:"column:label;comment:展示值" json:"label"`
	Value    string `gorm:"column:value;comment:字典值" json:"value"`
	Extend   string `gorm:"column:extend;comment:扩展值" json:"extend"`
	Status   uint8  `gorm:"column:status;comment:启用状态" json:"status,string"`
	Sort     uint   `gorm:"column:sort;comment:排序标记" json:"sort,string"`
	DictType string `gorm:"column:dict_type;comment:冗余" json:"dictType"`
	DictID   uint   `gorm:"column:dict_id;comment:关联标记" json:"dictID,string"`
}

func (i *DictItem) TableName() string {
	return "admin_dict_item"
}
