package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlConnector struct {
	opt *Options
}

func NewMysqlConnector(opt *Options) *MysqlConnector {
	return &MysqlConnector{opt: opt}
}
func (i *MysqlConnector) Open(cfg *gorm.Config) (*gorm.DB, error) {
	openDb, err := gorm.Open(mysql.Open(i.GetDSN()), cfg)
	return openDb, err
}

func (i *MysqlConnector) Options() *Options {
	if i.opt != nil {
		return i.opt
	}

	return i.DefaultOptions()
}

func (i *MysqlConnector) DefaultOptions() *Options {

	return &Options{
		Host:         "127.0.0.1",
		Port:         3306,
		Username:     "root",
		Password:     "root",
		Driver:       Mysql,
		Database:     "mysql",
		Charset:      "utf8mb4",
		Query:        "parseTime=True&loc=Local&timeout=3000ms",
		MaxIdleConns: 10,
		MaxConns:     100,
	}
}

func (i *MysqlConnector) GetDSN() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&%s",
		i.opt.Username,
		i.opt.Password,
		i.opt.Host,
		i.opt.Port,
		i.opt.Database,
		i.opt.Charset,
		i.opt.Query,
	)
	return dsn
}
