package v1

import (
	"net/http"

	"bit-labs.cn/owl/provider/router"
	"github.com/gin-gonic/gin"
)

var _ router.Handler = (*ApiHandle)(nil)

type ApiHandle struct {
	engine *gin.Engine
}

func NewApiHandle(engine *gin.Engine) *ApiHandle {
	return &ApiHandle{engine: engine}
}
func (i ApiHandle) ModuleName() (en string, zh string) {
	return "api", "接口管理"
}

//	@Tags			接口管理
//	@Summary		获取所有接口
//	@Description	获取系统中所有已注册的API接口信息
//	@Name			获取所有接口
//	@Success		200	{object}	router.RouterInfo	"接口列表获取成功"
//	@Failure		400	{object}	router.RouterInfo	"请求参数错误"
//	@Failure		500	{object}	router.RouterInfo	"服务器内部错误"
//	@Router			/api/v1/api [GET]
//	@Access			router.AccessAuthorized

func (i ApiHandle) GetAll(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"msg":     "获取api成功",
		"data":    router.GetAllRoutes(),
	})
	return
}
