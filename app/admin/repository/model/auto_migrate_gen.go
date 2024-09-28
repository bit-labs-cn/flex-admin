// generate by auto_migrate Do not edit it
package model

import "gorm.io/gorm"

func Migrate(db *gorm.DB) {
	_ = db.Migrator().AutoMigrate(

		&ModelDictionary{},

		&ModelDictionaryDetail{},

		&ModelApi{},

		&ModelMenu{},
		&ModelRole{},
		&ModelUser{},
		&ModelRoleMenu{},
		&ModelUserMenu{},
	)
}
