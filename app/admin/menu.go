package admin

import (
	"github.com/guoliang1994/gin-flex-admin/owl"
)

var menu = []*owl.Menu{
	{
		Name: "超级管理员",
		Url:  "",
		Icon: "el-icon-user",
		Child: []*owl.Menu{
			{
				Name: "用户管理",
				Url:  "/user/list",
				Icon: "el-icon-user",
				Child: []*owl.Menu{
					{
						Name: "添加用户",
						Url:  "/user/add",
						Icon: "el-icon-user",
					},
					{
						Name: "编辑用户",
						Url:  "/user/edit",
						Icon: "el-icon-user",
					},
					{
						Name: "删除用户",
						Url:  "/user/delete",
						Icon: "el-icon-user",
					},
					{
						Name: "重置密码",
						Url:  "/user/reset",
						Icon: "el-icon-user",
					},
				},
			},
			{
				Name: "角色管理",
				Url:  "/user/list",
				Icon: "el-icon-user",
			},
			{
				Name: "菜单管理",
				Url:  "/user/add",
				Icon: "el-icon-user",
			},
			{
				Name: "Api管理",
				Url:  "/user/add",
				Icon: "el-icon-user",
			},
			{
				Name: "字典管理",
				Url:  "/user/add",
				Icon: "el-icon-user",
			},
			{
				Name: "操作历史",
				Url:  "/user/add",
				Icon: "el-icon-user",
			},
		},
	},
}
