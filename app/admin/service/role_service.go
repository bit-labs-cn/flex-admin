package service

import (
	"github.com/guoliang1994/gin-flex-admin/app/admin/repository"
	"github.com/guoliang1994/gin-flex-admin/app/admin/repository/model"
	"github.com/guoliang1994/gin-flex-admin/owl"
	"github.com/jinzhu/copier"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

// RoleService 角色服务
type RoleService struct {
	db  *gorm.DB
	app *owl.Application
	BaseService[model.ModelRole]
}

func NewRoleService(db *gorm.DB, app *owl.Application) *RoleService {
	return &RoleService{
		db:          db,
		app:         app,
		BaseService: NewBaseService[model.ModelRole](db),
	}
}

type CreateRoleRequest struct {
	Name   string `json:"name" binding:"required"`
	Code   string `json:"code" binding:"required"`
	Remark string `json:"remark"`
}

func (i *RoleService) Create(req *CreateRoleRequest) error {
	newDB := i.db.Model(&model.ModelRole{})
	var m model.ModelRole
	_ = copier.Copy(&m, req)
	m.Status = 1 // 默认启用
	return newDB.Create(&m).Error
}

type UpdateRole struct {
	ID uint `json:"id,string" binding:"required"`
	CreateRoleRequest
}

func (i *RoleService) Update(req *UpdateRole) error {
	newDB := i.db.Model(&model.ModelRole{})
	var m model.ModelRole
	_ = copier.Copy(&m, req)
	return newDB.Where("id", req.ID).Updates(&m).Error
}

type ChangeStatus struct {
	ID     uint `json:"id,string" binding:"required"`
	Status int  `json:"status"`
}

// ChangeStatus 修改角色状态
func (i *RoleService) ChangeStatus(req *ChangeStatus) error {
	return i.BaseService.ChangeStatus(req)
}

// Delete 删除角色
func (i *RoleService) Delete(id uint) error {

	err := i.BaseService.Delete(id, func(db *gorm.DB) (msg string, b bool) {
		var count int64
		db.Model(model.ModelUser{}).Where("role_id", id).Count(&count)
		if count > 0 {
			return "该角色下还有用户，无法删除", false
		}
		return "", true
	})
	return err
}

func (i *RoleService) Retrieve(req RetrieveUserReq) []model.ModelRole {
	var users []model.ModelRole
	i.db.Find(&users)
	return users
}

type RoleItem struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// GetRolesSimple 获取所有角色
func (i *RoleService) GetRolesSimple() (result []RoleItem, err error) {
	newDb := i.db.Model(&model.ModelRole{})
	newDb.Scan(&result)
	return
}

// AssignMenuToRole 分配菜单给角色, 菜单和按钮权限
type AssignMenuToRole struct {
	MenuIDs []string `json:"menuIds"`
}

func (i *RoleService) AssignMenusToRole(roleID uint, req *AssignMenuToRole) error {
	role := model.ModelRole{}
	role.ID = roleID

	i.db.Model(&model.ModelRoleMenu{}).Where("role_id", roleID).Delete(nil)
	menus := repository.GetModelsByModelIDs[model.ModelMenu](req.MenuIDs)
	err := i.db.Model(&role).Association("Menus").Replace(&menus)

	permissions := i.app.MenuManager().GetPermissionsByMenuIDs(req.MenuIDs...)
	var rules [][]string
	for _, permission := range permissions {
		rules = append(rules, []string{cast.ToString(role.ID), permission})
	}

	i.app.Enforcer().RemoveFilteredPolicy(0, cast.ToString(role.ID))
	_, err = i.app.Enforcer().AddPolicies(rules)

	return err
}

type AssignRoleToUser struct {
	RoleIDs []string `json:"roleIDs"`
}

// AssignRoleToUser 分配角色给用户
func (i *RoleService) AssignRoleToUser(userID uint, req *AssignRoleToUser) error {
	user := model.ModelUser{}
	user.ID = userID
	roles := repository.GetModelsByModelIDs[model.ModelRole](req.RoleIDs)

	var rules [][]string
	for _, roleID := range req.RoleIDs {
		rules = append(rules, []string{cast.ToString(userID), roleID})
	}
	i.app.Enforcer().RemoveFilteredGroupingPolicy(0, cast.ToString(userID))
	i.app.Enforcer().AddGroupingPolicies(rules)

	return i.db.Model(&user).Association("Roles").Replace(&roles)
}

// GetRolesMenuIDs 获取角色的菜单IDs
func (i *RoleService) GetRolesMenuIDs(ids ...string) (result []string) {
	i.db.Where("role_id in ?", ids).Model(&model.ModelRoleMenu{}).Select("menu_Id").Scan(&result)
	return
}
