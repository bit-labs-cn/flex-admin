package v1

import (
	"bit-labs.cn/flex-admin/app/model"
	"bit-labs.cn/flex-admin/app/service"
	"bit-labs.cn/owl"
	"bit-labs.cn/owl/contract"
	"bit-labs.cn/owl/db"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
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
	owl.Auto(ctx, nil, err)
}

func (i *UserHandle) Delete(ctx *gin.Context) {

	id := cast.ToUint(ctx.Param("id"))
	err := i.userSvc.DeleteUser(id)
	owl.Auto(ctx, nil, err)
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
	owl.Auto(ctx, nil, err)
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

	err := i.userSvc.ChangeStatus(req)
	owl.Auto(ctx, nil, err)
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
	owl.Auto(ctx, gin.H{"list": list, "pageSize": req.PageSize, "currentPage": req.Page, "total": count}, err)
}

// AssignRolesToUser 分配角色给用户
func (i *UserHandle) AssignRolesToUser(ctx *gin.Context) {
	req := new(service.AssignRoleToUser)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		owl.Auto(ctx, nil, err)
		return
	}
	req.UserID = cast.ToUint(ctx.Param("id"))
	menuIds := i.userSvc.AssignRoleToUser(req)
	owl.Auto(ctx, menuIds, nil)
}

// GetRoleIdsByUserId 获取用户角色
func (i *UserHandle) GetRoleIdsByUserId(ctx *gin.Context) {

	userID := cast.ToUint(ctx.Param("id"))
	ids, err := i.userSvc.GetUserRoleIDs(userID)
	owl.Auto(ctx, ids, err)
}

// AssignMenuToUser 分配菜单给用户
func (i *UserHandle) AssignMenuToUser(ctx *gin.Context) {
	req := new(service.AssignRoleToUser)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		owl.Auto(ctx, nil, err)
		return
	}
	req.UserID = cast.ToUint(ctx.Param("id"))
	menuIds := i.userSvc.AssignRoleToUser(req)
	owl.Auto(ctx, menuIds, nil)
}

// GetMyMenus 获取用户菜单
func (i *UserHandle) GetMyMenus(ctx *gin.Context) {

	user, _ := ctx.Get("user")
	if user.(*model.User).IsSuperAdmin {
		owl.Auto(ctx, i.menuRepo.GetMenuWithoutBtn(), nil)
		return
	}
	menus := i.userSvc.GetUserMenus(user.(*model.User).ID)
	owl.Auto(ctx, menus, nil)
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
	owl.Auto(ctx, login, err)
}
func (i *UserHandle) Register(ctx *gin.Context) {

}

func (i *UserHandle) Me(ctx *gin.Context) {
	value, _ := ctx.Get("user")

	owl.Auto(ctx, value, nil)
}
