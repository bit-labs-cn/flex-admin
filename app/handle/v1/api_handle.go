package v1

import (
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

// @Summary		获取所有接口
// @Description	获取系统中所有已注册的API接口信息
// @Tags			接口管理
// @Produce		json
// @Success		200	{object}	router.Resp{data=[]router.RouterInfo}	"操作成功"
// @Router			/api/v1/api [GET]
func (i ApiHandle) GetAll(c *gin.Context) {

	router.Success(c, router.GetAllRoutes())
}
