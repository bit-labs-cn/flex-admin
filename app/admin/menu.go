package admin

import (
	. "github.com/guoliang1994/gin-flex-admin/app/admin/router"
	"github.com/guoliang1994/gin-flex-admin/owl"
)

// 一个权限可能会用到多个接口（接口受到访问控制，一个接口需要一个权限）
var menu = []*owl.Menu{
	{
		Path: "/system",
		Name: "System",
		Meta: owl.Meta{
			Title: "超级管理员",
			Icon:  "el-icon-user",
		},
		Rank: 13,
		Children: []*owl.Menu{
			{
				Name: "SystemUser",
				Path: "/system/user/index",
				Meta: owl.Meta{
					Title: "用户管理",
					Icon:  "el-icon-user",
				},
				MenuType: owl.MenuTypeMenu,
				Children: []*owl.Menu{
					{
						Name: "SystemUserAdd",
						Path: "/system/user/add",
						Meta: owl.Meta{
							Title: "新增用户",
							Icon:  "el-icon-user",
						},
						MenuType: owl.MenuTypeBtn,
						Apis:     []string{PermissionAddRole},
					},
					{
						Name: "SystemUserUpdate",
						Path: "/system/user/update",
						Meta: owl.Meta{
							Title: "修改用户",
							Icon:  "el-icon-user",
						},
						MenuType: owl.MenuTypeBtn,
						Apis:     []string{PermissionDetailRole, PermissionUpdateRole},
					},
					{
						Name: "SystemUserDelete",
						Path: "/system/user/delete",
						Meta: owl.Meta{
							Title: "删除用户",
							Icon:  "el-icon-user",
						},
						MenuType: owl.MenuTypeBtn,
						Apis:     []string{PermissionDeleteUser},
					},
					{
						Name: "SystemUserDetail",
						Path: "/system/user/detail",
						Meta: owl.Meta{
							Title: "用户详情",
							Icon:  "el-icon-user",
						},
						MenuType: owl.MenuTypeBtn,
						Apis:     []string{PermissionDetailUser},
					},
					{
						Name: "SystemUserChangePassword",
						Path: "/system/user/changePassword",
						Meta: owl.Meta{
							Title: "修改密码",
							Icon:  "el-icon-user",
						},
						MenuType: owl.MenuTypeBtn,
						Apis:     []string{PermissionChangeUserPassword},
					},
					{
						Name: "SystemUserRetrieve",
						Path: "/system/user/retrieve",
						Meta: owl.Meta{
							Title: "分页获取用户",
							Icon:  "el-icon-user",
						},
						MenuType: owl.MenuTypeBtn,
						Apis:     []string{PermissionRetrieveUser},
					},
				},
			},
			{
				Name: "SystemRole",
				Path: "/system/role/index",
				Meta: owl.Meta{
					Title: "角色管理",
					Icon:  "el-icon-user",
				},
				Children: []*owl.Menu{
					{
						Name: "SystemRoleAdd",
						Path: "/system/user/add",
						Meta: owl.Meta{
							Title: "新增角色",
							Icon:  "el-icon-user",
						},
						MenuType: owl.MenuTypeBtn,
						Apis:     []string{PermissionAddRole},
					},
					{
						Name: "SystemRoleUpdate",
						Path: "/system/user/update",
						Meta: owl.Meta{
							Title: "修改角色",
							Icon:  "el-icon-user",
						},
						MenuType: owl.MenuTypeBtn,
						Apis:     []string{PermissionDetailRole, PermissionUpdateRole},
					},
					{
						Name: "SystemRoleDelete",
						Path: "/system/user/delete",
						Meta: owl.Meta{
							Title: "删除角色",
							Icon:  "el-icon-user",
						},
						MenuType: owl.MenuTypeBtn,
						Apis:     []string{PermissionDeleteRole},
					},
					{
						Name: "SystemRoleDetail",
						Path: "/system/user/detail",
						Meta: owl.Meta{
							Title: "角色详情",
							Icon:  "el-icon-user",
						},
						MenuType: owl.MenuTypeBtn,
						Apis:     []string{PermissionDetailRole},
					},
					{
						Name: "SystemRoleRetrieve",
						Path: "/system/role/index",
						Meta: owl.Meta{
							Title: "分页获取角色",
							Icon:  "el-icon-user",
						},
						MenuType: owl.MenuTypeBtn,
						Apis:     []string{PermissionRetrieveRole},
					},
				},
			},
			{
				Name: "SystemMenu",
				Path: "/system/menu/index",
				Meta: owl.Meta{
					Title: "菜单管理",
					Icon:  "el-icon-user",
				},
			},

			{
				Name: "SystemDept",
				Path: "/system/dept/index",
				Meta: owl.Meta{
					Title: "部门管理",
					Icon:  "el-icon-user",
				},
			},

			{
				Name: "SystemMenu",
				Path: "/system/dept/index",
				Meta: owl.Meta{
					Title: "字典管理",
					Icon:  "el-icon-user",
				},
				//Actions: []owl.Action{
				//	{"新增字典", PermissionAddRole, []string{PermissionAddRole}},
				//	{"删除字典", PermissionDeleteRole, []string{PermissionDeleteRole}},
				//	{"字典详情", PermissionDetailRole, []string{PermissionDetailRole}},
				//	{"修改字典", PermissionUpdateRole, []string{PermissionDetailRole, PermissionUpdateRole}},
				//	{"分页获取字典", PermissionRetrieveRole, []string{PermissionRetrieveRole}},
				//},
			},

			{
				Name: "SystemMenu",
				Path: "/system/dept/index",
				Meta: owl.Meta{
					Title: "操作日志",
					Icon:  "el-icon-user",
				},
			},

			{
				Name: "SystemApi",
				Path: "/system/api/index",
				Meta: owl.Meta{
					Title: "api列表",
					Icon:  "el-icon-user",
				},
			},
		},
	},
	{
		Path: "/system",
		Meta: owl.Meta{

			Title: "服务器状态",
			Icon:  "el-icon-user",
		},
		Rank: 19,
	},
}
