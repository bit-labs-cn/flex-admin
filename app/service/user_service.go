package service

import (
	"errors"

	"github.com/spf13/cast"

	"bit-labs.cn/flex-admin/app/event"
	"bit-labs.cn/flex-admin/app/model"
	"bit-labs.cn/flex-admin/app/provider/jwt"
	"bit-labs.cn/flex-admin/app/repository"
	"bit-labs.cn/owl/provider/conf"
	"bit-labs.cn/owl/provider/db"
	"bit-labs.cn/owl/provider/router"
	"bit-labs.cn/owl/utils"
	"github.com/asaskevich/EventBus"
	"github.com/casbin/casbin/v2"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

var (
	ErrLogin = errors.New("用户名或密码错误")
)

type UserBatchFields struct {
	Username string `json:"username" validate:"required,min=3,max=32"`
	NickName string `json:"nickName" validate:"omitempty,max=32"`
	Email    string `json:"email" validate:"omitempty,email"`
	Phone    string `json:"phone" validate:"omitempty,numeric"`
	Remark   string `json:"remark" validate:"omitempty,max=255"`
	Status   int    `json:"status" validate:"omitempty,oneof=0 1"`
	Sex      int    `json:"sex" validate:"omitempty,oneof=1 2 3"`
	Source   string `json:"source" validate:"omitempty,max=32"`
	SourceID string `json:"sourceId" validate:"omitempty,max=64"`
}

// CreateUserReq 创建用户
type CreateUserReq struct {
	UserBatchFields
	Password string `json:"password" validate:"required,min=6,max=64"`
}

type UpdateUserReq struct {
	ID uint `json:"id,string,omitempty"`
	UserBatchFields
}

type RetrieveUserReq struct {
	router.PageReq
	Keyword string `json:"keyword" validate:"omitempty,max=64"`
}

type LoginReq struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResp struct {
	*model.User
	AccessToken string `json:"accessToken"`
}

type ChangePasswordReq struct {
	UserID      uint   `json:"userId,string"`
	OldPassword string `json:"oldPassword" validate:"required,min=6,max=64"`
	NewPassword string `json:"newPassword" validate:"required,min=6,max=64"`
}

// ResetPasswordReq 超管重置用户密码
type ResetPasswordReq struct {
	UserID      uint   `json:"userId,string" validate:"required"`
	NewPassword string `json:"newPassword" validate:"required,min=6,max=64"`
}

type AssignPermissionReq struct {
	PermissionIds []uint `json:"permissionIds"`
	UserId        uint   `json:"userId" binding:"required"`
}

// AssignMenuToUser 分配菜单给用户
type AssignMenuToUser struct {
	UserID  uint     `json:"userId,string"` // 用户id
	MenuIDs []string `json:"menuIds"`       // 菜单列表
}

type AssignRolesReq struct {
	UserID  uint   `json:"userId,string"` // 用户id
	RoleIDs []uint `json:"roleIds"`       // 角色ids
}

type UserService struct {
	db         *gorm.DB
	menuManger *router.MenuRepository
	jwtSvc     *jwt.JWTService
	db.BaseRepository[model.User]
	roleSvc   *RoleService
	enforcer  casbin.IEnforcer
	userRepo  repository.UserRepositoryInterface
	eventBus  EventBus.Bus
	configure *conf.Configure
}

func NewUserService(
	roleSvc *RoleService,
	userRepo repository.UserRepositoryInterface,
	tx *gorm.DB,
	enforcer casbin.IEnforcer,
	eventBus EventBus.Bus,
	configure *conf.Configure,
	jwtSvc *jwt.JWTService,
	menuManager *router.MenuRepository,
) *UserService {
	return &UserService{
		db:             tx,
		roleSvc:        roleSvc,
		userRepo:       userRepo,
		enforcer:       enforcer,
		BaseRepository: db.NewBaseRepository[model.User](tx),
		eventBus:       eventBus,
		configure:      configure,
		jwtSvc:         jwtSvc,
		menuManger:     menuManager,
	}
}

// Login 用户登录
func (i *UserService) Login(req *LoginReq) (resp *LoginResp, err error) {

	user, err := i.GetUserByName(req.Username)
	if err != nil {
		return nil, err
	}

	// 校验密码是否正确
	if ok := utils.BcryptCheck(req.Password, user.Password); !ok {
		return nil, ErrLogin
	}

	// 非超管才需要获取菜单及权限
	if !user.IsSuperAdmin {
		roleIDs := user.GetRoleIDs()
		menuIDs := i.roleSvc.GetRolesMenuIDs(roleIDs...)
		user.Permissions = i.menuManger.GetPermissionsByMenuIDs(menuIDs...)
	}

	token, err := i.jwtSvc.GenerateToken(user)
	return &LoginResp{
		User:        user,
		AccessToken: token,
	}, err
}

// LoginByThirdParty 第三方用户登录
func (i *UserService) LoginByThirdParty(username, provider string) (resp *LoginResp, err error) {

	thirdProvider, err := i.userRepo.GetByNameAndThirdProvider(username, provider)
	if err != nil {
		return nil, err
	}

	token, err := i.jwtSvc.GenerateToken(&thirdProvider)
	return &LoginResp{
		User:        &thirdProvider,
		AccessToken: token,
	}, err
}

// GetUserByName 根据用户名查找用户
func (i *UserService) GetUserByName(name string) (*model.User, error) {
	var user model.User

	var adminLoginReq LoginReq
	err := i.configure.GetConfig("app.admin", &adminLoginReq)
	if err != nil {
		return nil, err
	}

	// 查找用户，优先从配置里面找
	if name == adminLoginReq.Username {
		user = model.NewSuperUser()
		user.Password = adminLoginReq.Password
	} else {
		// 从数据库查找用户
		user, err = i.userRepo.GetByName(name)
		if errors.Is(err, repository.ErrUserNotExists) {
			return nil, ErrLogin
		}
	}

	return &user, nil
}

// GetUserMenus 获取用户菜单
func (i *UserService) GetUserMenus(userID uint) []*router.Menu {

	user, err := i.userRepo.FindById(userID)
	if err != nil {
		return nil
	}

	menuIDs := i.roleSvc.GetRolesMenuIDs(user.GetRoleIDs()...)

	menus := i.menuManger.GetMenuByMenuIDs(menuIDs...)
	return menus
}

// AssignRoleToUser 分配角色给用户
func (i *UserService) AssignRoleToUser(req *AssignRoleToUser) error {

	roles := db.GetModelsByIDs[model.Role](req.RoleIDs)

	user, err := i.userRepo.FindById(req.UserID)
	if err != nil {
		return err
	}
	user.SetRoles(roles)
	err = i.userRepo.Save(user)

	i.eventBus.Publish(event.AssignRoleToUser, req)
	return err
}

// GetUserRoleIDs 获取用户的角色IDs
func (i *UserService) GetUserRoleIDs(id uint) ([]string, error) {
	user, err := i.userRepo.FindById(id)
	if err != nil {
		return nil, err
	}
	return user.GetRoleIDs(), nil
}

// ChangeUserPassword 修改用户密码
func (i *UserService) ChangeUserPassword(req *ChangePasswordReq) error {
	user, err := i.userRepo.FindById(req.UserID)
	if err != nil {
		return err
	}
	err = user.ChangePassword(req.OldPassword, req.NewPassword)
	if err != nil {
		return err
	}
	return i.userRepo.Save(user)
}

// ResetUserPassword 超管重置用户密码（不校验旧密码）
func (i *UserService) ResetUserPassword(req *ResetPasswordReq) error {
	user, err := i.userRepo.FindById(req.UserID)
	if err != nil {
		return err
	}
	user.Password = utils.BcryptHash(req.NewPassword)
	return i.userRepo.Save(user)
}

type ChangeAvatarReq struct {
	UserID uint   `json:"userId,string"`
	Avatar string `json:"avatar" binding:"required,url"`
}

// ChangeUserAvatar 修改用户头像
func (i *UserService) ChangeUserAvatar(req *ChangeAvatarReq) error {
	user, err := i.userRepo.FindById(req.UserID)
	if err != nil {
		return err
	}
	user.SetAvatar(req.Avatar)

	return i.userRepo.Save(user)
}

// RetrieveUsers 获取用户列表
func (i *UserService) RetrieveUsers(req *RetrieveUserReq) (count int, list []model.User, err error) {
	c, u, e := i.userRepo.Retrieve(req.Page, req.PageSize, func(tx *gorm.DB) {
		db.AppendWhereFromStruct(tx, req)
		tx.Preload("Roles")
	})

	return cast.ToInt(c), u, e
}

// CreateUser 创建用户
func (i *UserService) CreateUser(req *CreateUserReq) error {
	var user model.User
	err := copier.Copy(&user, req)
	if err != nil {
		return err
	}
	user.Password = utils.BcryptHash(req.Password)

	if err = i.userRepo.Save(&user); err != nil {
		return err
	}
	return err
}

// Register 注册用户
func (i *UserService) Register(req *model.User) error {

	var user model.User
	err := copier.Copy(&user, req)
	if err != nil {
		return err
	}

	return i.userRepo.Save(&user)
}

// UpdateUser 更新用户
func (i *UserService) UpdateUser(req *UpdateUserReq) error {

	user, err := i.userRepo.FindById(req.ID)
	if err != nil {
		return err
	}

	err = copier.Copy(&user, req)
	if err != nil {
		return err
	}

	return i.userRepo.Save(user)
}

// ChangeUserStatus 修改用户状态
func (i *UserService) ChangeUserStatus(req *db.ChangeStatus) error {
	return i.BaseRepository.ChangeStatus(req)
}

func (i *UserService) DeleteUser(id uint) error {
	err := i.BaseRepository.Delete(id)
	return err
}
