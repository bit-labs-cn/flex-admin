package main

import (
	"github.com/guoliang1994/gin-flex-admin/app/admin"
	"github.com/guoliang1994/gin-flex-admin/app/admin/repository/model"
	"github.com/guoliang1994/gin-flex-admin/app/code_gen"
	"github.com/guoliang1994/gin-flex-admin/owl"
)

//go:generate go run cmd/main.go
func main() {
	owl.Run(
		new(admin.SubAppAdmin),
		new(code_gen.SubAppCodeGen),
	)
	m := owl.MenuMange.GetMenus()
	for _, m2 := range m {
		iter(m2, 1)
	}
}
func iter(menu *owl.Menu, level int) {
	db := owl.DB
	if level == 1 {
		menu.Ancestors = menu.Name
		db.Model(&model.ModelMenu{}).Create(&model.ModelMenu{
			Name:       menu.Name,
			Path:       menu.Path,
			Rank:       menu.Rank,
			Level:      1,
			ParentName: "",
		})
	}
	if menu.Children != nil && len(menu.Children) > 0 {
		for j, v := range menu.Children {
			v.Ancestors = menu.Ancestors + "," + v.Name
			iter(v, level+1)

			db.Model(&model.ModelMenu{}).Create(&model.ModelMenu{
				Name:       v.Name,
				Path:       v.Path,
				Rank:       j,
				Level:      level,
				ParentName: menu.Name,
				Ancestors:  v.Ancestors,
			})
		}
	}
}
