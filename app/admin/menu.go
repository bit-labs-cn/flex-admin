package admin

import "github.com/guoliang1994/gin-flex-admin/owl"

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
			menuMenu,
			{
				Name: "SystemDept",
				Path: "/system/dept/index",
				Meta: owl.Meta{
					Title: "部门管理",
					Icon:  "ep:unlock",
				},
				MenuType: owl.MenuTypeMenu,
			},
			apiMenu,
		},
	}
}
