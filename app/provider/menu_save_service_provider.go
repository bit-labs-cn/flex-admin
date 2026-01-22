package provider

import (
	"bit-labs.cn/flex-admin/app/model"
	"bit-labs.cn/owl"
	"bit-labs.cn/owl/contract/foundation"
	"bit-labs.cn/owl/provider/router"
	jsoniter "github.com/json-iterator/go"
	"gorm.io/gorm"
)

type MenuSaveServiceProvider struct {
	app foundation.Application
	db  *gorm.DB
}

var _ foundation.ServiceProvider = (*MenuSaveServiceProvider)(nil)

func (i *MenuSaveServiceProvider) Register() {

}

func (i *MenuSaveServiceProvider) Boot() {
	err := i.app.Invoke(func(db *gorm.DB, manager *router.MenuRepository) {
		i.db = db
		m := manager.CloneMenus()
		db.Model(&model.Menu{}).Unscoped().Where("1=1").Delete(nil)
		for _, m2 := range m {
			i.iter(m2, 1)
		}
	})
	owl.PanicIf(err)
}

func (i *MenuSaveServiceProvider) iter(menu *router.Menu, level int) {
	if level == 1 {
		menu.ID = menu.Name
		meta, _ := jsoniter.MarshalToString(menu.Meta)
		i.db.Model(&model.Menu{}).Create(&model.Menu{
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
	if len(menu.Children) > 0 {
		for j, v := range menu.Children {
			v.ID = menu.ID + "," + v.Name
			i.iter(v, level+1)
			meta, _ := jsoniter.MarshalToString(menu.Meta)
			i.db.Model(&model.Menu{}).Create(&model.Menu{
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
func (i *MenuSaveServiceProvider) GenerateConf() map[string]string {
	return nil
}
