package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type sqliteGetter struct {
	opt *Options
}

func NewSqliteGetter(opt *Options) *sqliteGetter {
	return &sqliteGetter{opt}
}
func (i *sqliteGetter) Open(cfg *gorm.Config) (*gorm.DB, error) {
	openDb, err := gorm.Open(sqlite.Open(i.GetDSN()), cfg)
	return openDb, err

}
func (i *sqliteGetter) Options() *Options {
	if i.opt != nil {
		return i.opt
	}

	return i.DefaultOptions()
}

func (i *sqliteGetter) DefaultOptions() *Options {
	return &Options{
		Host:         "owlSqlite.db",
		Username:     "root",
		Password:     "root",
		Database:     "main",
		Driver:       Sqlite,
		MaxIdleConns: 10,
		MaxConns:     100,
	}
}

func (i *sqliteGetter) GetDSN() string {
	return i.opt.Host
}
