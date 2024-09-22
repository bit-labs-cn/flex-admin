package model

import (
	"time"
)

type ModelUser struct {
	ModelBase
	Avatar      *string    `gorm:"comment:用户头像" json:"avatar"`
	Name        string     `gorm:"comment:用户名称" json:"name"`
	NickName    string     `gorm:"comment:用户昵称" json:"nickName"`
	Account     string     `gorm:"comment:用户账号" json:"account"`
	Password    string     `gorm:"comment:用户密码" json:"password"`
	Description *string    `gorm:"comment:用户描述" json:"description"`
	Phone       *string    `gorm:"comment:手机" json:"phone"`
	Email       *string    `gorm:"comment:邮箱" json:"email"`
	Status      int        `gorm:"comment:状态" json:"status"`
	VerifiedAt  *time.Time `gorm:"comment:验证时间" json:"verified_at"`
}

func (i *ModelUser) TableName() string {
	return "admin_user"
}
