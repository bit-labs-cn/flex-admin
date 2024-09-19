package service

import (
	"errors"
	"github.com/guoliang1994/gin-flex-admin/app/admin/repository/model"
	"github.com/guoliang1994/gin-flex-admin/owl"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type UserServiceInterface interface {
	Create(req *CreateUserRequest) error
	Register(req model.UserModel) error
	exists(id uint, value model.UserModel) bool
}

// CreateUserRequest 创建用户，同时绑定角色
type CreateUserRequest struct {
	Account  string `json:"account" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
	RoleIds  []uint `json:"roleIds"`
}
type UpdateUserRequest struct {
	Id uint `json:"id,string"`
	CreateUserRequest
}

type UserService struct {
	db *gorm.DB
}

func NewUserService() *UserService {
	return &UserService{
		db: owl.DB,
	}
}

type LoginReq struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserResp struct {
	model.UserModel
	Permissions []string          `json:"permissions"`
	Roles       []model.RoleModel `json:"roles"`
	Menus       []model.MenuModel `json:"menus"`
	Btn         []model.ApiModel  `json:"btn"`
}

func (i *UserService) Login(req *LoginReq) (user *UserResp, err error) {

	err = i.db.Where("account = ?", req.Account).Preload("Roles").First(&user).Error
	if err == nil {
		return nil, errors.New("密码错误")
	}
	return user, err
}

type ChangePasswordReq struct {
	OldPassword string `json:"oldPassword" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required"`
}

func (i *UserService) ChangePassword(req ChangePasswordReq) {

}

type RetrieveUserReq struct {
	Page     int    `json:"page" binding:"required"`
	PageSize int    `json:"pageSize" binding:"required"`
	Keyword  string `json:"keyword"`
}

func (i *UserService) Retrieve(req RetrieveUserReq) {

}

type AssignPermissionReq struct {
	PermissionIds []uint `json:"permissionIds"`
	UserId        uint   `json:"userId" binding:"required"`
}

// AssignPermissions 分配权限给用户,用户可以附加权限
func (i *UserService) AssignPermissions(req AssignPermissionReq) {

}

type AssignRolesReq struct {
	RoleIds []uint `json:"roleIds"`
	UserId  uint
}

func (i *UserService) AssignRoles(req AssignRolesReq) {
	i.RevokeRole(req.UserId)
	for _, id := range req.RoleIds {
		i.db.Create(model.UserRoleModel{UserId: req.UserId, RoleId: id})
	}
}
func (i *UserService) RevokeRole(userId uint) {
	i.db.Model(model.UserRoleModel{}).Where("user_id", userId).Delete(model.UserRoleModel{UserId: userId})
}

var ErrUserExists = errors.New("用户已存在")

func (i *UserService) Create(req *CreateUserRequest) error {
	if i.exists(0, model.UserModel{Account: req.Account}) {
		return ErrUserExists
	}

	var user model.UserModel
	err := copier.Copy(&user, req)
	if err != nil {
		return err
	}

	err = i.db.Transaction(func(tx *gorm.DB) error {
		if err = i.db.Create(&user).Error; err != nil {
			return err
		}
		i.AssignRoles(AssignRolesReq{
			RoleIds: req.RoleIds,
			UserId:  user.ID,
		})
		return nil
	})
	return err
}

// Register 注册用户
func (i *UserService) Register(req model.UserModel) error {
	var user model.UserModel
	err := copier.Copy(&user, req)
	if err != nil {
		return err
	}

	return i.db.Create(&user).Error
}

// Update 更新用户
func (i *UserService) Update(req *UpdateUserRequest) error {
	if i.exists(req.Id, model.UserModel{Account: req.Account}) {
		return ErrUserExists
	}

	return i.db.Transaction(func(tx *gorm.DB) error {
		var user model.UserModel
		err := copier.Copy(&user, req)
		if err != nil {
			return err
		}
		return i.db.Model(&user).Where("id", req.Id).Updates(user).Error
	})
}

func (i *UserService) exists(id uint, value model.UserModel) bool {
	if id > 0 {
		i.db.Where("id != ?", id).Where(value).Find(&model.UserModel{})
	} else {
		i.db.Where(value).Find(&model.UserModel{})
	}
	return true
}
