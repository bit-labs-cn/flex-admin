package service

import (
	"bit-labs.cn/flex-admin/app/model"
	"bit-labs.cn/flex-admin/app/repository"
	validatorv10 "github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type AreaService struct {
	areaRepo repository.AreaRepositoryInterface
	validate *validatorv10.Validate
}

func NewAreaService(areaRepo repository.AreaRepositoryInterface, validate *validatorv10.Validate) *AreaService {
	return &AreaService{areaRepo: areaRepo, validate: validate}
}

type RetrieveAllAreaReq struct {
	ParentID uint   `json:"parentId,string" validate:"omitempty"` // 父级区域ID
	Name     string `json:"name" validate:"omitempty,max=64"`     // 区域名称
}

func (i *AreaService) RetrieveAll(req *RetrieveAllAreaReq) (list []model.Area, err error) {
	if err := i.validate.Struct(req); err != nil {
		return nil, err
	}
	return i.areaRepo.ListAll(func(db *gorm.DB) {
		if req == nil {
			return
		}
		if req.ParentID > 0 {
			db.Where("parent_id = ?", req.ParentID)
		}
		if req.Name != "" {
			db.Where("name like ?", "%"+req.Name+"%")
		}
	})
}
