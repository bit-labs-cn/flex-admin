package code_gen

import (
	"github.com/gin-gonic/gin"
	"github.com/guoliang1994/gin-flex-admin/owl"
	"github.com/spf13/cobra"
)

type SubAppCodeGen struct {
}

func (s SubAppCodeGen) RegisterRouter(r *gin.Engine) {

}

func (s SubAppCodeGen) RegisterMenu(manager *owl.MenuManager) {
	manager.AddMenu(menu...)
}

func (s SubAppCodeGen) RegisterCommand(command *cobra.Command) {

}
