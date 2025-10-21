package v1

import (
	"bit-labs.cn/flex-admin/app/service"
	"bit-labs.cn/owl"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

var _ owl.Handler = (*DeptHandle)(nil)
var _ owl.CrudHandler = (*DeptHandle)(nil)

type DeptHandle struct {
	deptSvc *service.DeptService
}

func NewDeptHandle(deptSvc *service.DeptService) *DeptHandle {
	return &DeptHandle{deptSvc: deptSvc}
}
func (i DeptHandle) ModuleName() (en string, zh string) {
	return "dept", "部门管理"
}

// Create
// @Tags      Dept
// @Summary   创建部门
// @Security  JWT
// @Accept    json
// @Produce   json
// @Param    createDeptReq body  service.CreateDeptReq  true  "请求参数"
// @Success   200  {object}  gin.RouteInfo{data=[]gin.RouteInfo}
// @Router    /api/v1/dept [post]
func (i DeptHandle) Create(ctx *gin.Context) {
	var req service.CreateDeptReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		owl.Fail(ctx, "参数绑定失败")
		return
	}

	err := i.deptSvc.CreateDept(&req)
	owl.Auto(ctx, nil, err)
}

// Update
// @Tags      Dept
// @Summary   更新部门
// @Security  JWT
// @Accept    json
// @Produce   json
// @Param    createDeptReq body  service.UpdateDeptReq  true  "请求参数"
// @Success   200  {object}  gin.RouteInfo{data=[]gin.RouteInfo}
// @Router    /api/v1/dept [post]
func (i DeptHandle) Update(ctx *gin.Context) {
	var req service.UpdateDeptReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		owl.Fail(ctx, "参数绑定失败")
		return
	}

	err := i.deptSvc.UpdateDept(&req)
	owl.Auto(ctx, nil, err)
}

// Delete
// @Tags      Dept
// @Summary   删除部门
// @Security  JWT
// @Accept    json
// @Produce   json
// @Param    createDeptReq body  service.UpdateDeptReq  true  "请求参数"
// @Success   200  {object}  gin.RouteInfo{data=[]gin.RouteInfo}
// @Router    /api/v1/dept [delete]
func (i DeptHandle) Delete(ctx *gin.Context) {
	id := cast.ToUint(ctx.Param("id"))
	err := i.deptSvc.DeleteDept(id)
	owl.Auto(ctx, nil, err)
}

// Retrieve
// @Tags      Dept
// @Summary   部门列表
// @Security  JWT
// @Accept    json
// @Produce   json
// @Param    createDeptReq body  service.UpdateDeptReq  true  "请求参数"
// @Success   200  {object}  gin.RouteInfo{data=[]gin.RouteInfo}
// @Router    /api/v1/dept [get]
func (i DeptHandle) Retrieve(ctx *gin.Context) {
	_, list, err := i.deptSvc.RetrieveDepts()
	owl.Auto(ctx, list, err)
}

func (i DeptHandle) Detail(ctx *gin.Context) {

}
