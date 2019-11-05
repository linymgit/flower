package service

import (
	"flower/entity/gen"
	"flower/mysql"
)

var AccSrv = &AccountSrv{}

type AccountSrv struct {
}

func (a *AccountSrv) GetAccountByName(name string) (exist bool, account *gen.Account, err error) {
	account = &gen.Account{}
	exist, err = mysql.Db.Where("name = ?", name).Get(account)
	return
}

func (a *AccountSrv) GetAccountPwById(id int64) (exist bool, password string, err error) {
	exist, err = mysql.Db.Table(&gen.Account{}).ID(id).Cols("password").Get(&password)
	return
}

func (a *AccountSrv) UpdateAccountPwById(id int64, password string) (err error) {
	affected, err := mysql.Db.ID(id).Cols("password").Update(&gen.Account{
		Id:       id,
		Password: password,
	})
	if err != nil {
		// TODO
	}
	if affected != 1 {
		// TODO
	}
	return
}
