package owl

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"runtime"
	"strings"
	"sync"
)

type AccessLevel string

const (
	Public        AccessLevel = "开放接口"
	Authenticated AccessLevel = "需要登录"
	Authorized    AccessLevel = "需要授权"
	AdminOnly     AccessLevel = "仅超管"
)

type RouterInfo struct {
	Method      string          `json:"method"`      // 请求方式
	Path        string          `json:"path"`        // 路由路径
	Name        string          `json:"name"`        // 路由名称
	Module      string          `json:"module"`      // 模块名称
	Permission  string          `json:"permission"`  // 权限标识
	Description string          `json:"description"` // 描述
	AccessLevel AccessLevel     `json:"accessLevel"` // 访问级别
	handle      gin.HandlerFunc // 处理函数
}

type Dep struct {
	Handler Handler
	Method  gin.HandlerFunc
}
type RouterInfoBuilder struct {
	menu     *Menu
	moduleEn string // 模块英文名称
	moduleZh string // 模块中文名称
	appName  string
	router   gin.IRoutes
	handler  Handler // 定义 handler，可以获取 handler 的 moduleName
	lock     sync.RWMutex

	currentRouter *RouterInfo
	deps          []Dep
}

type MenuOption struct {
	ComponentName string // 组件名称
	Path          string // 路由地址
	Icon          string // 图标
}

func NewRouterInfoBuilder(appName string, handle Handler, router gin.IRoutes, meta MenuOption) *RouterInfoBuilder {
	en, zh := handle.ModuleName()
	return &RouterInfoBuilder{
		moduleEn: en,
		moduleZh: zh,
		router:   router,
		handler:  handle,
		appName:  appName,
		menu: &Menu{
			Name: meta.ComponentName,
			Path: meta.Path,
			Meta: Meta{
				Title: zh,
				Icon:  meta.Icon,
			},
			MenuType: MenuTypeMenu,
		},
		currentRouter: &RouterInfo{},
	}
}

func (i *RouterInfoBuilder) add(method, path string, accessLevel AccessLevel, handle gin.HandlerFunc) *RouterInfoBuilder {
	i.lock.Lock()
	i.currentRouter = &RouterInfo{
		Method:      method,
		Path:        path,
		Module:      i.moduleZh,
		AccessLevel: accessLevel,
		handle:      handle,
	}

	return i
}

func (i *RouterInfoBuilder) getPermissionStr(funcName string) string {
	return i.appName + ":" + i.moduleEn + ":" + funcName
}

// Build 构建路由,以及菜单
func (i *RouterInfoBuilder) Build() {
	defer i.lock.Unlock()

	// 设置路由授权标识，前端可用于控制按钮的显示
	handleName := nameOfFunction(i.currentRouter.handle)
	i.currentRouter.Permission = i.getPermissionStr(handleName)

	var permissions []string
	permissions = append(permissions, i.currentRouter.Permission)

	// 设置接口依赖，完成一个动作可能会需要很多个接口配合
	for _, dep := range i.deps {
		funcName := nameOfFunction(dep.Method)
		en, _ := dep.Handler.ModuleName()
		permission := i.appName + ":" + en + ":" + funcName
		permissions = append(permissions, permission)
	}
	i.deps = nil // 清理依赖

	// 构建菜单（按钮）
	if i.currentRouter.AccessLevel == Authorized {
		i.menu.Children = append(i.menu.Children, &Menu{
			Name: handleName,
			Meta: Meta{
				Title: i.currentRouter.Name,
			},
			MenuType:    MenuTypeBtn,
			Permissions: permissions,
		})
	}
	// 添加路由
	i.router.Handle(&gin.RouteInfo{
		Method: i.currentRouter.Method,
		Path:   i.currentRouter.Path,
		Extra:  i.currentRouter,
	}, i.currentRouter.handle)
}

func (i *RouterInfoBuilder) Description(d string) *RouterInfoBuilder {
	i.currentRouter.Description = d
	return i
}

func (i *RouterInfoBuilder) Name(n string) *RouterInfoBuilder {
	i.currentRouter.Name = n
	return i
}

// Deps 此接口依赖的其他接口，比如说修改用户，则需要 `获取用户详情` `更新用户`  两个接口
func (i *RouterInfoBuilder) Deps(dep ...Dep) *RouterInfoBuilder {
	i.deps = append(i.deps, dep...)
	return i
}
func (i *RouterInfoBuilder) Post(path string, accessLevel AccessLevel, handle gin.HandlerFunc) *RouterInfoBuilder {
	return i.add(http.MethodPost, path, accessLevel, handle)
}
func (i *RouterInfoBuilder) Put(path string, accessLevel AccessLevel, handle gin.HandlerFunc) *RouterInfoBuilder {
	return i.add(http.MethodPut, path, accessLevel, handle)
}
func (i *RouterInfoBuilder) Delete(path string, accessLevel AccessLevel, handle gin.HandlerFunc) *RouterInfoBuilder {
	return i.add(http.MethodDelete, path, accessLevel, handle)
}
func (i *RouterInfoBuilder) Get(path string, accessLevel AccessLevel, handle gin.HandlerFunc) *RouterInfoBuilder {
	return i.add(http.MethodGet, path, accessLevel, handle)
}

func (i *RouterInfoBuilder) GetMenu() *Menu {
	return i.menu
}

// 获取函数名
func nameOfFunction(f any) string {
	fullName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	// github.com/guoliang1994/gin-flex-admin/app/admin/handle/v1.(*RoleHandle).Detail-fm
	lastIndex := strings.LastIndex(fullName, ".")
	if lastIndex > 0 {
		return strings.Replace(fullName[lastIndex+1:], "-fm", "", 1)
	}
	return fullName
}
