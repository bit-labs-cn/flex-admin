package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/guoliang1994/gin-flex-admin/app/admin/service"
)

type UserHandle struct {
}

func NewUserHandle() UserHandle {
	return UserHandle{}
}

func (i *UserHandle) Create(ctx *gin.Context) {
	req := new(service.CreateUpdateUserRequest)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	svc := service.NewUserService()
	err := svc.Create(req)
	if err != nil {
		return
	}
	ctx.JSON(200, gin.H{"msg": err})
}
