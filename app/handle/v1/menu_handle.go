package v1

import (
	"bit-labs.cn/owl"
	"github.com/gin-gonic/gin"
	"net/http"
)

var _ owl.Handler = (*MenuHandle)(nil)

type MenuHandle struct {
	menuRepo *owl.MenuRepository
}

func NewMenuHandle(menuManager *owl.MenuRepository) *MenuHandle {
	return &MenuHandle{menuRepo: menuManager}
}
func (i MenuHandle) ModuleName() (en string, zh string) {
	return "menu", "菜单管理"
}

// Assignable
// @Tags      Admin
// @Summary   获取所有菜单，包括按钮，用于权限分配
// @Security  JWT
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  owl.Menu{data=[]owl.Menu}
// @Router    /api/v1/menus/assignable [get]
func (i MenuHandle) Assignable(c *gin.Context) {
	menus := i.menuRepo.GetMenusWithBtn()
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"msg":     "获取menus成功",
		"data":    menus,
	})
	return
}

// Menus
// @Tags      Admin
// @Summary   获取所有菜单,不包含按钮，用于前端渲染菜单
// @Security  JWT
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  owl.Menu{data=[]owl.Menu}
// @Router    /api/v1/menus [get]
func (i MenuHandle) Menus(c *gin.Context) {
	menus := i.menuRepo.GetMenusWithBtn()
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"msg":     "获取menus成功",
		"data":    menus,
	})
	return
}
