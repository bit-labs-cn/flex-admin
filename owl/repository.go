package owl

import (
	"github.com/guoliang1994/gin-flex-admin/owl/db"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Repository[T schema.Tabler] struct {
	db *gorm.DB
}

func NewRepository[T schema.Tabler](db *gorm.DB) *Repository[T] {
	return &Repository[T]{db: db}
}

// Create 创建
func (i *Repository[T]) Create(value T) error {
	return i.db.Create(value).Error
}

func (i *Repository[T]) Delete(ids ...int64) error {
	return i.db.Where("id in ?", ids).Delete(new(T)).Error
}

func (i *Repository[T]) Update(id int64, value T) error {
	return i.db.Where("id", id).Save(value).Error
}

// Retrieve 分页查询
func (i *Repository[T]) Retrieve(page, pageSize int, condition any) (int64, []T) {
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
