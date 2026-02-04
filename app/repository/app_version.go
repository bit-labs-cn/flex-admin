package repository

import (
	"errors"

	"bit-labs.cn/flex-admin/app/model"
	"gorm.io/gorm"
)

var ErrAppVersionNotFound = errors.New("app版本不存在")

type AppVersionRepositoryInterface interface {
	Latest(apkType *int32) (*model.AppVersion, error)
}

var _ AppVersionRepositoryInterface = (*AppVersionRepository)(nil)

type AppVersionRepository struct {
	db *gorm.DB
}

func NewAppVersionRepository(tx *gorm.DB) AppVersionRepositoryInterface {
	return &AppVersionRepository{db: tx}
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
