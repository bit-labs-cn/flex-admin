package service

import (
	"bit-labs.cn/flex-admin/app/model"
	"bit-labs.cn/flex-admin/app/repository"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type DeptService struct {
	deptRepo repository.DeptRepositoryInterface
}

func NewDeptService(deptRepo repository.DeptRepositoryInterface) *DeptService {
	return &DeptService{
		deptRepo: deptRepo,
	}
}

type CreateDeptReq struct {
	Name        string `gorm:"comment:部门名称" json:"name" validate:"required,max=64"`
	ParentId    int    `gorm:"comment:父级部门" json:"parentId,string" validate:"omitempty"`
	Level       uint   `gorm:"comment:部门层级" json:"level" validate:"omitempty"`
	Sort        uint   `gorm:"comment:排序" json:"sort" validate:"omitempty"`
	Status      uint   `gorm:"comment:状态" json:"status" validate:"omitempty,oneof=0 1"`
	Description string `gorm:"comment:描述" json:"description" binding:"omitempty,max=255"`
}

type UpdateDeptReq struct {
	ID uint `json:"id,string,omitempty"`
	CreateDeptReq
}

// CreateDept 创建部门
// 就算 CreateDeptReq 直接使用了 model.Dept 作为了结构体，但是也要单独声明 CreateDeptReq 来接收参数，因为这样可扩展性更高
func (i DeptService) CreateDept(req *CreateDeptReq) error {
	var dept model.Dept
	err := copier.Copy(&dept, req)
	if err != nil {
		return err
	}
	return i.deptRepo.Create(&dept)
}

func (i DeptService) UpdateDept(req *UpdateDeptReq) error {
	var dept model.Dept
	err := copier.Copy(&dept, req)
	if err != nil {
		return err
	}
	return i.deptRepo.Update(&dept)
}
func (i DeptService) DeleteDept(id uint) error {
	return i.deptRepo.Delete(id)
}

func (i DeptService) RetrieveDepts() (count int64, list []model.Dept, err error) {
	return i.deptRepo.Retrieve(1, 1000, func(db *gorm.DB) {

	})
}
