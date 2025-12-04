package model

import (
	"bit-labs.cn/owl/provider/db"
	"gorm.io/gorm/schema"
)

var _ schema.Tabler = (*Position)(nil)

type Position struct {
	db.BaseModel
	Name   string `gorm:"comment:岗位名称" json:"name"`
	Remark string `gorm:"comment:岗位备注" json:"remark"`
	Status int    `gorm:"comment:状态(1启用,2禁用)" json:"status"`
}

func (Position) TableName() string {
	return "admin_position"
}

func (p *Position) Enable()  { p.Status = 1 }
func (p *Position) Disable() { p.Status = 0 }
