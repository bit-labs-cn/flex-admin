package admin

import (
	"github.com/guoliang1994/gin-flex-admin/app/admin/handle/v1"
	"github.com/guoliang1994/gin-flex-admin/app/admin/service"
	"github.com/guoliang1994/gin-flex-admin/owl"
	"github.com/guoliang1994/gin-flex-admin/owl/middleware"
)

var userMenu, roleMenu, menuMenu, apiMenu *owl.Menu

func InitApi(app *owl.Application, appName string) {
	gv1 := app.Engine().Group("/api/v1").Use(middleware.Cors(), middleware.PermissionCheck(app))

	userHandle := v1.NewUserHandle(service.NewUserService(app, service.NewRoleService(app.DB(), app)), service.NewRoleService(app.DB(), app), app)
	roleHandle := v1.NewRoleHandle(service.NewRoleService(app.DB(), app))
	apiHandle := v1.NewApiHandle(app)
	menuHandle := v1.NewMenuHandle(app)

	// user
	{
		r := owl.NewRouterInfoBuilder(appName, &userHandle, gv1, owl.MenuOption{
			ComponentName: "SystemUser",
			Path:          "/system/user/index",
			Icon:          "ep:user",
		})

		r.Post("/users/login", owl.Public, userHandle.Login).Name("用户登录").Build()

		r.Put("/users/me/password", owl.Authenticated, userHandle.ChangePassword).Name("修改我的密码").Build()
		r.Get("/users/me/menus", owl.Authenticated, userHandle.GetMyMenus).Name("我的菜单").Build()
		r.Get("/users/me", owl.Authenticated, userHandle.GetMyMenus).Name("我的信息").Build()

		r.Post("/users", owl.Authorized, userHandle.Create).Name("创建用户").Build()
		r.Delete("/users/:id", owl.Authorized, userHandle.Delete).Name("删除用户").Build()

		deps := []owl.Dep{
			{
				Handler: &userHandle,
				Method:  userHandle.Detail,
			},
		}
		r.Put("/users/:id", owl.Authorized, userHandle.Update).Deps(deps...).Name("更新用户").Build()

		r.Put("/users/:id/status", owl.Authorized, userHandle.ChangeStatus).Name("启用，禁用用户").Build()
		r.Get("/users", owl.Authorized, userHandle.Retrieve).Name("分页获取用户").Build()
		r.Get("/users/:id", owl.Authorized, userHandle.Detail).Name("获取用户详情").Build()
		r.Put("/users/:id/reset", owl.Authorized, userHandle.ResetPassword).Name("重置用户密码").Build()

		deps = []owl.Dep{
			{
				Handler: &userHandle,
				Method:  userHandle.GetRoleIdsByUserId,
			},
			{
				Handler: &roleHandle,
				Method:  roleHandle.GetRolesSimple,
			},
		}
		r.Post("/users/:id/roles", owl.Authorized, userHandle.AssignRolesToUser).Deps(deps...).Name("分配角色给用户").Build()
		r.Get("/users/:id/roles", owl.Authorized, userHandle.GetRoleIdsByUserId).Name("获取用户角色").Build()

		deps = []owl.Dep{
			{
				Handler: menuHandle,
				Method:  menuHandle.GetMenu,
			},
		}
		r.Post("/users/:id/menus", owl.Authorized, userHandle.AssignMenuToUser).Deps(deps...).Name("分配菜单给用户").Build()

		userMenu = r.GetMenu()
	}

	// role
	{
		router := owl.NewRouterInfoBuilder(appName, &roleHandle, gv1, owl.MenuOption{
			ComponentName: "SystemRole",
			Path:          "/system/role/index",
			Icon:          "fa-solid:users",
		})

		router.Post("/roles", owl.Authorized, roleHandle.Create).Name("创建角色").Build()
		router.Delete("/roles/:id", owl.Authorized, roleHandle.Delete).Name("删除角色").Build()
		router.Put("/roles/:id", owl.Authorized, roleHandle.Update).Name("更新角色").Build()
		router.Put("/roles/:id/status", owl.Authorized, roleHandle.ChangeStatus).Name("禁用，启用角色").Build()
		router.Get("/roles", owl.Authorized, roleHandle.Retrieve).Name("角色列表").Build()
		router.Get("/roles/simple", owl.Authenticated, roleHandle.GetRolesSimple).Name("所有角色(id,name)").Build()
		router.Get("/roles/:id", owl.Authorized, roleHandle.Detail).Name("获取角色详情").Build()
		router.Get("/roles/:id/menu-ids", owl.Authorized, roleHandle.GetRoleMenuIDs).Name("获取角色拥有的菜单").Build()
		router.Put("/roles/:id/menus", owl.Authorized, roleHandle.AssignMenusToRole).Name("分配菜单给角色").Build()
		roleMenu = router.GetMenu()
	}

	// api(permission)
	{
		router := owl.NewRouterInfoBuilder(appName, apiHandle, gv1, owl.MenuOption{
			ComponentName: "SystemApi",
			Path:          "/system/api/index",
			Icon:          "ep:user",
		})
		router.Get("/api", owl.Authorized, apiHandle.GetAll).Name("查询所有接口").Build()

		apiMenu = router.GetMenu()
	}

	// menu
	{
		router := owl.NewRouterInfoBuilder(appName, menuHandle, gv1, owl.MenuOption{
			ComponentName: "SystemMenu",
			Path:          "/system/menu/index",
			Icon:          "ep:menu",
		})

		router.Get("/menus/assignable", owl.Authorized, menuHandle.GetMenu).Name("查询可分配的菜单").Description("查询可分配的菜单（包含按钮）").Build()
		menuMenu = router.GetMenu()
	}
}
