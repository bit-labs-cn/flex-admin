package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/guoliang1994/gin-flex-admin/app/admin/service"
)

type RoleHandle struct {
}

func NewRoleHandle() RoleHandle {
	return RoleHandle{}
}

func (i *RoleHandle) Retrieve(ctx *gin.Context) {
	svc := service.NewRoleService()
	x := svc.Retrieve(service.RetrieveUserReq{})
	ctx.JSON(200, gin.H{"success": true, "msg": "ok", "data": gin.H{"list": x}, "pageSize": 2, "currentPage": 1, "total": 100})
}

func (i *RoleHandle) MenuIds(ctx *gin.Context) {
	svc := service.NewRoleService()
	x := svc.Retrieve(service.RetrieveUserReq{})
	ctx.JSON(200, gin.H{"success": true, "msg": "ok", "data": gin.H{"list": x}, "pageSize": 2, "currentPage": 1, "total": 100})
}

func (i *RoleHandle) AssignMenu(ctx *gin.Context) {
	svc := service.NewRoleService()
	x := svc.Retrieve(service.RetrieveUserReq{})
	ctx.JSON(200, gin.H{"success": true, "msg": "ok", "data": gin.H{"list": x}, "pageSize": 2, "currentPage": 1, "total": 100})
}
