package service

import (
	"github.com/guoliang1994/gin-flex-admin/app/admin/repository/model"
	"github.com/guoliang1994/gin-flex-admin/owl"
	"gorm.io/gorm"
)

// role 1:n menu 1:n action 1:0.1 btn 1:n api
type RoleService struct {
	db *gorm.DB
}

type CreateRoleRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Sort        string `json:"sort" binding:"required"`
}

func NewRoleService() *RoleService {
	return &RoleService{
		db: owl.DB,
	}
}

func (i *RoleService) Create(CreateRoleRequest) error {
	return nil
}

func (i *RoleService) GetMenus() error {
	return nil
}

// AssignRolePermission 角色分配权限, 菜单和按钮权限
type AssignRolePermission struct {
	RoleId    uint     `json:"roleId" binding:"required"`
	MenuNames []string `json:"menuNames"` // name 是唯一键
	BtnKeys   []string `json:"btnKeys"`
}

func (i *RoleService) AssignPermissions(AssignRolePermission) error {
	return nil
}

func (i *RoleService) GetBtnPermissions() error {
	return nil
}
func (i *RoleService) Retrieve(req RetrieveUserReq) []model.ModelRole {
	var users []model.ModelRole
	i.db.Find(&users)
	return users
}
