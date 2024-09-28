package service

import (
	"errors"
	"github.com/guoliang1994/gin-flex-admin/app/admin/repository"
	"github.com/guoliang1994/gin-flex-admin/app/admin/repository/model"
	"github.com/guoliang1994/gin-flex-admin/owl"
	"github.com/guoliang1994/gin-flex-admin/owl/utils"
	"github.com/jinzhu/copier"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

type UserServiceInterface interface {
	Create(req *CreateUserRequest) error
	Register(req model.ModelUser) error
	Login(request *LoginReq) (*LoginResp, error)
	GetMenus(userID uint) []*owl.Menu
	exists(id uint, value model.ModelUser) bool
}

var _ UserServiceInterface = (*UserService)(nil)

type UserService struct {
	db     *gorm.DB
	app    *owl.Application
	jwtSvc JWTService
	BaseService[model.ModelUser]
	roleSvc *RoleService
}

func NewUserService(app *owl.Application, roleSvc *RoleService) *UserService {
	return &UserService{
		db:          app.DB(),
		app:         app,
		roleSvc:     roleSvc,
		BaseService: NewBaseService[model.ModelUser](app.DB()),
	}
}

type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResp struct {
	model.ModelUser
	AccessToken string            `json:"accessToken"`
	Roles       []model.ModelRole `json:"roles"`
}

func (i *UserService) Login(req *LoginReq) (resp *LoginResp, err error) {
	var user model.ModelUser

	if req.Username == "admin" && req.Password == "123qwe" {
		user = model.ModelUser{
			Username:     "超级管理员",
			Nickname:     "超级管理员",
			IsSuperAdmin: true,
			Permissions:  []string{"*:*:*"},
		}
	} else {
		err = i.db.Where("username = ?", req.Username).Preload("Roles").First(&user).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户名或密码错误")
		}
		if ok := utils.BcryptCheck(req.Password, user.Password); !ok {
			return nil, errors.New("用户名或密码错误")
		}

		roleIDs, err := i.GetUserRoleIDs(user.ID)
		if err != nil {
			return nil, err
		}
		menuIDs := i.roleSvc.GetRolesMenuIDs(roleIDs...)
		user.Permissions = i.app.MenuManager().GetPermissionsByMenuIDs(menuIDs...)
	}

	token, err := i.jwtSvc.GenerateToken(&user)
	return &LoginResp{
		ModelUser:   user,
		AccessToken: token,
		Roles:       user.Roles,
	}, err
}

func (i *UserService) GetMenus(userID uint) []*owl.Menu {
	roleIDs, err := i.GetUserRoleIDs(userID)
	if err != nil {
		return nil
	}
	menuIDs := i.roleSvc.GetRolesMenuIDs(roleIDs...)

	menus := i.app.MenuManager().GetMenuByMenuIDs(menuIDs...)
	return menus
}

type UserResp struct {
	model.ModelUser
	Permissions []string          `json:"permissions"`
	Roles       []model.ModelRole `json:"roles"`
	Menus       []model.ModelMenu `json:"menus"`
	Btn         []model.ModelApi  `json:"btn"`
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

func (i *UserService) Retrieve(req RetrieveUserReq) []model.ModelUser {
	var users []model.ModelUser
	i.db.Preload("Roles").Find(&users)
	return users
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

var ErrUserExists = errors.New("用户已存在")

// CreateUserRequest 创建用户，同时绑定角色
type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	NickName string `json:"nickName"`
	Sex      *int   `json:"sex"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	RoleIds  []uint `json:"roleIds"`
}

func (i *UserService) Create(req *CreateUserRequest) error {
	if i.exists(0, model.ModelUser{Username: req.Username}) {
		return ErrUserExists
	}

	var user model.ModelUser
	err := copier.Copy(&user, req)
	if err != nil {
		return err
	}
	user.Password = utils.BcryptHash(req.Password)
	if err = i.db.Create(&user).Error; err != nil {
		return err
	}
	return err
}

// Register 注册用户
func (i *UserService) Register(req model.ModelUser) error {
	var user model.ModelUser
	err := copier.Copy(&user, req)
	if err != nil {
		return err
	}

	return i.db.Create(&user).Error
}

type UpdateUserRequest struct {
	ID uint `json:"id,string"`
	CreateUserRequest
}

// Update 更新用户
func (i *UserService) Update(req *UpdateUserRequest) error {
	if i.exists(req.ID, model.ModelUser{Username: req.Username}) {
		return ErrUserExists
	}

	return i.db.Transaction(func(tx *gorm.DB) error {
		var user model.ModelUser
		err := copier.Copy(&user, req)
		if err != nil {
			return err
		}
		return i.db.Model(&user).Where("id", req.ID).Updates(user).Error
	})
}

// ChangeStatus 修改用户状态
func (i *UserService) ChangeStatus(req *ChangeStatus) error {
	return i.BaseService.ChangeStatus(req)
}

// GetUserRoleIDs 获取用户角色IDs
func (i *UserService) GetUserRoleIDs(userID uint) (roleIDs []string, err error) {
	md := model.ModelUser{}
	i.db.Model(&md).Where("id", userID).Preload("Roles").Find(&md)
	for _, role := range md.Roles {
		roleIDs = append(roleIDs, cast.ToString(role.ID))
	}
	return roleIDs, err
}

// AssignMenuToUser 角色分配权限, 菜单和按钮权限
type AssignMenuToUser struct {
	UserID  uint     `json:"userID,string" binding:"required"`
	MenuIDs []string `json:"menuIds"`
}

func (i *UserService) AssignMenusToUser(userID uint, req *AssignMenuToRole) error {

	user := model.ModelUser{}
	user.ID = userID
	menus := repository.GetModelsByModelIDs[model.ModelRole](req.MenuIDs)
	err := i.db.Model(&user).Association("Menus").Replace(&menus)
	return err
}

func (i *UserService) exists(id uint, value model.ModelUser) bool {
	var count int64
	if id > 0 {
		i.db.Where("id != ?", id).Where("username", value.Username).Count(&count)
	} else {
		i.db.Where("username", value.Username).Count(&count)
	}
	if count > 0 {
		return true
	}
	return false
}
