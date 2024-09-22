package code_gen

import (
	"github.com/guoliang1994/gin-flex-admin/owl"
)

var menu = []*owl.Menu{
	{
		Path: "/system",
		Name: "File",
		Meta: owl.Meta{
			Title: "文件管理",
			Icon:  "el-icon-user",
		},
		Rank: 13,
		Children: []*owl.Menu{
			{
				Name: "SystemUser",
				Path: "/system/user/index",
				Meta: owl.Meta{
					Title: "媒体库",
					Icon:  "el-icon-user",
				},
			},
			{
				Name: "SystemRole",
				Path: "/system/role/index",
				Meta: owl.Meta{
					Title: "大文件上传",
					Icon:  "el-icon-user",
				},
			},
			{
				Name: "SystemRole",
				Path: "/system/role/index",
				Meta: owl.Meta{
					Title: "大文件上传",
					Icon:  "el-icon-user",
				},
			},
		},
	},
	{
		Path: "/system",
		Name: "DevKit",
		Meta: owl.Meta{
			Title: "开发工具",
			Icon:  "el-icon-user",
		},
		Rank: 13,
		Children: []*owl.Menu{
			{
				Name: "SystemUser",
				Path: "/system/user/index",
				Meta: owl.Meta{
					Title: "代码生成器",
					Icon:  "el-icon-user",
				},
			},
			{
				Name: "SystemRole",
				Path: "/system/role/index",
				Meta: owl.Meta{
					Title: "角色管理",
					Icon:  "el-icon-user",
				},
			},
		}},
}

var device = &owl.Menu{
	Path: "/system",
	Meta: owl.Meta{
		Title: "设备管理",
		Icon:  "el-icon-user",
	},
	Rank: 13,
	Children: []*owl.Menu{
		{
			Name: "SystemUser",
			Path: "/system/user/index",
			Meta: owl.Meta{
				Title: "设备类型",
				Icon:  "el-icon-user",
			},
		},
		{
			Name: "SystemRole",
			Path: "/system/role/index",
			Meta: owl.Meta{
				Title: "设备列表",
				Icon:  "el-icon-user",
			},
		},
	},
}
