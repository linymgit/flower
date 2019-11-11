package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
	"xorm.io/xorm"
)

var Db *xorm.Engine

func Init(dbUrl string) (err error) {
	Db, err = xorm.NewEngine("mysql", dbUrl)
	Db.ShowSQL(true)
	Db.SetConnMaxLifetime(1 * time.Hour)
	Db.SetMaxIdleConns(500)
	Db.SetMaxOpenConns(1000)
	if err != nil {
		// TODO
	}
	if err = Db.Ping(); err != nil {
		// TODO
		return
	}
	return
}
