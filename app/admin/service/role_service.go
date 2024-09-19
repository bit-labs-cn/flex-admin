package service

type RoleService struct {
}

type CreateRoleRequest struct {
	Name          string `json:"name" binding:"required"`
	Description   string `json:"description" binding:"required"`
	Sort          string `json:"sort" binding:"required"`
	PermissionIds []uint `json:"permissionIds"`
}

func NewRoleService() *RoleService {
	return &RoleService{}
}
