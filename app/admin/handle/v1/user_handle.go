package v1

import (
	"github.com/guoliang1994/gin-flex-admin/app/admin/repository/model"
	"github.com/guoliang1994/gin-flex-admin/owl"
)

type UserHandle struct {
	owl.Handle[*model.UserModel]
}

func NewUserHandle() UserHandle {
	return UserHandle{
		Handle: owl.NewHandle[*model.UserModel](),
	}
}
