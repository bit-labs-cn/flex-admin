package owl

import (
	"context"
	"errors"
	"github.com/guoliang1994/gin-flex-admin/owl/db"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Repository[T schema.Tabler] struct {
	ctx   context.Context
	db    *gorm.DB
	check func(T) map[string]interface{}
}

func NewRepository[T schema.Tabler](db *gorm.DB) *Repository[T] {
	return &Repository[T]{db: db}
}

// Create 创建
func (i *Repository[T]) Create(value T) error {
	newDb := i.db.Session(&gorm.Session{})
	if i.check != nil {
		var c int64
		condition := i.check(value)
		newDb.Model(new(T)).Where(condition).Count(&c)
		if c > 0 {
			return errors.New("数据已存在")
		}
		if newDb.Error != nil {
			return i.db.Error
		}
	}
	return newDb.Create(value).Error
}

func (i *Repository[T]) Delete(ids ...int64) error {
	return i.db.Where("id in ?", ids).Delete(new(T)).Error
}

func (i *Repository[T]) Update(id int64, value T) error {
	if i.check != nil {
		var c int64
		condition := i.check(value)
		newDb := i.db.Model(new(T)).Where("id != ?", id).Where(condition)
		newDb.Count(&c)
		if c > 0 {
			return errors.New("数据已存在")
		}
		if newDb.Error != nil {
			return i.db.Error
		}
	}
	var c int64
	i.db.Where("id", id).Count(&c)
	if c == 0 {
		return errors.New("数据不存在")
	}
	return i.db.Save(value).Error
}

// Retrieve 分页查询
func (i *Repository[T]) Retrieve(page, pageSize int, condition map[string]interface{}) (int64, []T) {
	var list []T
	i.db.Scopes(db.Paginate(page, pageSize)).Where(condition).Find(&list)
	var count int64
	i.db.Where(condition).Count(&count)
	return count, list
}

// List 获取所有数据
func (i *Repository[T]) List(condition any) []T {
	var list []T
	i.db.Where(condition).Find(&list)
	return list
}

func (i *Repository[T]) Detail(id int64) T {
	var d T
	i.db.Where("id", id).Find(&d)
	return d
}

func (i *Repository[T]) UniqueCheckFn(fn func(T) map[string]interface{}) *Repository[T] {
	i.check = fn
	return i
}
