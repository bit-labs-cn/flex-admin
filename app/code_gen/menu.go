package code_gen

import (
	"github.com/guoliang1994/gin-flex-admin/app/admin/repository/model"
	"github.com/guoliang1994/gin-flex-admin/owl"
)

var menu = []*owl.Menu{
	{
		Name: "开发工具",
		Url:  "",
		Icon: "el-icon-user",
		Child: []*owl.Menu{
			{
				Name: "代码生成器",
				Url:  "/user/list",
				Icon: "el-icon-user",
			},
		},
	},
}

func iter(menu *owl.Menu, level int) {
	db := owl.DB
	if level == 1 {
		menu.Path = menu.Name
		db.Model(&model.MenuModel{}).Create(&model.MenuModel{
			Name:     menu.Name,
			Url:      menu.Url,
			Sort:     menu.Sort,
			Icon:     menu.Icon,
			Meta:     menu.Meta,
			Path:     menu.Path,
			Level:    1,
			ParentId: "",
		})
	}
	if menu.Child != nil && len(menu.Child) > 0 {
		for j, v := range menu.Child {
			v.Path = menu.Path + "," + v.Name
			iter(v, level+1)
			db.Model(&model.MenuModel{}).Create(&model.MenuModel{
				Name:     v.Name,
				Url:      v.Url,
				Sort:     j,
				Level:    level,
				Icon:     v.Icon,
				Meta:     v.Meta,
				Path:     v.Path,
				ParentId: menu.Name,
			})
		}
	}
}
