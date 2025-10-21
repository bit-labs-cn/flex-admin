// generate by auto_migrate Do not edit it
package database

import (
	. "bit-labs.cn/flex-admin/app/model"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	_ = db.Migrator().AutoMigrate(

		&Dict{},
		&DictItem{},

		&Api{},

		&Menu{},
		&Role{},
		&User{},
		&RoleMenu{},
		&UserMenu{},
		&Dept{},
	)
}
