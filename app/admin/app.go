package admin

import (
	"github.com/guoliang1994/gin-flex-admin/app/admin/cmd"
	"github.com/guoliang1994/gin-flex-admin/app/admin/repository/model"
	"github.com/guoliang1994/gin-flex-admin/owl"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var _ owl.SubApp = (*SubAppAdmin)(nil)

type SubAppAdmin struct {
	app *owl.Application
}

func (i *SubAppAdmin) Construct(application *owl.Application) owl.SubApp {
	return &SubAppAdmin{app: application}
}

func (i *SubAppAdmin) Name() string {
	return "admin"
}

func (i *SubAppAdmin) Migrate(db *gorm.DB) {
	model.Migrate(db)
}

func (i *SubAppAdmin) Seed(db *gorm.DB) {

}

func (i *SubAppAdmin) RegisterRouter() {
	InitApi(i.app, i.Name())
}

func (i *SubAppAdmin) RegisterMenu() {
	i.app.MenuManager().AddMenu(InitMenu())
}

func (i *SubAppAdmin) RegisterCommand(command *cobra.Command) {
	command.AddCommand(cmd.Version, cmd.GenMigrate)
}
