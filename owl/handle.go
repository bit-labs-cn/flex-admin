package owl

import "github.com/gin-gonic/gin"

// Handler 控制器需要实现此接口，用于获取模块名称，用于设置 api 接口信息
type Handler interface {
	ModuleName() (en string, zh string)
}

type CrudHandler interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Retrieve(ctx *gin.Context)
	Detail(ctx *gin.Context)
}
