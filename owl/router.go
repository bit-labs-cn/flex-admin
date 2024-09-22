package owl

type AccessLevel string

const (
	Public        AccessLevel = "开放"
	Authenticated AccessLevel = "需要登录"
	Authorized    AccessLevel = "需要授权"
	AdminOnly     AccessLevel = "仅超管"
	OwnerOnly     AccessLevel = "仅拥有者"
)

type RouterInfo struct {
	Name        string      `json:"name"`
	Module      string      `json:"module"`
	Permission  string      `json:"permission"`
	Description string      `json:"description"`
	AccessLevel AccessLevel `json:"accessLevel"`
}

type RouterInfoBuilder struct {
	module string
	router []RouterInfo
}

func NewRouterInfoBuilder(module string) *RouterInfoBuilder {
	return &RouterInfoBuilder{
		module: module,
	}
}

func (i *RouterInfoBuilder) Add(name, permission string, accessLevel AccessLevel, description string) *RouterInfoBuilder {
	i.router = append(i.router, RouterInfo{
		Name:        name,
		Module:      i.module,
		Permission:  permission,
		Description: description,
		AccessLevel: accessLevel,
	})
	return i
}

func (i *RouterInfoBuilder) Get(index int) RouterInfo {
	return i.router[index]
}
