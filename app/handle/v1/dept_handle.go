package v1

import (
	"bit-labs.cn/flex-admin/app/service"
	"bit-labs.cn/owl/provider/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

var _ router.Handler = (*DeptHandle)(nil)
var _ router.CrudHandler = (*DeptHandle)(nil)

type DeptHandle struct {
	deptSvc *service.DeptService
}

func NewDeptHandle(deptSvc *service.DeptService) *DeptHandle {
	return &DeptHandle{deptSvc: deptSvc}
}
func (i DeptHandle) ModuleName() (en string, zh string) {
	return "dept", "部门管理"
}

//	@Summary		创建部门
//	@Description	创建新的部门信息
//	@Tags			部门管理
//	@Name			创建部门
//	@Param			createDeptReq	body		service.CreateDeptReq	true	"部门创建请求参数"
//	@Success		200				{object}	router.RouterInfo		"部门创建成功"
//	@Failure		400				{object}	router.RouterInfo		"请求参数错误"
//	@Failure		500				{object}	router.RouterInfo		"服务器内部错误"
//	@Router			/api/v1/dept [POST]

func (i DeptHandle) Create(ctx *gin.Context) {
	var req service.CreateDeptReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		router.Fail(ctx, "参数绑定失败")
		return
	}

	err := i.deptSvc.CreateDept(&req)
	router.Auto(ctx, nil, err)
}

//	@Summary		更新部门
//	@Description	根据部门ID更新部门信息
//	@Tags			部门管理
//	@Name			更新部门
//	@Param			id				path		int						true	"部门ID"
//	@Param			updateDeptReq	body		service.UpdateDeptReq	true	"部门更新请求参数"
//	@Success		200				{object}	router.RouterInfo		"部门更新成功"
//	@Failure		400				{object}	router.RouterInfo		"请求参数错误"
//	@Failure		500				{object}	router.RouterInfo		"服务器内部错误"
//	@Router			/api/v1/dept/:id [PUT]

func (i DeptHandle) Update(ctx *gin.Context) {
	var req service.UpdateDeptReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		router.Fail(ctx, "参数绑定失败")
		return
	}

	err := i.deptSvc.UpdateDept(&req)
	router.Auto(ctx, nil, err)
}

// Delete 删除部门
//
//	@Summary		删除部门
//	@Description	根据部门ID删除指定部门
//	@Tags			部门管理
//	@Router			/api/v1/dept/{id} [DELETE]
//	@Access			AccessAuthorized
//	@Name			删除部门
//	@Param			id	path		int					true	"部门ID"
//	@Success		200	{object}	router.RouterInfo	"部门删除成功"
//	@Failure		400	{object}	router.RouterInfo	"请求参数错误"
//	@Failure		500	{object}	router.RouterInfo	"服务器内部错误"
func (i DeptHandle) Delete(ctx *gin.Context) {
	id := cast.ToUint(ctx.Param("id"))
	err := i.deptSvc.DeleteDept(id)
	router.Auto(ctx, nil, err)
}

// Retrieve 获取部门列表
//
//	@Summary		获取部门列表
//	@Description	获取所有部门的列表信息
//	@Tags			部门管理
//	@Router			/api/v1/dept [GET]
//	@Access			AccessAuthorized
//	@Name			获取部门列表
//	@Success		200	{object}	router.RouterInfo	"部门列表获取成功"
//	@Failure		400	{object}	router.RouterInfo	"请求参数错误"
//	@Failure		500	{object}	router.RouterInfo	"服务器内部错误"
func (i DeptHandle) Retrieve(ctx *gin.Context) {
	_, list, err := i.deptSvc.RetrieveDepts()
	router.Auto(ctx, list, err)
}

func (i DeptHandle) Detail(ctx *gin.Context) {

}
