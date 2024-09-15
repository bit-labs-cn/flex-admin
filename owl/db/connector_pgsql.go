package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type pgsqlConnector struct {
	opt *Options
}

func NewPgSqlGetter(opt *Options) *pgsqlConnector {
	return &pgsqlConnector{opt}
}
func (i *pgsqlConnector) Open(cfg *gorm.Config) (*gorm.DB, error) {
	openDb, err := gorm.Open(postgres.Open(i.GetDSN()), cfg)
	return openDb, err
}

func (i *pgsqlConnector) Options() *Options {
	if i.opt != nil {
		return i.opt
	}
	return i.DefaultOptions()
}

func (i *pgsqlConnector) DefaultOptions() *Options {
	return &Options{
		Host:         "localhost",
		Port:         5432,
		Username:     "postgres",
		Password:     "postgres",
		Driver:       Pgsql,
		Database:     "postgres",
		Schema:       "public",
		MaxIdleConns: 10,
		MaxConns:     100,
	}
}

func (i *pgsqlConnector) GetDSN() string {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d TimeZone=%s search_path=%s",
		i.opt.Host,
		i.opt.Username,
		i.opt.Password,
		i.opt.Database,
		i.opt.Port,
		i.opt.TimeZone,
		i.opt.Schema,
	)
	return dsn
}
