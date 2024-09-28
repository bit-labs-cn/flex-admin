package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/guoliang1994/gin-flex-admin/app/admin/repository/model"
	"github.com/guoliang1994/gin-flex-admin/app/admin/service"
	"github.com/guoliang1994/gin-flex-admin/owl"
	"github.com/spf13/cast"
)

var _ owl.Handler = (*UserHandle)(nil)
var _ owl.CrudHandler = (*UserHandle)(nil)

type UserHandle struct {
	app     *owl.Application
	svc     *service.UserService
	roleSvc *service.RoleService
}

func (i *UserHandle) ModuleName() (string, string) {
	return "user", "用户管理"
}

func NewUserHandle(userService *service.UserService, roleService *service.RoleService, app *owl.Application) UserHandle {
	return UserHandle{
		svc:     userService,
		roleSvc: roleService,
		app:     app,
	}
}

func (i *UserHandle) Create(ctx *gin.Context) {
	req := new(service.CreateUserRequest)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	err := i.svc.Create(req)
	if err != nil {
		return
	}
	ctx.JSON(200, gin.H{"msg": err})
}

func (i *UserHandle) Delete(ctx *gin.Context) {

	id := cast.ToUint(ctx.Param("id"))
	err := i.svc.Delete(id, nil)
	ctx.JSON(200, gin.H{"success": true, "msg": "ok", "data": err})
}

func (i *UserHandle) Detail(ctx *gin.Context) {

}

func (i *UserHandle) List(ctx *gin.Context) {

}

func (i *UserHandle) Update(ctx *gin.Context) {
	req := new(service.UpdateUserRequest)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	id := cast.ToUint(ctx.Param("id"))
	req.ID = id

	err := i.svc.Update(req)
	if err != nil {
		return
	}
	ctx.JSON(200, gin.H{"msg": err})
}

// ChangeStatus 修改角色状态
func (i *UserHandle) ChangeStatus(ctx *gin.Context) {
	req := new(service.ChangeStatus)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	id := cast.ToUint(ctx.Param("id"))
	req.ID = id

	_ = i.svc.ChangeStatus(req)
	ctx.JSON(200, gin.H{"success": true, "msg": "ok"})
}

func (i *UserHandle) Retrieve(ctx *gin.Context) {

	x := i.svc.Retrieve(service.RetrieveUserReq{})
	ctx.JSON(200, gin.H{"success": true, "msg": "ok", "data": gin.H{"list": x}, "pageSize": 2, "currentPage": 1, "total": 100})
}

// AssignRolesToUser 分配角色给用户
func (i *UserHandle) AssignRolesToUser(ctx *gin.Context) {
	req := new(service.AssignRoleToUser)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	id := cast.ToUint(ctx.Param("id"))
	menuIds := i.roleSvc.AssignRoleToUser(id, req)
	ctx.JSON(200, gin.H{"success": true, "msg": "ok", "data": menuIds})
}

// GetRoleIdsByUserId 获取用户角色
func (i *UserHandle) GetRoleIdsByUserId(ctx *gin.Context) {

	userID := cast.ToUint(ctx.Param("id"))
	ids, err := i.svc.GetUserRoleIDs(userID)
	if err != nil {
		return
	}
	ctx.JSON(200, gin.H{"success": true, "msg": "ok", "data": ids})
}

// AssignMenuToUser 分配菜单给用户
func (i *UserHandle) AssignMenuToUser(ctx *gin.Context) {
	req := new(service.AssignRoleToUser)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	id := cast.ToUint(ctx.Param("id"))
	menuIds := i.roleSvc.AssignRoleToUser(id, req)
	ctx.JSON(200, gin.H{"success": true, "msg": "ok", "data": menuIds})
}

// GetMyMenus 获取用户菜单
func (i *UserHandle) GetMyMenus(ctx *gin.Context) {

	user, _ := ctx.Get("user")
	if user.(*model.ModelUser).IsSuperAdmin {
		ctx.JSON(200, gin.H{
			"success": true,
			"msg":     "获取menus成功",
			"data":    i.app.MenuManager().GetMenuWithoutBtn(),
		})
		return
	}
	menus := i.svc.GetMenus(user.(*model.ModelUser).ID)
	ctx.JSON(200, gin.H{
		"success": true,
		"msg":     "获取menus成功",
		"data":    menus,
	})
	return
}
func (i *UserHandle) ChangePassword(ctx *gin.Context) {

}

func (i *UserHandle) ResetPassword(ctx *gin.Context) {

}

func (i *UserHandle) Login(ctx *gin.Context) {
	var req service.LoginReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	login, err := i.svc.Login(&req)
	if err != nil {
		ctx.JSON(200, gin.H{"success": false, "msg": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"success": true, "msg": "ok", "data": login})
}
func (i *UserHandle) Register(ctx *gin.Context) {

}

func (i *UserHandle) Me(ctx *gin.Context) {

}
