package repository

import (
	"context"
	"errors"

	"bit-labs.cn/flex-admin/app/model"
	"bit-labs.cn/owl/contract"
	"gorm.io/gorm"
)

var ErrAppVersionNotFound = errors.New("app版本不存在")

type AppVersionRepositoryInterface interface {
	Latest(apkType *int32) (*model.AppVersion, error)
	contract.WithContext[AppVersionRepositoryInterface]
}

var _ AppVersionRepositoryInterface = (*AppVersionRepository)(nil)

type AppVersionRepository struct {
	db  *gorm.DB
	ctx context.Context
}

func NewAppVersionRepository(tx *gorm.DB) AppVersionRepositoryInterface {
	return &AppVersionRepository{db: tx}
}

func (i *AppVersionRepository) WithContext(ctx context.Context) AppVersionRepositoryInterface {
	i.db = i.db.WithContext(ctx)
	i.ctx = ctx
	return i
}

func (i *AppVersionRepository) Latest(apkType *int32) (*model.AppVersion, error) {
	var v model.AppVersion

	tx := i.db.Model(&model.AppVersion{})

	if apkType != nil {
		tx = tx.Where("apk_type = ?", *apkType)
	}

	err := tx.Order("id desc").First(&v).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrAppVersionNotFound
	}
	return &v, err
}
