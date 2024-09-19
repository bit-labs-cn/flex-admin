package v1

import (
	"github.com/guoliang1994/gin-flex-admin/app/admin/repository/model"
	"github.com/guoliang1994/gin-flex-admin/owl"
	"github.com/guoliang1994/gin-flex-admin/owl/utils/structs"
)

type RoleHandle struct {
	owl.Handle[*model.RoleModel]
}

func NewRoleHandle() RoleHandle {
	repo := owl.NewRepository[*model.RoleModel](owl.DB).UniqueCheckFn(func(value *model.RoleModel) map[string]interface{} {
		m := structs.Map(value, structs.Underline)
		return map[string]interface{}{
			"name": m["Name"],
		}
	})
	return RoleHandle{
		Handle: owl.NewHandle[*model.RoleModel](repo),
	}
}
