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
	req := new(service.CreateUserRequest)
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

func (i *UserHandle) Delete(ctx *gin.Context) {

}

func (i *UserHandle) Detail(ctx *gin.Context) {

}

func (i *UserHandle) List(ctx *gin.Context) {

}

func (i *UserHandle) Update(ctx *gin.Context) {

}

func (i *UserHandle) Retrieve(ctx *gin.Context) {

	svc := service.NewUserService()
	x := svc.Retrieve(service.RetrieveUserReq{})
	ctx.JSON(200, gin.H{"success": true, "msg": "ok", "data": gin.H{"list": x}, "pageSize": 2, "currentPage": 1, "total": 100})
}

func (i *UserHandle) RevokeRole(ctx *gin.Context) {

}
func (i *UserHandle) AssignPermissions(ctx *gin.Context) {

}

func (i *UserHandle) AssignRoles(ctx *gin.Context) {

}

func (i *UserHandle) ChangePassword(ctx *gin.Context) {

}

func (i *UserHandle) Login(ctx *gin.Context) {

}

func (i *UserHandle) Register(ctx *gin.Context) {

}
