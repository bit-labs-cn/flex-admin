package v1

import (
	"net/http"

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

// Assignable 获取可分配菜单
//
//	@Summary		获取可分配菜单
//	@Description	获取所有菜单，包括按钮，用于权限分配
//	@Tags			菜单管理
//	@Router			/api/v1/menus/assignable [GET]
//	@Access			AccessAuthorized
//	@Name			获取可分配菜单
//	@Success		200	{object}	router.RouterInfo{data=[]router.Menu}	"可分配菜单获取成功"
//	@Failure		400	{object}	router.RouterInfo						"请求参数错误"
//	@Failure		500	{object}	router.RouterInfo						"服务器内部错误"
func (i MenuHandle) Assignable(c *gin.Context) {
	menus := i.menuRepo.GetMenusWithBtn()
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"msg":     "获取menus成功",
		"data":    menus,
	})
	return
}

// Menus 获取菜单列表
//
//	@Summary		获取菜单列表
//	@Description	获取所有菜单，不包含按钮，用于前端渲染菜单
//	@Tags			菜单管理
//	@Router			/api/v1/menus [GET]
//	@Access			AccessAuthorized
//	@Name			获取菜单列表
//	@Success		200	{object}	router.RouterInfo{data=[]router.Menu}	"菜单列表获取成功"
//	@Failure		400	{object}	router.RouterInfo						"请求参数错误"
//	@Failure		500	{object}	router.RouterInfo						"服务器内部错误"
func (i MenuHandle) Menus(c *gin.Context) {
	menus := i.menuRepo.GetMenusWithBtn()
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"msg":     "获取menus成功",
		"data":    menus,
	})
	return
}
