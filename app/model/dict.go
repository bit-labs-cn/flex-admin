package model

type Dict struct {
	Base
	Name   string `json:"name"  gorm:"column:name;comment:字典名（中）"`        // 字典名（中）
	Type   string `json:"type"  gorm:"column:type;comment:字典名（英）"`        // 字典名（英）
	Status uint8  `json:"status,string"  gorm:"column:status;comment:状态"` // 状态
	Desc   string `json:"desc"  gorm:"column:desc;comment:描述"`            // 描述
	Sort   uint8  `json:"sort,string"  gorm:"column:sort;comment:排序"`     // 排序
}

func (i Dict) TableName() string {
	return "admin_dict"
}

type DictItem struct {
	Base
	Label    string `json:"label"  gorm:"column:label;comment:展示值"`            // 展示值
	Value    string `json:"value"  gorm:"column:value;comment:字典值"`            // 字典值
	Extend   string `json:"extend"  gorm:"column:extend;comment:扩展值"`          // 扩展值(分组)
	Status   uint8  `json:"status,string"  gorm:"column:status;comment:启用状态"`  // 启用状态
	Sort     uint   `json:"sort,string"  gorm:"column:sort;comment:排序标记"`      // 排序标记
	DictType string `json:"dictType"  gorm:"column:dict_type;comment:冗余"`      // 冗余
	DictID   uint   `json:"dictID,string"  gorm:"column:dict_id;comment:关联标记"` // 关联标记
}

func (i *DictItem) TableName() string {
	return "admin_dict_item"
}
