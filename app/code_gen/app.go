package code_gen

import (
	"github.com/gin-gonic/gin"
	"github.com/guoliang1994/gin-flex-admin/owl"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var _ owl.SubApp = (*SubAppCodeGen)(nil)

type SubAppCodeGen struct {
}

func (s SubAppCodeGen) Name() string {
	return "code_gen"
}

func (s SubAppCodeGen) Migrate(db *gorm.DB) {
	//TODO implement me
}

func (s SubAppCodeGen) Seed(db *gorm.DB) {
	//TODO implement me
}

func (s SubAppCodeGen) RegisterRouter(r *gin.Engine) {

}

func (s SubAppCodeGen) RegisterMenu(manager *owl.MenuManager) {
	manager.AddMenu(menu...)
	manager.AddMenu(device)
}

func (s SubAppCodeGen) RegisterCommand(command *cobra.Command) {

}
