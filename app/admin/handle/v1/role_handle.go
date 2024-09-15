package v1

import (
	"github.com/guoliang1994/gin-flex-admin/app/admin/repository/model"
	"github.com/guoliang1994/gin-flex-admin/owl"
)

type RoleHandle struct {
	owl.Handle[*model.RoleModel]
}

func NewRoleHandle() RoleHandle {
	return RoleHandle{
		Handle: owl.NewHandle[*model.RoleModel](),
	}
}
