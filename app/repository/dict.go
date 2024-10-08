package repository

import (
	"bit-labs.cn/gin-flex-admin/app/model"
	"bit-labs.cn/owl/db"
	"errors"
	"gorm.io/gorm"
)

var ErrDictExists = errors.New("字典已存在")
var ErrDictNotExists = errors.New("字典不存在")
var ErrDictItemExists = errors.New("字典项已存在")

type DictRepositoryInterface interface {
	Save(dict *model.Dict) error
	DeleteDict(ids ...string) error
	Unique(id uint, name string, Type string) (*model.Dict, bool)
	Detail(id any) (*model.Dict, error)
	Retrieve(page, pageSize int, fn func(db *gorm.DB)) (count int64, list []model.Dict, err error)

	CreateItem(item *model.DictItem) error
	DeleteItem(itemIds ...string) error
	UpdateItem(item *model.DictItem) error
	RetrieveItems(dictID any) (count int64, list []model.DictItem, err error)
}

var _ DictRepositoryInterface = (*DictRepository)(nil)

type DictRepository struct {
	db *gorm.DB
	db.BaseRepository[model.Dict]
}

func NewDictRepository(d *gorm.DB) DictRepositoryInterface {
	return &DictRepository{
		db:             d,
		BaseRepository: db.NewBaseRepository[model.Dict](d),
	}
}
func (i DictRepository) Save(dict *model.Dict) error {
	_, exists := i.Unique(dict.ID, dict.Name, dict.Type)
	if exists {
		return ErrDictExists
	}

	err := i.db.Where("id", dict.ID).Save(&dict).Error
	return err
}

func (i DictRepository) Detail(id any) (*model.Dict, error) {
	var m model.Dict
	err := i.db.Where("id = ?", id).First(&m).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &m, ErrDictNotExists
	}

	return &m, err
}

func (i DictRepository) DeleteDict(ids ...string) error {
	return i.db.Model(&model.Dict{}).Where("id in ?", ids).Delete(nil).Error
}

func (i DictRepository) CreateItem(item *model.DictItem) error {
	return i.db.Create(item).Error
}

func (i DictRepository) DeleteItem(itemIds ...string) error {
	return i.db.Delete(&model.DictItem{}, itemIds).Error
}
func (i DictRepository) UpdateItem(item *model.DictItem) error {
	return i.db.Updates(item).Error
}
func (i DictRepository) RetrieveItems(dictID any) (count int64, list []model.DictItem, err error) {
	i.db.Where("dict_id", dictID).Order("sort asc").Find(&list)
	return 100, list, nil
}
func (i DictRepository) Unique(id uint, name string, Type string) (*model.Dict, bool) {
	return i.BaseRepository.Unique(id, func(db *gorm.DB) {
		db.Where("name = ? or type = ?", name, Type)
	})
}
