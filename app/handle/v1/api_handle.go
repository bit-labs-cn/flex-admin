package v1

import (
	"bit-labs.cn/owl"
	"github.com/gin-gonic/gin"
	"net/http"
)

var _ owl.Handler = (*ApiHandle)(nil)

type ApiHandle struct {
	engine *gin.Engine
}

func NewApiHandle(engine *gin.Engine) *ApiHandle {
	return &ApiHandle{engine: engine}
}
func (i ApiHandle) ModuleName() (en string, zh string) {
	return "api", "接口管理"
}

// GetAll
// @Tags      Admin
// @Summary   获取所有接口
// @Security  JWT
// @Accept    application/json
// @Produce   application/json
// @Param     data  body      string  false  "请求参数"
// @Success   200  {object}  gin.RouteInfo{data=[]gin.RouteInfo}
// @Router    /api/v1/api [get]
func (i ApiHandle) GetAll(c *gin.Context) {
	routers := i.engine.GetAllRoutes()
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"msg":     "获取api成功",
		"data":    routers,
	})
	return
}
