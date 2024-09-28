package service

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type BaseService[T schema.Tabler] struct {
	db *gorm.DB
}

func NewBaseService[T schema.Tabler](db *gorm.DB) BaseService[T] {
	return BaseService[T]{
		db: db,
	}
}

// Delete 删除数据
func (i *BaseService[T]) Delete(id any, fn func(db *gorm.DB) (msg string, b bool)) error {
	model := new(T)

	if fn != nil {
		msg, b := fn(i.db)
		if !b {
			return errors.New(msg)
		}
	}

	db := i.db.Model(model).Where("id", id).Delete(nil)
	return db.Error
}

// Detail 获取详情
func (i *BaseService[T]) Detail(id any) (*T, error) {
	model := new(T)
	db := i.db.Model(model).Where("id", id).First(model)
	return model, db.Error
}

// ChangeStatus 更改状态，我们经常会需要单独的更改状态，比如禁用，启用等。
func (i *BaseService[T]) ChangeStatus(req *ChangeStatus) error {
	return i.db.Model(new(T)).
		Where("id", req.ID).
		Update("status", req.Status).Error
}
