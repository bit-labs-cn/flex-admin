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
		var extra = []owl.RouterInfo{
			{"创建用户", module, PermissionAddUser, "", owl.Authorized},
			{"更新用户", module, PermissionUpdateUser, "", owl.Authorized},
			{"删除用户", module, PermissionDeleteUser, "", owl.Authorized},
			{"搜索用户", module, PermissionRetrieveUser, "", owl.Authorized},
			{"用户详情", module, PermissionDetailUser, "", owl.Authorized},
			{"用户登录", module, "", "", owl.Public},
			{"修改密码", module, PermissionChangeUserPassword, "", owl.Authenticated},
		}

		handle := v1.NewUserHandle()

		gv1.POST(&gin.RouteInfo{Path: "/user", Extra: extra[0]}, handle.Create)
		gv1.DELETE(&gin.RouteInfo{Path: "/user/:id", Extra: extra[1]}, handle.Delete)
		gv1.PUT(&gin.RouteInfo{Path: "/user/:id", Extra: extra[2]}, handle.Update)
		gv1.GET(&gin.RouteInfo{Path: "/user", Extra: extra[3]}, handle.Retrieve)
		gv1.GET(&gin.RouteInfo{Path: "/user/:id", Extra: extra[4]}, handle.Detail)
		gv1.PUT(&gin.RouteInfo{Path: "/user/changePassword", Extra: extra[5]}, handle.Detail)
		gv1.PUT(&gin.RouteInfo{Path: "/user/resetPassword", Extra: extra[6]}, handle.Detail)
	}

	// role
	{
		handle := v1.RoleHandle{}
		module := "角色管理"
		var extra = []owl.RouterInfo{
			{"创建角色", module, "admin:role:add", "", owl.Authorized},
			{"更新角色", module, "admin:role:update", "", owl.Authorized},
			{"删除角色", module, "admin:role:delete", "", owl.Authorized},
			{"搜索角色", module, "admin:role:retrieve", "", owl.Authorized},
			{"角色详情", module, "admin:role:detail", "", owl.Authorized},
			{"查询角色菜单", module, "admin:role:menuIds", "", owl.Authorized},
			{"分配菜单", module, "admin:role:assignMenu", "", owl.Authorized},
		}

		gv1.POST(&gin.RouteInfo{Path: "/role", Extra: extra[0]})
		gv1.DELETE(&gin.RouteInfo{Path: "/role/:id", Extra: extra[1]})
		gv1.PUT(&gin.RouteInfo{Path: "/role/:id", Extra: extra[2]})
		gv1.GET(&gin.RouteInfo{Path: "/role", Extra: extra[3]}, handle.Retrieve)
		gv1.GET(&gin.RouteInfo{Path: "/role/:id", Extra: extra[4]})
		gv1.GET(&gin.RouteInfo{Path: "/role-menu-ids/:id", Extra: extra[5]}, handle.MenuIds)
		gv1.POST(&gin.RouteInfo{Path: "/role/assign-menu", Extra: extra[6]}, handle.AssignMenu)
	}

	// dict
	{
		module := "字典管理"
		var extra = []owl.RouterInfo{
			{"创建字典", module, "admin:role:add", "", owl.Authorized},
			{"更新字典", module, "admin:role:update", "", owl.Authorized},
			{"删除字典", module, "admin:role:delete", "", owl.Authorized},
			{"搜索字典", module, "admin:role:retrieve", "", owl.Authorized},
			{"字典详情", module, "admin:role:detail", "", owl.Authorized},
		}

		gv1.POST(&gin.RouteInfo{Path: "/dict", Extra: extra[0]})
		gv1.DELETE(&gin.RouteInfo{Path: "/dict/:id", Extra: extra[1]})
		gv1.PUT(&gin.RouteInfo{Path: "/dict/:id", Extra: extra[2]})
		gv1.GET(&gin.RouteInfo{Path: "/dict", Extra: extra[3]})
		gv1.GET(&gin.RouteInfo{Path: "/dict/:id", Extra: extra[4]})
	}

	// api(permission)
	{
		module := "api管理"
		var extra = []owl.RouterInfo{
			{"查询api", module, "", "", owl.Authenticated},
		}
		gv1.GET(&gin.RouteInfo{Path: "/api", Extra: extra[0]}, func(c *gin.Context) {
			routers := r.GetAllRoutes()
			c.JSON(200, gin.H{
				"success": true,
				"msg":     "获取api成功",
				"data":    routers,
			})
			return
		})
	}

	// menu
	{
		module := "菜单管理"
		gv1.GET(&gin.RouteInfo{Path: "/assign-menu", Extra: owl.RouterInfo{Name: "查询菜单", Module: module, AccessLevel: owl.Authorized}}, func(c *gin.Context) {
			menus := owl.MenuMange.GetMenus()
			for _, m2 := range menus {
				iter(m2, 1)
			}
			c.JSON(200, gin.H{
				"success": true,
				"msg":     "获取menus成功",
				"data":    menus,
			})
			return
		})
		gv1.GET(&gin.RouteInfo{Path: "/menu", Extra: owl.RouterInfo{Name: "查询菜单", Module: module, AccessLevel: owl.Authorized}}, func(c *gin.Context) {
			menus := owl.MenuMange.GetMenus()
			for _, m2 := range menus {
				iter2(m2, 1)
			}
			c.JSON(200, gin.H{
				"success": true,
				"msg":     "获取menus成功",
				"data":    menus,
			})
			return
		})
	}
}

func iter(menu *owl.Menu, level int) {
	if level == 1 {
		menu.Ancestors = menu.Name
	}
	if menu.Children != nil && len(menu.Children) > 0 {
		for _, v := range menu.Children {
			v.Ancestors = menu.Ancestors + "," + v.Name
			v.ParentName = menu.Name
			iter(v, level+1)
		}
	}
}

func iter2(menu *owl.Menu, level int) owl.Menu {
	if level == 1 {
		menu.Ancestors = menu.Name
	}
	if menu.Children != nil && len(menu.Children) > 0 {
		// 创建一个新的切片来保存过滤后的子节点
		filteredChildren := make([]*owl.Menu, 0)
		for _, v := range menu.Children {
			if v.MenuType == owl.MenuTypeBtn {
				continue
			}
			v.Ancestors = menu.Ancestors + "," + v.Name
			v.ParentName = menu.Name

			iter2(v, level+1)

			filteredChildren = append(filteredChildren, v)

		}
		menu.Children = filteredChildren
	}
	return *menu
}
