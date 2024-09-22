package code_gen

import (
	"github.com/guoliang1994/gin-flex-admin/owl"
)

var menu = []*owl.Menu{
	{
		Path: "/system",
		Meta: owl.Meta{
			Title: "CMS系统",
			Icon:  "el-icon-user",
		},
		Rank: 13,
		Children: []*owl.Menu{
			{
				Name: "SystemUser",
				Path: "/system/user/index",
				Meta: owl.Meta{
					Title: "文章分类",
					Icon:  "el-icon-user",
				},
			},
			{
				Name: "SystemRole",
				Path: "/system/role/index",
				Meta: owl.Meta{
					Title: "文章列表",
					Icon:  "el-icon-user",
				},
			},
			{
				Name: "SystemRole",
				Path: "/system/role/index",
				Meta: owl.Meta{
					Title: "文档列表",
					Icon:  "el-icon-user",
				},
			},
		},
	},
}
