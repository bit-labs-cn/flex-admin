package service

import (
	"errors"
	"strings"

	"bit-labs.cn/flex-admin/app/model"
	"bit-labs.cn/flex-admin/app/repository"
	"bit-labs.cn/owl/provider/db"
	"bit-labs.cn/owl/provider/redis"
	"bit-labs.cn/owl/provider/router"
	validatorv10 "github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

type CreateDictReq struct {
	Name   string `json:"name" validate:"required,max=32"`               // 字典名（中）
	Type   string `json:"type"`                                          // 字典名（英）
	Status uint8  `json:"status,string" validate:"required,min=1,max=2"` // 状态(1启用,2禁用)
	Desc   string `json:"desc"`                                          // 描述
	Sort   uint8  `json:"sort,string" validate:"required,min=1,max=255"` // 排序
}

type UpdateDictReq struct {
	ID uint `json:"id,string"` // 字典ID
	CreateDictReq
}

type DictService struct {
	dictRepo repository.DictRepositoryInterface // 字典仓储接口
	locker   redis.LockerFactory                // 分布式锁工厂
	validate *validatorv10.Validate
}

func NewDictService(dictRepo repository.DictRepositoryInterface, locker redis.LockerFactory, validate *validatorv10.Validate) *DictService {
	return &DictService{
		dictRepo: dictRepo,
		locker:   locker,
		validate: validate,
	}
}
func (i DictService) CreateDict(req *CreateDictReq) error {
	if err := i.validate.Struct(req); err != nil {
		return err
	}
	l := i.locker.New()
	if err := l.Lock("dict:create"); err != nil {
		return err
	}
	defer l.Unlock()
	dict := new(model.Dict)
	err := copier.Copy(&dict, req)
	if err != nil {
		return err
	}

	return i.dictRepo.Save(dict)
}

func (i DictService) UpdateDict(req *UpdateDictReq) error {
	if err := i.validate.Struct(req); err != nil {
		return err
	}
	l := i.locker.New()
	if err := l.Lock("dict:update:" + cast.ToString(req.ID)); err != nil {
		return err
	}
	defer l.Unlock()
	dict, err := i.dictRepo.Detail(req.ID)
	if err != nil {
		return err
	}

	err = copier.Copy(&dict, req)
	if err != nil {
		return err
	}

	return i.dictRepo.Save(dict)
}

type RetrieveDictReq struct {
	router.PageReq
	NameLike string `json:"nameLike" binding:"omitempty,max=64" validate:"omitempty,max=64"` // 名称模糊搜索
	StatusIn string `json:"statusIn" binding:"omitempty" validate:"omitempty"`               // 状态 in 查询
	Type     string `json:"type" binding:"omitempty,max=32" validate:"omitempty,max=32"`     // 字典类型
}

func (i DictService) RetrieveDicts(req *RetrieveDictReq) (count int64, list []model.Dict, err error) {
	if err := i.validate.Struct(req); err != nil {
		return 0, nil, err
	}
	return i.dictRepo.Retrieve(req.Page, req.PageSize, func(tx *gorm.DB) {
		db.AppendWhereFromStruct(tx, req)
		tx.Order("sort asc")
	})
}

// 字典项创建请求
type CreateDictItemReq struct {
	Label    string `json:"label" validate:"required,max=64"`                 // 展示值
	Value    string `json:"value" validate:"required,max=128"`                // 字典值
	Extend   string `json:"extend" validate:"omitempty,max=255"`              // 扩展值
	Status   uint8  `json:"status,string" validate:"required,min=1,max=2"`    // 启用状态
	Sort     uint   `json:"sort,string" validate:"omitempty,min=1,max=65535"` // 排序标记
	DictType string `json:"dictType" validate:"omitempty,max=64"`             // 字典类型
	DictID   uint   `json:"dictID,string" validate:"required"`                // 字典ID
}

// 字典项更新请求
type UpdateDictItemReq struct {
	ID       uint   `json:"id,string" validate:"required"`                    // 字典项ID
	Label    string `json:"label" validate:"omitempty,max=64"`                // 展示值
	Value    string `json:"value" validate:"omitempty,max=128"`               // 字典值
	Extend   string `json:"extend" validate:"omitempty,max=255"`              // 扩展值
	Status   uint8  `json:"status,string" validate:"omitempty,min=1,max=2"`   // 启用状态
	Sort     uint   `json:"sort,string" validate:"omitempty,min=1,max=65535"` // 排序标记
	DictType string `json:"dictType" validate:"omitempty,max=64"`             // 字典类型
	DictID   uint   `json:"dictID,string" validate:"required"`                // 字典ID
}

func (i DictService) CreateItem(req *CreateDictItemReq) error {
	if err := i.validate.Struct(req); err != nil {
		return err
	}
	l := i.locker.New()
	if err := l.Lock("dict:item:create:" + cast.ToString(req.DictID)); err != nil {
		return err
	}
	defer l.Unlock()
	_, err := i.dictRepo.Detail(req.DictID)
	if err != nil {
		return err
	}
	var item model.DictItem
	if err := copier.Copy(&item, req); err != nil {
		return err
	}
	return i.dictRepo.CreateItem(&item)
}

func (i DictService) DeleteItems(dictID any, itemIds ...string) error {

	l := i.locker.New()
	if err := l.Lock("dict:item:delete:" + cast.ToString(dictID) + ":" + strings.Join(itemIds, ",")); err != nil {
		return err
	}
	defer l.Unlock()
	_, err := i.dictRepo.Detail(dictID)
	if err != nil {
		return err
	}
	return i.dictRepo.DeleteItem(itemIds...)
}
func (i DictService) UpdateItem(req *UpdateDictItemReq) error {
	if err := i.validate.Struct(req); err != nil {
		return err
	}

	l := i.locker.New()
	if err := l.Lock("dict:item:update:" + cast.ToString(req.DictID) + ":" + cast.ToString(req.ID)); err != nil {
		return err
	}
	defer l.Unlock()

	_, err := i.dictRepo.Detail(req.DictID)
	if err != nil {
		return err
	}

	var item model.DictItem
	if err := copier.Copy(&item, req); err != nil {
		return err
	}
	item.ID = req.ID

	return i.dictRepo.UpdateItem(&item)
}

func (i DictService) RetrieveItems(dictID any) (int64, []model.DictItem, error) {

	return i.dictRepo.RetrieveItems(dictID)
}

func (i DictService) DeleteDict(ids ...string) error {
	l := i.locker.New()
	if err := l.Lock("dict:delete:" + strings.Join(ids, ",")); err != nil {
		return err
	}
	defer l.Unlock()

	return i.dictRepo.DeleteDict(ids...)
}

func (i DictService) GetDictByType(dictType string) ([]model.DictItem, error) {

	return nil, errors.New("not implemented")
}
