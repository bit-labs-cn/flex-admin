package model

import (
	"time"
)

type ModelUser struct {
	ModelBase
	Avatar     *string     `gorm:"comment:用户头像" json:"avatar"`
	Username   string      `gorm:"comment:用户名称" json:"username"`
	Nickname   string      `gorm:"comment:用户昵称" json:"nickname"`
	Password   string      `gorm:"comment:用户密码" json:"-"`
	Remark     *string     `gorm:"comment:remark" json:"remark"`
	Phone      *string     `gorm:"comment:手机" json:"phone"`
	Email      *string     `gorm:"comment:邮箱" json:"email"`
	Status     int         `gorm:"comment:状态" json:"status"`
	Sex        int         `gorm:"comment:性别" json:"sex"`
	VerifiedAt *time.Time  `gorm:"comment:验证时间" json:"verified_at"`
	Roles      []ModelRole `gorm:"many2many:user_role;joinForeignKey:user_id;References:id;JoinReferences:role_id" json:"roles"`
	Menus      []ModelMenu `gorm:"many2many:user_menu;joinForeignKey:user_id;References:id;JoinReferences:menu_id" json:"menus"`

	IsSuperAdmin bool     `json:"isSuperAdmin" gorm:"-"`
	Permissions  []string `json:"permissions" gorm:"-"`
}

func (i ModelUser) TableName() string {
	return "admin_user"
}
