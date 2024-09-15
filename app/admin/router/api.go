package router

import (
	"github.com/gin-gonic/gin"
	"github.com/guoliang1994/gin-flex-admin/app/admin/handle/v1"
	"github.com/guoliang1994/gin-flex-admin/owl"
	"github.com/guoliang1994/gin-flex-admin/owl/middleware"
)

func InitApi(r *gin.Engine) {
	gv1 := r.Group("/api/v1").Use(middleware.Cors())

	// user
	{
		module := "用户管理"
		handle := v1.NewUserHandle()

		gv1.POST(&gin.RouteInfo{Path: "/user", Extra: owl.RouterInfo{Name: "创建用户", Module: module, AccessLevel: owl.Authorized}}, handle.Create)
		gv1.DELETE(&gin.RouteInfo{Path: "/user/:id", Extra: owl.RouterInfo{Name: "更新用户", Module: module, AccessLevel: owl.Authorized}}, handle.Delete)
		gv1.PUT(&gin.RouteInfo{Path: "/user/:id", Extra: owl.RouterInfo{Name: "删除用户", Module: module, AccessLevel: owl.Authorized}}, handle.Update)
		gv1.GET(&gin.RouteInfo{Path: "/user", Extra: owl.RouterInfo{Name: "搜索用户", Module: module, AccessLevel: owl.Authorized}}, handle.Retrieve)
		gv1.GET(&gin.RouteInfo{Path: "/user/:id", Extra: owl.RouterInfo{Name: "用户详情", Module: module, AccessLevel: owl.Authorized}}, handle.Detail)
	}

	// role
	{
		module := "角色管理"
		handle := v1.NewRoleHandle()
		gv1.POST(&gin.RouteInfo{Path: "/role", Extra: owl.RouterInfo{Name: "创建角色", Module: module, AccessLevel: owl.Authorized}}, handle.Create)
		gv1.DELETE(&gin.RouteInfo{Path: "/role/:id", Extra: owl.RouterInfo{Name: "更新角色", Module: module, AccessLevel: owl.Authorized}}, handle.Delete)
		gv1.PUT(&gin.RouteInfo{Path: "/role/:id", Extra: owl.RouterInfo{Name: "删除角色", Module: module, AccessLevel: owl.Authorized}}, handle.Update)
		gv1.GET(&gin.RouteInfo{Path: "/role", Extra: owl.RouterInfo{Name: "查询角色", Module: module, AccessLevel: owl.Authorized}}, handle.Retrieve)
		gv1.GET(&gin.RouteInfo{Path: "/role/:id", Extra: owl.RouterInfo{Name: "查询角色", Module: module, AccessLevel: owl.Authorized}}, handle.Detail)
	}

	// api(permission)
	{
		module := "api管理"
		gv1.GET(&gin.RouteInfo{Path: "/api", Extra: owl.RouterInfo{Name: "查询api", Module: module, AccessLevel: owl.Authorized}}, func(c *gin.Context) {
			routers := r.GetAllRoutes()
			c.JSON(200, gin.H{
				"msg":  "获取api成功",
				"data": routers,
			})
			return
		})
	}

	// menu
	{
		module := "菜单管理"
		gv1.GET(&gin.RouteInfo{Path: "/menu", Extra: owl.RouterInfo{Name: "查询菜单", Module: module, AccessLevel: owl.Authorized}}, func(c *gin.Context) {
			menus := owl.MM.GetMenus()
			c.JSON(200, gin.H{
				"msg":  "获取menus成功",
				"data": menus,
			})
			return
		})
	}
}
