package service

import (
	"bit-labs.cn/flex-admin/app/model"
	"bit-labs.cn/flex-admin/app/repository"
	"bit-labs.cn/owl/provider/db"
	"bit-labs.cn/owl/provider/redis"
	"bit-labs.cn/owl/provider/router"
	"github.com/spf13/cast"
	"gorm.io/gorm"

	"strings"

	"github.com/jinzhu/copier"
)

type CreateDictReq struct {
	Name   string `json:"name" validate:"required,max=32"`
	Type   string `json:"type"`
	Status uint8  `json:"status,string" validate:"required,min=0,max=1"`
	Desc   string `json:"desc"`
	Sort   uint8  `json:"sort,string" validate:"required,min=0,max=255"`
}

type UpdateDictReq struct {
	ID uint `json:"id,string"`
	CreateDictReq
}

type DictService struct {
	dictRepo repository.DictRepositoryInterface
	locker   redis.LockerFactory
}

func NewDictService(dictRepo repository.DictRepositoryInterface, locker redis.LockerFactory) *DictService {
	return &DictService{
		dictRepo: dictRepo,
		locker:   locker,
	}
}
func (i DictService) CreateDict(req *CreateDictReq) error {
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
	NameLike string `json:"nameLike" binding:"omitempty,max=64"`
	StatusIn string `json:"statusIn" binding:"omitempty"`
	Type     string `json:"type" binding:"omitempty,max=32"`
}

func (i DictService) RetrieveDicts(req *RetrieveDictReq) (count int64, list []model.Dict, err error) {
	return i.dictRepo.Retrieve(req.Page, req.PageSize, func(tx *gorm.DB) {
		db.AppendWhereFromStruct(tx, req)
		tx.Order("sort asc")
	})
}

func (i DictService) CreateItem(item *model.DictItem) error {
	l := i.locker.New()
	if err := l.Lock("dict:item:create:" + cast.ToString(item.DictID)); err != nil {
		return err
	}
	defer l.Unlock()
	_, err := i.dictRepo.Detail(item.DictID)
	if err != nil {
		return err
	}
	return i.dictRepo.CreateItem(item)
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
func (i DictService) UpdateItem(item *model.DictItem) error {
	l := i.locker.New()
	if err := l.Lock("dict:item:update:" + cast.ToString(item.DictID) + ":" + cast.ToString(item.ID)); err != nil {
		return err
	}
	defer l.Unlock()
	_, err := i.dictRepo.Detail(item.DictID)
	if err != nil {
		return err
	}
	return i.dictRepo.UpdateItem(item)
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
	//TODO implement me
	panic("implement me")
}
