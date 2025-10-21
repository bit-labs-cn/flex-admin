package v1

import (
	"bit-labs.cn/flex-admin/app/repository"
	"bit-labs.cn/flex-admin/app/service"
	"bit-labs.cn/owl"
	"bit-labs.cn/owl/contract"
	"bit-labs.cn/owl/db"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net/http"
)

var _ owl.Handler = (*RoleHandle)(nil)
var _ owl.CrudHandler = (*RoleHandle)(nil)

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
func (i *RoleHandle) Update(ctx *gin.Context) {
	req := new(service.UpdateRoleReq)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		owl.Fail(ctx, err.Error())
		return
	}

	id := cast.ToUint(ctx.Param("id"))
	req.ID = id

	err := i.roleService.UpdateRole(req)
	if err != nil {
		owl.Fail(ctx, err.Error())
		return
	}
	owl.Success(ctx, nil)
}

// Delete 删除角色
func (i *RoleHandle) Delete(ctx *gin.Context) {
	id := cast.ToUint(ctx.Param("id"))
	err := i.roleService.DeleteRole(id)
	ctx.JSON(http.StatusOK, gin.H{"success": true, "data": err})
}

// Retrieve 分页获取角色列表
func (i *RoleHandle) Retrieve(ctx *gin.Context) {
	req := service.RetrieveRoleReq{
		PageReq: contract.PageReq{
			Page:     1,
			PageSize: 10,
		},
	}

	count, list, err := i.roleService.RetrieveRoles(&req)
	if err != nil {
		owl.Fail(ctx, err.Error())
		return
	}
	owl.Success(ctx, gin.H{"list": list, "pageSize": req.PageSize, "currentPage": req.Page, "total": count})
}

// AssignMenusToRole 分配菜单给角色
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
func (i *RoleHandle) GetRoleMenuIDs(ctx *gin.Context) {
	id := ctx.Param("id")
	menuIds := i.roleService.WithContext(ctx).GetRolesMenuIDs(id)
	ctx.JSON(http.StatusOK, gin.H{"success": true, "data": menuIds})

}

// RoleOptions 获取所有角色，简单的返回 id 和 name
func (i *RoleHandle) RoleOptions(ctx *gin.Context) {
	x, err := i.roleRepo.Options()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"success": false, "msg": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"success": true, "data": x})
}

// ChangeStatus 修改角色状态
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
