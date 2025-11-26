package v1

import (
	"bit-labs.cn/flex-admin/app/model"
	"bit-labs.cn/flex-admin/app/service"
	"bit-labs.cn/owl/contract"
	"bit-labs.cn/owl/provider/db"
	"bit-labs.cn/owl/provider/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

var _ router.Handler = (*UserHandle)(nil)
var _ router.CrudHandler = (*UserHandle)(nil)

type UserHandle struct {
	userSvc  *service.UserService
	roleSvc  *service.RoleService
	menuRepo *router.MenuRepository
}

func (i *UserHandle) ModuleName() (string, string) {
	return "user", "用户管理"
}

func NewUserHandle(userService *service.UserService, roleService *service.RoleService, manager *router.MenuRepository) *UserHandle {
	return &UserHandle{
		userSvc:  userService,
		roleSvc:  roleService,
		menuRepo: manager,
	}
}

// Create 创建用户
//
//	@Summary		创建新用户
//	@Description	创建一个新的用户账户，需要提供用户名、邮箱等基本信息
//	@Tags			用户管理
//	@Router			/api/v1/users [POST]

// @Permission		admin:user:create
// @Name			创建用户
// @Param			request	body		service.CreateUserReq	true	"用户创建请求"
// @Success		200		{object}	router.RouterInfo		"用户创建成功"
// @Failure		400		{object}	router.RouterInfo		"请求参数错误"
// @Failure		500		{object}	router.RouterInfo		"服务器内部错误"
func (i *UserHandle) Create(ctx *gin.Context) {
	req := new(service.CreateUserReq)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	err := i.userSvc.CreateUser(req)
	router.Auto(ctx, nil, err)
}

// Delete 删除用户
//
//	@Summary		删除用户
//	@Description	根据用户ID删除指定用户
//	@Tags			用户管理
//	@Router			/api/v1/users/{id} [DELETE]

// @Name			删除用户
// @Param			id	path		int					true	"用户ID"
// @Success		200	{object}	router.RouterInfo	"用户删除成功"
// @Failure		400	{object}	router.RouterInfo	"请求参数错误"
// @Failure		500	{object}	router.RouterInfo	"服务器内部错误"
func (i *UserHandle) Delete(ctx *gin.Context) {

	id := cast.ToUint(ctx.Param("id"))
	err := i.userSvc.DeleteUser(id)
	router.Auto(ctx, nil, err)
}

func (i *UserHandle) Detail(ctx *gin.Context) {

}

//	@Summary		更新用户信息
//	@Description	根据用户ID更新用户的基本信息
//	@Tags			用户管理
//	@Name			更新用户
//	@Param			id		path		int						true	"用户ID"
//	@Param			request	body		service.UpdateUserReq	true	"用户更新请求"
//	@Success		200		{object}	router.RouterInfo		"用户更新成功"
//	@Failure		400		{object}	router.RouterInfo		"请求参数错误"
//	@Failure		500		{object}	router.RouterInfo		"服务器内部错误"
//	@Router			/api/v1/users/:id [PUT]

func (i *UserHandle) Update(ctx *gin.Context) {

	req := new(service.UpdateUserReq)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	id := cast.ToUint(ctx.Param("id"))
	req.ID = id

	err := i.userSvc.UpdateUser(req)
	router.Auto(ctx, nil, err)
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
	router.Auto(ctx, nil, err)
}

// Retrieve 获取用户列表
//
//	@Summary		获取用户列表
//	@Description	分页获取用户列表，支持搜索和筛选
//	@Tags			用户管理
//	@Router			/api/v1/users [GET]

// @Name			获取用户列表
// @Param			page		query		int										false	"页码"	default(1)
// @Param			pageSize	query		int										false	"每页数量"	default(10)
// @Param			keyword		query		string									false	"搜索关键词"
// @Success		200			{object}	router.RouterInfo{list=[]model.User}	"用户列表获取成功"
// @Failure		400			{object}	router.RouterInfo						"请求参数错误"
// @Failure		500			{object}	router.RouterInfo						"服务器内部错误"
func (i *UserHandle) Retrieve(ctx *gin.Context) {

	req := &service.RetrieveUserReq{
		PageReq: contract.PageReq{
			Page:     1,
			PageSize: 10,
		},
	}
	count, list, err := i.userSvc.RetrieveUsers(req)
	if err != nil {
		router.Fail(ctx, err.Error())
		return
	}
	router.Auto(ctx, gin.H{"list": list, "pageSize": req.PageSize, "currentPage": req.Page, "total": count}, err)
}

// AssignRolesToUser 分配角色给用户
func (i *UserHandle) AssignRolesToUser(ctx *gin.Context) {
	req := new(service.AssignRoleToUser)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		router.Auto(ctx, nil, err)
		return
	}
	req.UserID = cast.ToUint(ctx.Param("id"))
	menuIds := i.userSvc.AssignRoleToUser(req)
	router.Auto(ctx, menuIds, nil)
}

// GetRoleIdsByUserId 获取用户角色
func (i *UserHandle) GetRoleIdsByUserId(ctx *gin.Context) {

	userID := cast.ToUint(ctx.Param("id"))
	ids, err := i.userSvc.GetUserRoleIDs(userID)
	router.Auto(ctx, ids, err)
}

// AssignMenuToUser 分配菜单给用户
func (i *UserHandle) AssignMenuToUser(ctx *gin.Context) {
	req := new(service.AssignRoleToUser)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		router.Auto(ctx, nil, err)
		return
	}
	req.UserID = cast.ToUint(ctx.Param("id"))
	menuIds := i.userSvc.AssignRoleToUser(req)
	router.Auto(ctx, menuIds, nil)
}

// GetMyMenus 获取用户菜单
func (i *UserHandle) GetMyMenus(ctx *gin.Context) {

	user, _ := ctx.Get("user")
	if user.(*model.User).IsSuperAdmin {
		router.Auto(ctx, i.menuRepo.GetMenuWithoutBtn(), nil)
		return
	}
	menus := i.userSvc.GetUserMenus(user.(*model.User).ID)
	router.Auto(ctx, menus, nil)
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
	router.Auto(ctx, login, err)
}
func (i *UserHandle) Register(ctx *gin.Context) {

}

func (i *UserHandle) Me(ctx *gin.Context) {
	value, _ := ctx.Get("user")

	router.Auto(ctx, value, nil)
}
