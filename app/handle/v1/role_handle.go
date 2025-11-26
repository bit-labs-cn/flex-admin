package v1

import (
	"net/http"

	"bit-labs.cn/flex-admin/app/repository"
	"bit-labs.cn/flex-admin/app/service"
	"bit-labs.cn/owl/contract"
	"bit-labs.cn/owl/provider/db"
	"bit-labs.cn/owl/provider/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

var _ router.Handler = (*RoleHandle)(nil)
var _ router.CrudHandler = (*RoleHandle)(nil)

type RoleHandle struct {
	roleService *service.RoleService
	roleRepo    repository.RoleRepositoryInterface
}

func (i *RoleHandle) ModuleName() (string, string) {
	return "role", "角色管理"
}

func NewRoleHandle(roleService *service.RoleService, roleRepository repository.RoleRepositoryInterface) *RoleHandle {
	return &RoleHandle{
		roleService: roleService,
		roleRepo:    roleRepository,
	}
}

// Create 创建角色
//
//	@Summary		创建角色
//	@Description	创建新的角色信息
//	@Tags			角色管理
//	@Name			创建角色
//	@Param			createRoleReq	body		service.CreateRoleReq	true	"角色创建请求参数"
//	@Success		200				{object}	router.RouterInfo		"角色创建成功"
//	@Failure		400				{object}	router.RouterInfo		"请求参数错误"
//	@Failure		500				{object}	router.RouterInfo		"服务器内部错误"
//	@Router			/api/v1/role [POST]

// @Permission		admin:role:create
func (i *RoleHandle) Create(ctx *gin.Context) {
	req := new(service.CreateRoleReq)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	err := i.roleService.CreateRole(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": err})
}

// Detail 获取角色详情
func (i *RoleHandle) Detail(ctx *gin.Context) {

}

// Update 更新角色
//
//	@Summary		更新角色
//	@Description	根据角色ID更新角色信息
//	@Tags			角色管理
//	@Router			/api/v1/role/{id} [PUT]

// @Name			更新角色
// @Param			id				path		int						true	"角色ID"
// @Param			updateRoleReq	body		service.UpdateRoleReq	true	"角色更新请求参数"
// @Success		200				{object}	router.RouterInfo		"角色更新成功"
// @Failure		400				{object}	router.RouterInfo		"请求参数错误"
// @Failure		500				{object}	router.RouterInfo		"服务器内部错误"
func (i *RoleHandle) Update(ctx *gin.Context) {
	req := new(service.UpdateRoleReq)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		router.Fail(ctx, err.Error())
		return
	}

	id := cast.ToUint(ctx.Param("id"))
	req.ID = id

	err := i.roleService.UpdateRole(req)
	if err != nil {
		router.Fail(ctx, err.Error())
		return
	}
	router.Success(ctx, nil)
}

// Delete 删除角色
//
//	@Summary		删除角色
//	@Description	根据角色ID删除指定角色
//	@Tags			角色管理
//	@Router			/api/v1/role/{id} [DELETE]

// @Name			删除角色
// @Param			id	path		int					true	"角色ID"
// @Success		200	{object}	router.RouterInfo	"角色删除成功"
// @Failure		400	{object}	router.RouterInfo	"请求参数错误"
// @Failure		500	{object}	router.RouterInfo	"服务器内部错误"
func (i *RoleHandle) Delete(ctx *gin.Context) {
	id := cast.ToUint(ctx.Param("id"))
	err := i.roleService.DeleteRole(id)
	ctx.JSON(http.StatusOK, gin.H{"success": true, "data": err})
}

// Retrieve 获取角色列表
//
//	@Summary		获取角色列表
//	@Description	分页获取角色列表，支持搜索和筛选
//	@Tags			角色管理
//	@Router			/api/v1/role [GET]

// @Name			获取角色列表
// @Param			page		query		int										false	"页码"	default(1)
// @Param			pageSize	query		int										false	"每页数量"	default(10)
// @Param			keyword		query		string									false	"搜索关键词"
// @Success		200			{object}	router.RouterInfo{list=[]model.Role}	"角色列表获取成功"
// @Failure		400			{object}	router.RouterInfo						"请求参数错误"
// @Failure		500			{object}	router.RouterInfo						"服务器内部错误"
func (i *RoleHandle) Retrieve(ctx *gin.Context) {
	req := service.RetrieveRoleReq{
		PageReq: contract.PageReq{
			Page:     1,
			PageSize: 10,
		},
	}

	count, list, err := i.roleService.RetrieveRoles(&req)
	if err != nil {
		router.Fail(ctx, err.Error())
		return
	}
	router.Success(ctx, gin.H{"list": list, "pageSize": req.PageSize, "currentPage": req.Page, "total": count})
}

// AssignMenusToRole 分配菜单给角色
//
//	@Summary		分配菜单给角色
//	@Description	为指定角色分配菜单权限
//	@Tags			角色管理
//	@Router			/api/v1/role/{id}/menus [POST]

// @Name			分配菜单给角色
// @Param			id				path		int							true	"角色ID"
// @Param			assignMenuReq	body		service.AssignMenuToRole	true	"菜单分配请求参数"
// @Success		200				{object}	router.RouterInfo			"菜单分配成功"
// @Failure		400				{object}	router.RouterInfo			"请求参数错误"
// @Failure		500				{object}	router.RouterInfo			"服务器内部错误"
func (i *RoleHandle) AssignMenusToRole(ctx *gin.Context) {
	req := new(service.AssignMenuToRole)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	req.RoleID = cast.ToUint(ctx.Param("id"))
	err := i.roleService.WithContext(ctx).AssignMenusToRole(req)
	ctx.JSON(http.StatusOK, gin.H{"success": true, "data": err})
}

// GetRoleMenuIDs 获取角色菜单
//
//	@Summary		获取角色菜单ID列表
//	@Description	获取指定角色拥有的菜单ID列表
//	@Tags			角色管理
//	@Router			/api/v1/role/{id}/menus [GET]

// @Name			获取角色菜单ID列表
// @Param			id	path		string							true	"角色ID"
// @Success		200	{object}	router.RouterInfo{data=[]int}	"角色菜单ID列表获取成功"
// @Failure		400	{object}	router.RouterInfo				"请求参数错误"
// @Failure		500	{object}	router.RouterInfo				"服务器内部错误"
func (i *RoleHandle) GetRoleMenuIDs(ctx *gin.Context) {
	id := ctx.Param("id")
	menuIds := i.roleService.WithContext(ctx).GetRolesMenuIDs(id)
	ctx.JSON(http.StatusOK, gin.H{"success": true, "data": menuIds})

}

// RoleOptions 获取角色选项
//
//	@Summary		获取角色选项
//	@Description	获取所有角色的简单信息，用于下拉选择等场景
//	@Tags			角色管理
//	@Router			/api/v1/role/options [GET]

// @Name			获取角色选项
// @Success		200	{object}	router.RouterInfo	"角色选项获取成功"
// @Failure		400	{object}	router.RouterInfo	"请求参数错误"
// @Failure		500	{object}	router.RouterInfo	"服务器内部错误"
func (i *RoleHandle) RoleOptions(ctx *gin.Context) {
	x, err := i.roleRepo.Options()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"success": false, "msg": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"success": true, "data": x})
}

// @Summary		修改角色状态
// @Description	启用或禁用指定角色
// @Tags			角色管理
// @Name			修改角色状态
// @Param			id				path		int					true	"角色ID"
// @Param			changeStatusReq	body		db.ChangeStatus		true	"状态修改请求参数"
// @Success		200				{object}	router.RouterInfo	"角色状态修改成功"
// @Failure		400				{object}	router.RouterInfo	"请求参数错误"
// @Failure		500				{object}	router.RouterInfo	"服务器内部错误"
// @Router			/api/v1/role/{id}/status [PUT]

// @Permissions	PermissionAdminUserCreate
func (i *RoleHandle) ChangeStatus(ctx *gin.Context) {

	req := new(db.ChangeStatus)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	id := cast.ToUint(ctx.Param("id"))
	req.ID = id
	_ = i.roleService.ChangeStatus(req)
	ctx.JSON(http.StatusOK, gin.H{"success": true})
}
