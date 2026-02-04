package service

import (
	"errors"

	"bit-labs.cn/flex-admin/app/model"
	"bit-labs.cn/flex-admin/app/repository"
)

var ErrNoAvailableAppVersion = errors.New("暂无可用版本")

type AppVersionService struct {
	repo repository.AppVersionRepositoryInterface
}

func NewAppVersionService(repo repository.AppVersionRepositoryInterface) *AppVersionService {
	return &AppVersionService{repo: repo}
}

func (i *AppVersionService) Latest(apkType *int32) (*model.AppVersion, error) {
	v, err := i.repo.Latest(apkType)
	if errors.Is(err, repository.ErrAppVersionNotFound) {
		return nil, ErrNoAvailableAppVersion
	}
	return v, err
}
