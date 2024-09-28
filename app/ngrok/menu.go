package ngrok

import (
	"github.com/guoliang1994/gin-flex-admin/owl"
)

// admin:controller 的 module + 方法名    admin:user:create
// 一个权限可能会用到多个接口（接口受到访问控制，一个接口需要一个权限）
var menu = []*owl.Menu{
	{
		Path: "/system",
		Name: "System",
		Meta: owl.Meta{
			Title: "Ngrok穿透",
			Icon:  "el-icon-user",
		},
		MenuType: owl.MenuTypeDir,
		Children: []*owl.Menu{
			{
				Name: "SystemApi",
				Path: "/system/api/index",
				Meta: owl.Meta{
					Title: "客户端",
					Icon:  "el-icon-user",
				},
				MenuType: owl.MenuTypeMenu,
			},
			{
				Name: "SystemApi",
				Path: "/system/api/index",
				Meta: owl.Meta{
					Title: "域名解析",
					Icon:  "el-icon-user",
				},
				MenuType: owl.MenuTypeMenu,
			},
			{
				Name: "SystemApi",
				Path: "/system/api/index",
				Meta: owl.Meta{
					Title: "TCP隧道",
					Icon:  "el-icon-user",
				},
				MenuType: owl.MenuTypeMenu,
			},
			{
				Name: "SystemApi",
				Path: "/system/api/index",
				Meta: owl.Meta{
					Title: "UDP隧道",
					Icon:  "el-icon-user",
				},
				MenuType: owl.MenuTypeMenu,
			},
			{
				Name: "SystemApi",
				Path: "/system/api/index",
				Meta: owl.Meta{
					Title: "HTTP代理",
					Icon:  "el-icon-user",
				},
				MenuType: owl.MenuTypeMenu,
			},
			{
				Name: "SystemApi",
				Path: "/system/api/index",
				Meta: owl.Meta{
					Title: "Socket代理",
					Icon:  "el-icon-user",
				},
				MenuType: owl.MenuTypeMenu,
			},
			{
				Name: "SystemApi",
				Path: "/system/api/index",
				Meta: owl.Meta{
					Title: "私密代理",
					Icon:  "el-icon-user",
				},
				MenuType: owl.MenuTypeMenu,
			},
			{
				Name: "SystemApi",
				Path: "/system/api/index",
				Meta: owl.Meta{
					Title: "P2P连接",
					Icon:  "el-icon-user",
				},
				MenuType: owl.MenuTypeMenu,
			},
			{
				Name: "SystemApi",
				Path: "/system/api/index",
				Meta: owl.Meta{
					Title: "文件访问",
					Icon:  "el-icon-user",
				},
				MenuType: owl.MenuTypeMenu,
			},
		},
	},
}
