package code_gen

import (
	"github.com/gin-gonic/gin"
	"github.com/guoliang1994/gin-flex-admin/owl"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

type SubAppCms struct {
}

func (s SubAppCms) Migrate(db *gorm.DB) {
	//TODO implement me
}

func (s SubAppCms) Seed(db *gorm.DB) {
	//TODO implement me
}

func (s SubAppCms) RegisterRouter(r *gin.Engine) {

}

func (s SubAppCms) RegisterMenu(manager *owl.MenuManager) {
	manager.AddMenu(menu...)
}

func (s SubAppCms) RegisterCommand(command *cobra.Command) {

}
