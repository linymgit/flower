package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var Db *xorm.Engine

func Init(dbUrl string) (err error) {
	Db, err = xorm.NewEngine("mysql", dbUrl)
	if err != nil {
		// TODO
	}
	if err = Db.Ping(); err != nil {
		// TODO
		return
	}
	return
}
