package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/guoliang1994/gin-flex-admin/owl"
)

var _ owl.Handler = (*MenuHandle)(nil)

type MenuHandle struct {
	app *owl.Application
}

func NewMenuHandle(app *owl.Application) *MenuHandle {
	return &MenuHandle{app: app}
}
func (i MenuHandle) ModuleName() (en string, zh string) {
	return "menu", "菜单管理"
}

func (i MenuHandle) GetMenu(c *gin.Context) {
	menus := i.app.MenuManager().GetAllMenusWithBtn()
	c.JSON(200, gin.H{
		"success": true,
		"msg":     "获取menus成功",
		"data":    menus,
	})
	return
}
