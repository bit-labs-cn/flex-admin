package router

import (
	v1 "github.com/guoliang1994/gin-flex-admin/app/cms/handle/v1"
	"github.com/guoliang1994/gin-flex-admin/owl"
	"github.com/guoliang1994/gin-flex-admin/owl/middleware"
)

var classifyMenu, articleMenu, documentMenu *owl.Menu

func InitApi(app *owl.Application, appName string) {
	gv1 := app.Engine().Group("/api/v1").Use(middleware.Cors(), middleware.PermissionCheck(app))

	{
		classifyHandle := v1.NewClassifyHandle(app)
		r := owl.NewRouterInfoBuilder(appName, classifyHandle, gv1, owl.MenuOption{
			ComponentName: "CMSClassify",
			Path:          "/cms/classify/index",
			Icon:          "ep:user",
		})

		r.Post("/classify", owl.Authorized, classifyHandle.Create).Name("创建分类").Build()
		r.Delete("/classify/:id", owl.Authorized, classifyHandle.Delete).Name("删除分类").Build()
		r.Put("/classify/:id", owl.Authorized, classifyHandle.Update).Name("更新分类").Build()
		r.Get("/classify", owl.Authorized, classifyHandle.Retrieve).Name("获取分类").Build()

		classifyMenu = r.GetMenu()
	}

	{
		articleHandler := v1.NewArticleHandle(app)
		r := owl.NewRouterInfoBuilder(appName, articleHandler, gv1, owl.MenuOption{
			ComponentName: "CMSArticle",
			Path:          "/cms/article/index",
			Icon:          "ep:user",
		})

		r.Post("/articles", owl.Authorized, articleHandler.Create).Name("创建文章").Build()
		r.Delete("/articles/:id", owl.Authorized, articleHandler.Delete).Name("删除文章").Build()
		r.Put("/articles/:id", owl.Authorized, articleHandler.Update).Name("更新文章").Build()
		r.Get("/articles", owl.Authorized, articleHandler.Retrieve).Name("分页获取文章").Build()

		articleMenu = r.GetMenu()
	}

}
