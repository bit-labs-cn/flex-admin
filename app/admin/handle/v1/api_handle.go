package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/guoliang1994/gin-flex-admin/owl"
)

var _ owl.Handler = (*ApiHandle)(nil)

type ApiHandle struct {
	app *owl.Application
}

func NewApiHandle(app *owl.Application) *ApiHandle {
	return &ApiHandle{app: app}
}
func (i ApiHandle) ModuleName() (en string, zh string) {
	return "api", "接口管理"
}

func (i ApiHandle) GetAll(c *gin.Context) {
	routers := i.app.Routers()
	c.JSON(200, gin.H{
		"success": true,
		"msg":     "获取api成功",
		"data":    routers,
	})
	return
}
