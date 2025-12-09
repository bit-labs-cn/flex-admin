package repository

import (
	"bit-labs.cn/flex-admin/app/model"
	"bit-labs.cn/owl/provider/db"
	"gorm.io/gorm"
)

// 日志仓储接口
type LogRepositoryInterface interface {
	// 保存登录日志
	SaveLoginLog(log *model.LoginLog) error
	// 保存操作日志
	SaveOperationLog(log *model.OperationLog) error
	// 分页查询登录日志
	RetrieveLoginLogs(page, pageSize int, fn func(tx *gorm.DB)) (count int64, list []model.LoginLog, err error)
	// 分页查询操作日志
	RetrieveOperationLogs(page, pageSize int, fn func(tx *gorm.DB)) (count int64, list []model.OperationLog, err error)
}

var _ LogRepositoryInterface = (*LogRepository)(nil)

type LogRepository struct {
	db        *gorm.DB
	loginBase db.BaseRepository[model.LoginLog]
	opernBase db.BaseRepository[model.OperationLog]
}

func NewLogRepository(d *gorm.DB) LogRepositoryInterface {
	return &LogRepository{
		db:        d,
		loginBase: db.NewBaseRepository[model.LoginLog](d),
		opernBase: db.NewBaseRepository[model.OperationLog](d),
	}
}

func (i *LogRepository) SaveLoginLog(log *model.LoginLog) error {
	return i.db.Create(log).Error
}

func (i *LogRepository) SaveOperationLog(log *model.OperationLog) error {
	return i.db.Create(log).Error
}

func (i *LogRepository) RetrieveLoginLogs(page, pageSize int, fn func(tx *gorm.DB)) (count int64, list []model.LoginLog, err error) {
	tx := i.db.Model(&model.LoginLog{})
	if fn != nil {
		fn(tx)
	}
	err = tx.Count(&count).Error
	if err != nil {
		return
	}
	err = tx.Scopes(db.Paginate(page, pageSize)).Order("id DESC").Find(&list).Error
	return
}

func (i *LogRepository) RetrieveOperationLogs(page, pageSize int, fn func(tx *gorm.DB)) (count int64, list []model.OperationLog, err error) {
	tx := i.db.Model(&model.OperationLog{})
	if fn != nil {
		fn(tx)
	}
	err = tx.Count(&count).Error
	if err != nil {
		return
	}
	err = tx.Scopes(db.Paginate(page, pageSize)).Order("id DESC").Find(&list).Error
	return
}
