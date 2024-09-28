package code_gen

import (
	"github.com/guoliang1994/gin-flex-admin/app/cms/router"
	"github.com/guoliang1994/gin-flex-admin/owl"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var _ owl.SubApp = (*SubAppCms)(nil)

type SubAppCms struct {
	app *owl.Application
}

func (i SubAppCms) RegisterRouter() {
	router.InitApi(i.app, i.Name())
}

func (i SubAppCms) RegisterMenu() {
	i.app.MenuManager().AddMenu(router.InitMenu())
}

func (i SubAppCms) Name() string {
	return "cms"
}

func (i SubAppCms) Construct(application *owl.Application) owl.SubApp {
	return SubAppCms{
		app: application,
	}
}

func (i SubAppCms) Migrate(db *gorm.DB) {
	//TODO implement me
}

func (i SubAppCms) Seed(db *gorm.DB) {
	//TODO implement me
}

func (i SubAppCms) RegisterCommand(command *cobra.Command) {

}
