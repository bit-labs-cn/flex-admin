package route

import (
	v1 "bit-labs.cn/gin-flex-admin/app/handle/v1"
	middleware2 "bit-labs.cn/gin-flex-admin/app/middleware"
	"bit-labs.cn/owl"
	"bit-labs.cn/owl/contract/foundation"
	"bit-labs.cn/owl/contract/log"
	"bit-labs.cn/owl/middleware"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

var userMenu, roleMenu, menuMenu, apiMenu, deptMenu *owl.Menu

func InitMenu() *owl.Menu {
	return &owl.Menu{
		Path: "/system",
		Name: "System",
		Meta: owl.Meta{
			Title: "超级管理员",
			Icon:  "ep:lock",
		},
		MenuType: owl.MenuTypeDir,
		Children: []*owl.Menu{
			userMenu,
			roleMenu,
			deptMenu,
			menuMenu,
			apiMenu,
		},
	}
}

func InitApi(app foundation.Application, appName string) {

	err := app.Invoke(func(
		userHandle *v1.UserHandle,
		roleHandle *v1.RoleHandle,
		apiHandle *v1.ApiHandle,
		menuHandle *v1.MenuHandle,
		dictHandle *v1.DictHandle,
		deptHandle *v1.DeptHandle,
		enforcer casbin.IEnforcer,
		engine *gin.Engine,
		log log.Logger,
	) {

		gv1 := engine.Group("/api/v1").Use(middleware.Cors(), middleware2.PermissionCheck(engine, enforcer))

		// user
		{
			r := owl.NewRouteInfoBuilder(appName, userHandle, gv1, owl.MenuOption{
				ComponentName: "SystemUser",
				Path:          "/system/user/index",
				Icon:          "ep:user",
			})

			r.Post("/users/login", owl.Public, userHandle.Login).Name("用户登录").Build()

			r.Put("/users/me/password", owl.Authenticated, userHandle.ChangePassword).Name("修改我的密码").Build()
			r.Get("/users/me/menus", owl.Authenticated, userHandle.GetMyMenus).Name("我的菜单").Build()
			r.Get("/users/me", owl.Authenticated, userHandle.Me).Name("我的信息").Build()

			r.Post("/users", owl.Authorized, userHandle.Create).Name("创建用户").Build()
			r.Delete("/users/:id", owl.Authorized, userHandle.Delete).Name("删除用户").Build()

			r.Put("/users/:id", owl.Authorized, userHandle.Update).Deps(
				[]owl.Dep{
					{Handler: userHandle, Method: userHandle.Detail},
				}...,
			).Name("更新用户").Build()

			r.Put("/users/:id/status", owl.Authorized, userHandle.ChangeStatus).Name("启用，禁用用户").Build()

			r.Get("/users", owl.Authorized, userHandle.Retrieve).Deps(
				[]owl.Dep{
					{Handler: userHandle, Method: deptHandle.Retrieve},
				}...,
			).Name("分页获取用户").Build()

			r.Get("/users/:id", owl.Authorized, userHandle.Detail).Name("获取用户详情").Build()
			r.Put("/users/:id/reset", owl.Authorized, userHandle.ResetPassword).Name("重置用户密码").Build()

			r.Post("/users/:id/roles", owl.Authorized, userHandle.AssignRolesToUser).Deps(
				[]owl.Dep{
					{Handler: userHandle, Method: userHandle.GetRoleIdsByUserId},
					{Handler: roleHandle, Method: roleHandle.RoleOptions},
				}...,
			).Name("分配角色给用户").Build()

			r.Get("/users/:id/roles", owl.Authorized, userHandle.GetRoleIdsByUserId).Name("获取用户角色").Build()

			userMenu = r.GetMenu()
		}

		// role
		{
			router := owl.NewRouteInfoBuilder(appName, roleHandle, gv1, owl.MenuOption{
				ComponentName: "SystemRole",
				Path:          "/system/role/index",
				Icon:          "fa-solid:users",
			})

			router.Post("/roles", owl.Authorized, roleHandle.Create).Name("创建角色").Build()
			router.Delete("/roles/:id", owl.Authorized, roleHandle.Delete).Name("删除角色").Build()
			router.Put("/roles/:id", owl.Authorized, roleHandle.Update).Name("更新角色").Build()
			router.Put("/roles/:id/status", owl.Authorized, roleHandle.ChangeStatus).Name("禁用，启用角色").Build()

			router.Get("/roles", owl.Authorized, roleHandle.Retrieve).Name("角色列表").Build()
			router.Get("/roles/:id", owl.Authorized, roleHandle.Detail).Name("获取角色详情").Build()
			router.Get("/roles/:id/menu-ids", owl.Authorized, roleHandle.GetRoleMenuIDs).Name("获取角色拥有的菜单").Build()
			router.Get("/roles/:id/users", owl.Authorized, roleHandle.GetRoleMenuIDs).Name("获取角色下的用户").Build()

			router.Get("/roles-options", owl.Authenticated, roleHandle.RoleOptions).Name("所有角色(id,name)").Build()

			deps := []owl.Dep{
				{Handler: roleHandle, Method: roleHandle.GetRoleMenuIDs},
				{Handler: menuHandle, Method: menuHandle.Assignable},
			}
			router.Put("/roles/:id/menus", owl.Authorized, roleHandle.AssignMenusToRole).Deps(deps...).Name("分配菜单给角色").Build()
			roleMenu = router.GetMenu()
		}

		// api(permission)
		{
			router := owl.NewRouteInfoBuilder(appName, apiHandle, gv1, owl.MenuOption{
				ComponentName: "SystemApi",
				Path:          "/system/api/index",
				Icon:          "ep:user",
			})
			router.Get("/api", owl.Authorized, apiHandle.GetAll).Name("查询所有接口").Build()

			apiMenu = router.GetMenu()
		}

		// menu
		{
			router := owl.NewRouteInfoBuilder(appName, menuHandle, gv1, owl.MenuOption{
				ComponentName: "SystemMenu",
				Path:          "/system/menu/index",
				Icon:          "ep:menu",
			})

			router.Get("/menus/assignable", owl.Authorized, menuHandle.Assignable).Name("查询可分配的菜单").Description("查询可分配的菜单（包含按钮）").Build()
			router.Get("/menus", owl.Authorized, menuHandle.Menus).Name("菜单列表").Build()

			menuMenu = router.GetMenu()
		}

		// dictionary
		{
			router := owl.NewRouteInfoBuilder(appName, dictHandle, gv1, owl.MenuOption{
				ComponentName: "SystemDict",
				Path:          "/system/dict/index",
				Icon:          "ep:menu",
			})

			router.Post("/dict", owl.AdminOnly, dictHandle.Create).Name("创建字典").Build()
			router.Delete("/dict/:id", owl.AdminOnly, dictHandle.Delete).Name("删除字典").Build()
			router.Put("/dict/:id", owl.AdminOnly, dictHandle.Update).Name("更新字典").Build()
			router.Get("/dict", owl.AdminOnly, dictHandle.Retrieve).Name("字典列表").Build()

			router.Post("/dict/:id/item", owl.AdminOnly, dictHandle.CreateItem).Name("新增字典项").Build()
			router.Put("/dict/:id/item/:itemID", owl.AdminOnly, dictHandle.UpdateItem).Name("更新字典项").Build()
			router.Get("/dict/:id/item", owl.AdminOnly, dictHandle.RetrieveItems).Name("获取字典列表").Build()
			router.Delete("/dict/:id/item/:itemID", owl.AdminOnly, dictHandle.DeleteItem).Name("删除字典项").Build()
		}

		// dept
		{
			router := owl.NewRouteInfoBuilder(appName, deptHandle, gv1, owl.MenuOption{
				ComponentName: "Dept",
				Path:          "/system/dept/index",
				Icon:          "ep:menu",
			})

			router.Post("/dept", owl.Authorized, deptHandle.Create).Name("新增部门").Build()
			router.Delete("/dept/:id", owl.Authorized, deptHandle.Delete).Name("删除部门").Build()
			router.Put("/dept/:id", owl.Authorized, deptHandle.Update).Name("更新部门").Build()
			router.Get("/dept", owl.Authorized, deptHandle.Retrieve).Name("获取部门列表").Build()
			router.Get("/dept/:id/users", owl.Authorized, roleHandle.GetRoleMenuIDs).Name("获取部门下的用户").Build()

			deptMenu = router.GetMenu()
		}
	})
	owl.PanicIf(err)
}
