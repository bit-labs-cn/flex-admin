package service

import (
	"github.com/guoliang1994/gin-flex-admin/app/admin/repository/model"
	"github.com/guoliang1994/gin-flex-admin/owl"
	jsoniter "github.com/json-iterator/go"
	"gorm.io/gorm"
)

var db *gorm.DB

func MenuStore(app *owl.Application) error {
	m := app.MenuManager().CloneMenus()
	db = app.DB()
	db.Model(&model.ModelMenu{}).Unscoped().Where("1=1").Delete(nil)
	for _, m2 := range m {
		iter(m2, 1)
	}
	return nil
}

func iter(menu *owl.Menu, level int) {
	if level == 1 {
		menu.ID = menu.Name
		meta, _ := jsoniter.MarshalToString(menu.Meta)
		db.Model(&model.ModelMenu{}).Create(&model.ModelMenu{
			ID:       menu.ID,
			Name:     menu.Name,
			Path:     menu.Path,
			Rank:     menu.Rank,
			Level:    1,
			ParentID: "",
			Meta:     meta,
			MenuType: menu.MenuType,
		})
	}
	if menu.Children != nil && len(menu.Children) > 0 {
		for j, v := range menu.Children {
			v.ID = menu.ID + "," + v.Name
			iter(v, level+1)
			meta, _ := jsoniter.MarshalToString(menu.Meta)
			db.Model(&model.ModelMenu{}).Create(&model.ModelMenu{
				ID:       v.ID,
				ParentID: menu.Name,
				Name:     v.Name,
				Path:     v.Path,
				Rank:     j,
				Level:    level,
				MenuType: v.MenuType,
				Meta:     meta,
			})
		}
	}
}
