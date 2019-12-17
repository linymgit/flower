package mysql

import (
	"flower/config"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"xorm.io/xorm"
)

var Db *xorm.Engine

func Init(dbUrl string) (err error) {
	Db, err = xorm.NewEngine("mysql", dbUrl)
	Db.ShowSQL(config.Conf.MysqlConfig.ShowSQL)
	Db.SetConnMaxLifetime(time.Duration(int(time.Minute) * config.Conf.MysqlConfig.ConnMaxLifetime))
	Db.SetMaxIdleConns(config.Conf.MysqlConfig.MaxIdleConns)
	Db.SetMaxOpenConns(config.Conf.MysqlConfig.MaxOpenConns)
	if err != nil {
		// TODO
	}
	if err = Db.Ping(); err != nil {
		// TODO
		return
	}
	return
}
