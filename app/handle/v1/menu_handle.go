package v1

import (
	"bit-labs.cn/owl/provider/router"
	"github.com/gin-gonic/gin"
)

var _ router.Handler = (*MenuHandle)(nil)

type MenuHandle struct {
	menuRepo *router.MenuRepository
}

func NewMenuHandle(menuManager *router.MenuRepository) *MenuHandle {
	return &MenuHandle{menuRepo: menuManager}
}
func (i MenuHandle) ModuleName() (en string, zh string) {
	return "menu", "菜单管理"
}

// @Summary		获取可分配菜单
// @Description	获取所有菜单（含按钮）用于权限分配
// @Tags			菜单管理
// @Produce		json
// @Success		200	{object}	router.Resp	"操作成功"
// @Router			/api/v1/menus/assignable [GET]
func (i MenuHandle) Assignable(c *gin.Context) {
	menus := i.menuRepo.GetMenusWithBtn()
	router.Success(c, menus)
	return
}

// @Summary		获取菜单列表
// @Description	获取所有菜单（不含按钮）用于前端渲染
// @Tags			菜单管理
// @Produce		json
// @Success		200	{object}	router.Resp	"操作成功"
// @Router			/api/v1/menus [GET]
func (i MenuHandle) Menus(c *gin.Context) {
	menus := i.menuRepo.GetMenusWithBtn()
	router.Success(c, menus)
	return
}
