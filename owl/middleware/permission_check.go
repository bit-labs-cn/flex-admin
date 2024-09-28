package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/guoliang1994/gin-flex-admin/app/admin/service"
	"github.com/guoliang1994/gin-flex-admin/owl"
	"github.com/guoliang1994/gin-flex-admin/owl/utils"
	"github.com/spf13/cast"
	"strings"
)

func PermissionCheck(app *owl.Application) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		url := ctx.Request.URL.Path
		method := ctx.Request.Method

		apis := app.Routers()

		var find bool
		for _, api := range apis {
			if api.Method == method && utils.UrlIsEq(url, api.Path) {
				find = true
				accessLevel := api.Extra.(*owl.RouterInfo).AccessLevel
				if accessLevel == owl.Public {
					ctx.Next()
					return
				}
				if accessLevel == owl.AdminOnly || accessLevel == owl.Authorized || accessLevel == owl.Authenticated {

					token := ctx.Request.Header.Get("Authorization")
					var JWTService = service.NewJWTService(service.JWTOptions{})
					user, err := JWTService.ParseToken(strings.Replace(token, "Bearer ", "", -1))
					if err != nil {
						return
					}

					ctx.Set("user", user)
					if user.IsSuperAdmin {
						ctx.Next()
						return
					}
					// 只有系统管理员才能访问
					if accessLevel == owl.AdminOnly && !user.IsSuperAdmin {
						_ = ctx.AbortWithError(403, errors.New("未授权的访问"))
						return
					}

					// 需要登录，token 有效则认为已经登录
					if accessLevel == owl.Authenticated {
						ctx.Next()
						return
					}

					// 需要授权
					if accessLevel == owl.Authorized {
						permissionKey := api.Extra.(*owl.RouterInfo).Permission
						app.Enforcer().LoadPolicy()
						can, err := app.Enforcer().Enforce(cast.ToString(user.ID), permissionKey)
						if err != nil {
							_ = ctx.AbortWithError(500, err)
							return
						}
						if !can {
							_ = ctx.AbortWithError(403, errors.New("未授权的访问"))
							return
						}
					}
				}
			}
		}
		if !find {
			_ = ctx.AbortWithError(404, errors.New("未找到匹配的路由"))
		}
	}
}
