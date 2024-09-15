package owl

type AccessLevel string

const (
	Public        AccessLevel = "开放"
	Authenticated AccessLevel = "需要登录"
	Authorized    AccessLevel = "需要授权"
	AdminOnly     AccessLevel = "仅超管"
)

type RouterInfo struct {
	Name        string      `json:"name"`
	Module      string      `json:"module"`
	Description string      `json:"description"`
	AccessLevel AccessLevel `json:"accessLevel"`
}
