package v1

import (
	"bit-labs.cn/flex-admin/app/service"
	"bit-labs.cn/owl/provider/router"
	"github.com/gin-gonic/gin"
)

var _ router.Handler = (*LogHandle)(nil)

type LogHandle struct {
	logSvc *service.LogService
}

func NewLogHandle(logSvc *service.LogService) *LogHandle {
	return &LogHandle{logSvc: logSvc}
}

func (i *LogHandle) ModuleName() (string, string) {
	return "monitor", "系统监控"
}

// @Summary        登录日志
// @Description    分页查询登录日志
// @Tags           系统监控
// @Produce        json
// @Param          page        query   int     false   "页码"
// @Param          pageSize    query   int     false   "每页数量"
// @Param          userName    query   string  false   "用户名"
// @Param          ip          query   string  false   "IP"
// @Param          userType    query   string  false   "用户类型"
// @Param          start       query   int     false   "开始时间(Unix)"
// @Param          end         query   int     false   "结束时间(Unix)"
// @Success        200         {object} router.PageResp     "操作成功"
// @Failure        500         {object} router.Resp         "服务器内部错误"
// @Router         /api/v1/monitor/login-logs [GET]
func (i *LogHandle) LoginLogs(ctx *gin.Context) {
	var req service.RetrieveLoginLogsReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		router.BadRequest(ctx, err.Error())
		return
	}
	count, list, err := i.logSvc.RetrieveLoginLogs(&req)
	if err != nil {
		router.InternalError(ctx, err)
		return
	}
	router.PageSuccess(ctx, int(count), req.Page, req.PageSize, list)
}

// @Summary        操作日志
// @Description    分页查询操作日志
// @Tags           系统监控
// @Produce        json
// @Param          page        query   int     false   "页码"
// @Param          pageSize    query   int     false   "每页数量"
// @Param          userName    query   string  false   "用户名"
// @Param          path        query   string  false   "路径"
// @Param          method      query   string  false   "方法"
// @Param          status      query   int     false   "状态码"
// @Param          start       query   int     false   "开始时间(Unix)"
// @Param          end         query   int     false   "结束时间(Unix)"
// @Success        200         {object} router.PageResp     "操作成功"
// @Failure        500         {object} router.Resp         "服务器内部错误"
// @Router         /api/v1/monitor/operation-logs [GET]
func (i *LogHandle) OperationLogs(ctx *gin.Context) {
	var req service.RetrieveOperationLogsReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		router.BadRequest(ctx, err.Error())
		return
	}
	count, list, err := i.logSvc.RetrieveOperationLogs(&req)
	if err != nil {
		router.InternalError(ctx, err)
		return
	}
	router.PageSuccess(ctx, int(count), req.Page, req.PageSize, list)
}
