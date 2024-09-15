package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/guoliang1994/gin-flex-admin/app/admin/cmd"
	"github.com/guoliang1994/gin-flex-admin/app/admin/repository/model"
	"github.com/guoliang1994/gin-flex-admin/app/admin/router"
	"github.com/guoliang1994/gin-flex-admin/owl"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

type SubAppAdmin struct {
}

func (s SubAppAdmin) Migrate(db *gorm.DB) {
	_ = db.Migrator().AutoMigrate(
		&model.ApiModel{},
		&model.UserModel{},
		&model.RoleModel{},
		&model.MenuModel{},
	)
}
func (s SubAppAdmin) RegisterRouter(r *gin.Engine) {
	router.InitApi(r)
}

func (s SubAppAdmin) RegisterMenu(manager *owl.MenuManager) {
	manager.AddMenu(menu...)
}

func (s SubAppAdmin) RegisterCommand(command *cobra.Command) {
	command.AddCommand(cmd.Version)
}
