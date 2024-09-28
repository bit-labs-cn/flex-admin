package router

import "github.com/guoliang1994/gin-flex-admin/owl"

func InitMenu() *owl.Menu {
	return &owl.Menu{

		Path: "/cms",
		Name: "CMS",
		Meta: owl.Meta{
			Title: "CMS系统",
			Icon:  "el-icon-user",
		},
		Rank:     14,
		MenuType: owl.MenuTypeDir,
		Children: []*owl.Menu{
			classifyMenu,
			articleMenu,
		},
	}
}
