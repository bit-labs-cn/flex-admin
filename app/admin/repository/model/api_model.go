package model

import (
	"github.com/guoliang1994/gin-flex-admin/owl"
	"gorm.io/gorm"
)

type ApiModel struct {
	gorm.Model
	Name        string          `json:"name"`
	Code        string          `json:"code"`
	Path        string          `gorm:"comment:api路径" json:"path"`
	Module      string          `gorm:"comment:模块名称" json:"module"`
	Group       string          `json:"group" gorm:"type:varchar(255);comment:api组"`
	Method      string          `gorm:"type:varchar(255);comment:方法" json:"method"`
	Description string          `gorm:"type:varchar(255);comment:api中文描述" json:"description"`
	AccessLevel owl.AccessLevel `json:"accessLevel"`
}

func (i *ApiModel) TableName() string {
	return "admin_api"
}
