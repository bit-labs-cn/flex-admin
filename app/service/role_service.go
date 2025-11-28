package service

import (
	"context"

	"bit-labs.cn/flex-admin/app/event"
	"bit-labs.cn/flex-admin/app/model"
	"bit-labs.cn/flex-admin/app/repository"
	"bit-labs.cn/owl/contract/log"
	"bit-labs.cn/owl/provider/db"
	"bit-labs.cn/owl/provider/router"
	"github.com/asaskevich/EventBus"
	"github.com/casbin/casbin/v2"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

// CreateRoleReq 创建角色
type CreateRoleReq struct {
	Name   string `json:"name" validate:"required,min=2,max=32"`
	Code   string `json:"code" validate:"required,alphanum,min=2,max=32"`
	Remark string `json:"remark" validate:"omitempty,max=255"`
}

// UpdateRoleReq 更新角色
type UpdateRoleReq struct {
	ID uint `json:"id,string" validate:"required"`
	CreateRoleReq
}

// AssignMenuToRole 分配菜单给角色, 菜单和按钮权限
type AssignMenuToRole struct {
	RoleID  uint     `json:"roleID,string" validate:"required"`
	MenuIDs []string `json:"menuIds" validate:"required"`
}

// AssignRoleToUser 分配角色给用户
type AssignRoleToUser struct {
	UserID  uint     `json:"userID" validate:"required"`
	RoleIDs []string `json:"roleIDs" validate:"required"`
}

type RetrieveRoleReq struct {
	router.PageReq
	NameLike string `json:"nameLike" validate:"omitempty,max=32"`
	Code     string `json:"code" validate:"omitempty,alphanum,max=32"`
	Status   uint8  `json:"status" validate:"omitempty,oneof=0 1"`
}

// RoleService 角色服务
type RoleService struct {
	db.BaseRepository[model.Role]
	enforcer casbin.IEnforcer
	ctx      context.Context
	log      log.Logger

	roleRepo repository.RoleRepositoryInterface
	menuRepo *router.MenuRepository
	eventbus EventBus.Bus
}

func NewRoleService(menuManager *router.MenuRepository, roleRepo repository.RoleRepositoryInterface, enforcer casbin.IEnforcer, bus EventBus.Bus) *RoleService {
	return &RoleService{
		menuRepo: menuManager,
		enforcer: enforcer,
		roleRepo: roleRepo,
		eventbus: bus,
	}
}
func (i *RoleService) WithContext(ctx context.Context) *RoleService {
	i.ctx = ctx
	return i
}
func (i *RoleService) CreateRole(req *CreateRoleReq) error {
	var role model.Role
	err := copier.Copy(&role, req)
	if err != nil {
		return err
	}

	role.Status = 1 // 默认启用
	return i.roleRepo.Save(&role)
}

func (i *RoleService) UpdateRole(req *UpdateRoleReq) error {

	role, err := i.roleRepo.Detail(req.ID)
	if err != nil {
		return err
	}
	err = copier.Copy(&role, req)
	if err != nil {
		return err
	}
	return i.roleRepo.Save(role)
}

// ChangeStatus 修改角色状态
func (i *RoleService) ChangeStatus(req *db.ChangeStatus) error {
	return i.BaseRepository.ChangeStatus(req)
}

// DeleteRole 删除角色
func (i *RoleService) DeleteRole(id uint) error {
	return i.BaseRepository.Delete(id)
}

func (i *RoleService) RetrieveRoles(req *RetrieveRoleReq) (count int64, list []model.Role, err error) {
	return i.roleRepo.Retrieve(req.Page, req.PageSize, func(tx *gorm.DB) {
		db.AppendWhereFromStruct(tx, req)
	})
}

func (i *RoleService) AssignMenusToRole(req *AssignMenuToRole) error {

	role, err := i.roleRepo.Detail(req.RoleID)
	if err != nil {
		return err
	}

	role.SetMenus(db.GetModelsByIDs[model.Menu](req.MenuIDs))

	err = i.roleRepo.Save(role)
	i.eventbus.Publish(event.AssignMenuToRole, req)
	return err
}

// GetRolesMenuIDs 获取角色的菜单IDs
func (i *RoleService) GetRolesMenuIDs(ids ...string) (result []string) {
	ds, err := i.roleRepo.GetRolesMenuIDs(ids...)
	if err != nil {
		i.log.Error("获取角色菜单IDs失败", err)
		return nil
	}
	return ds
}
