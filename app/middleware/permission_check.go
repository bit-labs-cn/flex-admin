package middleware

import (
	"bit-labs.cn/flex-admin/app/service"
	"bit-labs.cn/owl"
	"bit-labs.cn/owl/utils"
	"errors"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net/http"
	"strings"
)

func PermissionCheck(engine *gin.Engine, enforcer casbin.IEnforcer) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		url := ctx.Request.URL.Path
		method := ctx.Request.Method

		routes := engine.GetAllRoutes()

		var findRoute bool
		for _, api := range routes {
			if api.Method == method && utils.UrlIsEq(url, api.Path) {
				findRoute = true
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
						_ = ctx.AbortWithError(http.StatusUnauthorized, errors.New("未授权的访问"))
						return
					}

					ctx.Set("user", user)
					if user.IsSuperAdmin {
						ctx.Next()
						return
					}
					// 只有系统管理员才能访问
					if accessLevel == owl.AdminOnly && !user.IsSuperAdmin {
						_ = ctx.AbortWithError(http.StatusForbidden, errors.New("未授权的访问"))
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
						enforcer.LoadPolicy()
						can, err := enforcer.Enforce(cast.ToString(user.ID), permissionKey)
						if err != nil {
							_ = ctx.AbortWithError(http.StatusInternalServerError, err)
							return
						}
						if !can {
							_ = ctx.AbortWithError(http.StatusForbidden, errors.New("未授权的访问"))
							return
						}
					}
				}
			}
		}
		if !findRoute {
			_ = ctx.AbortWithError(http.StatusNotFound, errors.New("未找到匹配的路由"))
		}
	}
}
