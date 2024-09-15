package db

import "gorm.io/gorm"

type Connector interface {
	Open(cfg *gorm.Config) (*gorm.DB, error) // 连接数据库
	Options() *Options                       // 返回配置
	DefaultOptions() *Options                // 返回默认配置
	GetDSN() string                          // 获取 DSN
}
