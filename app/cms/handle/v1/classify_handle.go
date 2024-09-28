package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/guoliang1994/gin-flex-admin/owl"
)

var _ owl.CrudHandler = (*ClassifyHandle)(nil)
var _ owl.Handler = (*ClassifyHandle)(nil)

type ClassifyHandle struct {
	app *owl.Application
}

func (a ClassifyHandle) ModuleName() (en string, zh string) {
	return "classify", "分类管理"
}

func NewClassifyHandle(app *owl.Application) *ClassifyHandle {
	return &ClassifyHandle{
		app: app,
	}

}

func (a ClassifyHandle) Create(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (a ClassifyHandle) Update(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (a ClassifyHandle) Delete(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (a ClassifyHandle) Retrieve(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (a ClassifyHandle) Detail(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}
