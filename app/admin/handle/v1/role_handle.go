package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/guoliang1994/gin-flex-admin/app/admin/service"
	"github.com/guoliang1994/gin-flex-admin/owl"
	"github.com/spf13/cast"
)

var _ owl.Handler = (*RoleHandle)(nil)
var _ owl.CrudHandler = (*RoleHandle)(nil)

type RoleHandle struct {
	svc *service.RoleService
}

func (i *RoleHandle) ModuleName() (string, string) {
	return "role", "角色管理"
}

func NewRoleHandle(roleService *service.RoleService) RoleHandle {
	return RoleHandle{
		svc: roleService,
	}
}

// Create 创建角色
func (i *RoleHandle) Create(ctx *gin.Context) {
	req := new(service.CreateRoleRequest)
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

// Detail 获取角色详情
func (i *RoleHandle) Detail(ctx *gin.Context) {

}

// Update 更新角色
func (i *RoleHandle) Update(ctx *gin.Context) {
	req := new(service.UpdateRole)
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

// Delete 删除角色
func (i *RoleHandle) Delete(ctx *gin.Context) {
	id := cast.ToUint(ctx.Param("id"))
	err := i.svc.Delete(id)
	ctx.JSON(200, gin.H{"success": true, "msg": "ok", "data": err})
}

// Retrieve 分页获取角色列表
func (i *RoleHandle) Retrieve(ctx *gin.Context) {
	x := i.svc.Retrieve(service.RetrieveUserReq{})
	ctx.JSON(200, gin.H{"success": true, "msg": "ok", "data": gin.H{"list": x}, "pageSize": 2, "currentPage": 1, "total": 100})
}

// AssignMenusToRole 分配菜单给角色
func (i *RoleHandle) AssignMenusToRole(ctx *gin.Context) {
	req := new(service.AssignMenuToRole)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	id := cast.ToUint(ctx.Param("id"))
	err := i.svc.AssignMenusToRole(id, req)
	ctx.JSON(200, gin.H{"success": true, "msg": "ok", "data": err})
}

// GetRoleMenuIDs 获取角色菜单
func (i *RoleHandle) GetRoleMenuIDs(ctx *gin.Context) {
	id := ctx.Param("id")
	menuIds := i.svc.GetRolesMenuIDs(id)
	ctx.JSON(200, gin.H{"success": true, "msg": "ok", "data": menuIds})

}

// GetRolesSimple 获取所有角色，简单的返回 id 和 name
func (i *RoleHandle) GetRolesSimple(ctx *gin.Context) {
	x, err := i.svc.GetRolesSimple()
	if err != nil {
		ctx.JSON(200, gin.H{"success": false, "msg": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"success": true, "msg": "ok", "data": x})
}

// ChangeStatus 修改角色状态
func (i *RoleHandle) ChangeStatus(ctx *gin.Context) {
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
