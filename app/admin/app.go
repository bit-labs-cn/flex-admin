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

func (i *SubAppAdmin) Migrate(db *gorm.DB) {
	model.Migrate(db)
}

func (i *SubAppAdmin) Seed(db *gorm.DB) {

}
func (i *SubAppAdmin) RegisterRouter(r *gin.Engine) {
	router.InitApi(r)
}

func (i *SubAppAdmin) RegisterMenu(manager *owl.MenuManager) {
	manager.AddMenu(menu...)
}

func (i *SubAppAdmin) RegisterCommand(command *cobra.Command) {
	command.AddCommand(cmd.Version, cmd.GenMigrate)
}
