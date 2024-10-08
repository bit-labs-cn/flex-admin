package v1

import (
	"bit-labs.cn/gin-flex-admin/app/model"
	"bit-labs.cn/gin-flex-admin/app/service"
	"bit-labs.cn/owl"
	"bit-labs.cn/owl/contract"
	"bit-labs.cn/owl/db"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net/http"
)

var _ owl.Handler = (*UserHandle)(nil)
var _ owl.CrudHandler = (*UserHandle)(nil)

type UserHandle struct {
	userSvc  *service.UserService
	roleSvc  *service.RoleService
	menuRepo *owl.MenuRepository
}

func (i *UserHandle) ModuleName() (string, string) {
	return "user", "用户管理"
}

func NewUserHandle(userService *service.UserService, roleService *service.RoleService, manager *owl.MenuRepository) *UserHandle {
	return &UserHandle{
		userSvc:  userService,
		roleSvc:  roleService,
		menuRepo: manager,
	}
}

func (i *UserHandle) Create(ctx *gin.Context) {
	req := new(service.CreateUserReq)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	err := i.userSvc.CreateUser(req)
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": err})
}

func (i *UserHandle) Delete(ctx *gin.Context) {

	id := cast.ToUint(ctx.Param("id"))
	err := i.userSvc.DeleteUser(id)
	ctx.JSON(http.StatusOK, gin.H{"success": true, "data": err})
}

func (i *UserHandle) Detail(ctx *gin.Context) {

}

func (i *UserHandle) Update(ctx *gin.Context) {
	req := new(service.UpdateUserReq)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	id := cast.ToUint(ctx.Param("id"))
	req.ID = id

	err := i.userSvc.UpdateUser(req)
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": err})
}

// ChangeStatus 修改用户状态
func (i *UserHandle) ChangeStatus(ctx *gin.Context) {
	req := new(db.ChangeStatus)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	id := cast.ToUint(ctx.Param("id"))
	req.ID = id

	_ = i.userSvc.ChangeStatus(req)
	ctx.JSON(http.StatusOK, gin.H{"success": true, "msg": "ok"})
}

// Retrieve 获取用户列表
func (i *UserHandle) Retrieve(ctx *gin.Context) {

	req := &service.RetrieveUserReq{
		PageReq: contract.PageReq{
			Page:     1,
			PageSize: 10,
		},
	}
	count, list, err := i.userSvc.RetrieveUsers(req)
	if err != nil {
		owl.Fail(ctx, err.Error())
		return
	}
	owl.Success(ctx, gin.H{"list": list, "pageSize": req.PageSize, "currentPage": req.Page, "total": count})
}

// AssignRolesToUser 分配角色给用户
func (i *UserHandle) AssignRolesToUser(ctx *gin.Context) {
	req := new(service.AssignRoleToUser)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	req.UserID = cast.ToUint(ctx.Param("id"))
	menuIds := i.userSvc.AssignRoleToUser(req)
	ctx.JSON(http.StatusOK, gin.H{"success": true, "data": menuIds})
}

// GetRoleIdsByUserId 获取用户角色
func (i *UserHandle) GetRoleIdsByUserId(ctx *gin.Context) {

	userID := cast.ToUint(ctx.Param("id"))
	ids, err := i.userSvc.GetUserRoleIDs(userID)
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"success": true, "data": ids})
}

// AssignMenuToUser 分配菜单给用户
func (i *UserHandle) AssignMenuToUser(ctx *gin.Context) {
	req := new(service.AssignRoleToUser)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	req.UserID = cast.ToUint(ctx.Param("id"))
	menuIds := i.userSvc.AssignRoleToUser(req)
	ctx.JSON(http.StatusOK, gin.H{"success": true, "data": menuIds})
}

// GetMyMenus 获取用户菜单
func (i *UserHandle) GetMyMenus(ctx *gin.Context) {

	user, _ := ctx.Get("user")
	if user.(*model.User).IsSuperAdmin {
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"msg":     "获取menus成功",
			"data":    i.menuRepo.GetMenuWithoutBtn(),
		})
		return
	}
	menus := i.userSvc.GetUserMenus(user.(*model.User).ID)
	ctx.JSON(http.StatusOK, gin.H{
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

	login, err := i.userSvc.Login(&req)
	if err != nil {
		owl.Fail(ctx, err.Error())
		return
	}
	owl.Success(ctx, login)
}
func (i *UserHandle) Register(ctx *gin.Context) {

}

func (i *UserHandle) Me(ctx *gin.Context) {

}
