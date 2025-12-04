package model

import (
	"bit-labs.cn/owl/provider/db"
	"bit-labs.cn/owl/provider/router"
)

type Api struct {
	db.BaseModel
	Name        string             `json:"name"`
	Code        string             `json:"code"`
	Path        string             `gorm:"comment:api路径" json:"path"`
	Module      string             `gorm:"comment:模块名称" json:"module"`
	Group       string             `json:"group" gorm:"type:varchar(255);comment:api组"`
	Method      string             `gorm:"type:varchar(255);comment:方法" json:"method"`
	Description string             `gorm:"type:varchar(255);comment:api中文描述" json:"description"`
	AccessLevel router.AccessLevel `json:"accessLevel"`
}

func (i *Api) TableName() string {
	return "admin_api"
}
