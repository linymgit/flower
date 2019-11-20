package service

import (
	"flower/entity"
	"flower/entity/gen"
	"flower/entity/state"
	"flower/mysql"
	"time"
	"xorm.io/builder"
)

var timeTemplate = "2006-01-02 15:04:05" //常规类型

var CrmSrv = &CrmServer{}

type CrmServer struct {
}

func (c *CrmServer) ListCrm(query *entity.CrmListReq) (crms []*gen.Crm, total int64, err error) {
	crms = make([]*gen.Crm, 0)
	cond := builder.NewCond()
	cond = cond.And(builder.Eq{"deleted": state.CrmNormal})
	if query.Name != "" {
		cond = cond.And(builder.Eq{"name": &query.Name})
	}
	if query.EndTime > 0 && query.BeginTime > 0 {
		begin := time.Unix(query.BeginTime, 0).Add(-24 * time.Hour).Format(timeTemplate)
		end := time.Unix(query.EndTime, 0).Add(24 * time.Hour).Format(timeTemplate)
		cond = cond.And(builder.Lt{"save_time": &end})
		cond = cond.And(builder.Gt{"save_time": &begin})
	}
	session := mysql.Db.NewSession()
	defer session.Close()
	total, err = session.Where(cond).Asc("save_time").Limit(query.Page.PageSize, query.Page.DbPageIndex()).FindAndCount(&crms)
	return
}

func (c *CrmServer) DeleteCrmById(id int64) (ok bool, err error) {
	i, err := mysql.Db.ID(id).Cols("deleted").Update(&gen.Crm{
		Deleted: state.CrmDeleted,
	})
	if i > 0 {
		ok = true
	} else {
		//TODO
	}
	return
}

func (c *CrmServer) InsertCrm(query *entity.AddCrmReq) (ok bool, err error) {
	affected, err := mysql.Db.Cols("name", "email", "phone", "official_web", "company", "message", "deleted").InsertOne(&gen.Crm{
		Name:        query.Name,
		Email:       query.Email,
		Phone:       query.Phone,
		OfficialWeb: query.OfficialWeb,
		Company:     query.Company,
		Message:     query.Message,
		Deleted:     state.CrmNormal,
	})
	ok = affected == 1
	return
}
