package service

import (
	"bit-labs.cn/flex-admin/app/model"
	"bit-labs.cn/flex-admin/app/repository"
	"gorm.io/gorm"
)

type AreaService struct {
	areaRepo repository.AreaRepositoryInterface
}

func NewAreaService(areaRepo repository.AreaRepositoryInterface) *AreaService {
	return &AreaService{areaRepo: areaRepo}
}

type RetrieveAllAreaReq struct {
	ParentID uint   `json:"parentId,string"` // 父级区域ID
	Name     string `json:"name"`            // 区域名称
}

func (i AreaService) RetrieveAll(req *RetrieveAllAreaReq) (list []model.Area, err error) {
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
