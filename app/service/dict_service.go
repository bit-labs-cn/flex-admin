package service

import (
	"bit-labs.cn/flex-admin/app/model"
	"bit-labs.cn/flex-admin/app/repository"
	"bit-labs.cn/owl/contract"
	"bit-labs.cn/owl/provider/db"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type CreateDictReq struct {
	model.Dict
}

type UpdateDictReq struct {
	ID uint `json:"id,string"` // 主键
	CreateDictReq
}

type DictService struct {
	dictRepo repository.DictRepositoryInterface
}

func NewDictService(dictRepo repository.DictRepositoryInterface) *DictService {
	return &DictService{
		dictRepo: dictRepo,
	}
}
func (i DictService) CreateDict(req *CreateDictReq) error {
	dict := new(model.Dict)
	err := copier.Copy(&dict, req)
	if err != nil {
		return err
	}

	return i.dictRepo.Save(dict)
}

func (i DictService) UpdateDict(req *UpdateDictReq) error {
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
	contract.PageReq
	NameLike string `json:"name"`   // 名称模糊搜索
	StatusIn string `json:"status"` // 状态in查询
	Type     string `json:"type"`   // 类型
}

func (i DictService) RetrieveDicts(req *RetrieveDictReq) (count int64, list []model.Dict, err error) {
	return i.dictRepo.Retrieve(req.Page, req.PageSize, func(tx *gorm.DB) {
		db.AppendWhereFromStruct(tx, req)
		tx.Order("sort asc")
	})
}

func (i DictService) CreateItem(item *model.DictItem) error {
	_, err := i.dictRepo.Detail(item.DictID)
	if err != nil {
		return err
	}
	return i.dictRepo.CreateItem(item)
}

func (i DictService) DeleteItems(dictID any, itemIds ...string) error {
	_, err := i.dictRepo.Detail(dictID)
	if err != nil {
		return err
	}
	return i.dictRepo.DeleteItem(itemIds...)
}
func (i DictService) UpdateItem(item *model.DictItem) error {
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
	return i.dictRepo.DeleteDict(ids...)
}

func (i DictService) GetDictByType(dictType string) ([]model.DictItem, error) {
	//TODO implement me
	panic("implement me")
}
