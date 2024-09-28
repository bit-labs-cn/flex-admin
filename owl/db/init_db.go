package db

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

type Options struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Host         string `json:"host"`
	Port         int    `json:"port"`
	Driver       Driver `json:"driver"` // 数据库类型
	Database     string `json:"database"`
	Schema       string `json:"schema"`
	Charset      string `json:"charset"`
	Query        string `json:"query"`
	MaxIdleConns int    `json:"max-idle-conns"`
	MaxConns     int    `json:"max-conns"`
	TimeZone     string `json:"timezone"`
}

type CustomReplacer struct {
	f func(string) string
}

func (r CustomReplacer) Replace(name string) string {
	return r.f(name)
}

func InitDB(opt *Options, plugins ...gorm.Plugin) *gorm.DB {

	var dbGetter Connector
	switch opt.Driver {
	case Mysql:
		dbGetter = NewMysqlConnector(opt)
	case Pgsql:
		dbGetter = NewPgSqlGetter(opt)
	case Sqlite:
		dbGetter = NewSqliteGetter(opt)
	default:
		panic("不支持的数据库类型")
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      false,       // 禁用彩色打印
		},
	)

	gormCfg := &gorm.Config{
		PrepareStmt:                              false,
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			//单数表名
			SingularTable: true,
			TablePrefix:   "admin_",
		},
		Logger:                 newLogger,
		SkipDefaultTransaction: true,
	}

	var openDb *gorm.DB
	var err error
	openDb, err = dbGetter.Open(gormCfg)

	if err != nil {
		panic("数据库连接失败，请检查数据库是否启动，配置是否错误" + err.Error())
	}

	for _, plugin := range plugins {
		err = openDb.Use(plugin)
		if err != nil {
			panic("应用数据库插件失败" + err.Error())
		}
	}
	sqlDB, err := openDb.DB()
	if err != nil {
		panic("数据库连接失败，请检查数据库是否启动，配置是否错误" + err.Error())
	}
	// 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(opt.MaxIdleConns)

	// 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(opt.MaxConns)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	return openDb
}
