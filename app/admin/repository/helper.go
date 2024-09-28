package repository

import (
	"github.com/guoliang1994/gin-flex-admin/app/admin/repository/model"
	"github.com/spf13/cast"
	"reflect"
)

type ModelSet interface {
	model.ModelRole | model.ModelMenu
}

// GetMenuModelsByIDs 获取菜单模型，gorm 多对多关联写数据

func GetModelsByModelIDs[T ModelSet](modelIDs []string) []T {
	var models []T
	for _, id := range modelIDs {
		m := *new(T)
		value := reflect.ValueOf(&m).Elem()
		fieldType := value.FieldByName("ID").Type()
		switch fieldType.Kind() {
		case reflect.Uint:
			value.FieldByName("ID").SetUint(cast.ToUint64(id))
		case reflect.String:
			value.FieldByName("ID").SetString(id)
		default:
			panic("unhandled default case")
		}

		models = append(models, m)
	}
	return models
}
