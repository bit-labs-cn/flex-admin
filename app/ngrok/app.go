package ngrok

import (
	"github.com/gin-gonic/gin"
	"github.com/guoliang1994/gin-flex-admin/owl"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

type SubAppNgrok struct {
}

func (i SubAppNgrok) Name() string {
	return "admin"
}

func (i SubAppNgrok) Migrate(db *gorm.DB) {
	//TODO implement me
}

func (i SubAppNgrok) Seed(db *gorm.DB) {
	//TODO implement me
}

func (i SubAppNgrok) RegisterRouter(r *gin.Engine) {

}

func (i SubAppNgrok) RegisterMenu(manager *owl.MenuManager) {
	manager.AddMenu(menu...)
}

func (i SubAppNgrok) RegisterCommand(command *cobra.Command) {

}
