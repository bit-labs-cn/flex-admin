package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/guoliang1994/gin-flex-admin/owl"
)

var _ owl.CrudHandler = (*ArticleHandle)(nil)
var _ owl.Handler = (*ArticleHandle)(nil)

type ArticleHandle struct {
	app *owl.Application
}

func (a ArticleHandle) ModuleName() (en string, zh string) {
	return "article", "文章管理"
}

func NewArticleHandle(app *owl.Application) *ArticleHandle {
	return &ArticleHandle{
		app: app,
	}

}

func (a ArticleHandle) Create(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (a ArticleHandle) Update(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (a ArticleHandle) Delete(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (a ArticleHandle) Retrieve(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (a ArticleHandle) Detail(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}
