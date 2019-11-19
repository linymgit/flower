package service

import (
	"flower/entity"
	"flower/entity/gen"
	"flower/entity/state"
	"flower/mysql"
)

var IndexSrv = &IndexService{}

type IndexService struct {
}

func (i *IndexService) ListIndexAd(query *entity.IndexReq) (ads []*gen.Ad, total int64, err error) {
	ads = make([]*gen.Ad, 0)
	session := mysql.Db.NewSession()
	defer session.Close()
	total, err = session.Where("postion_id=?", state.AdIndex).Asc("save_time").Limit(query.Page.PageSize, query.Page.DbPageIndex()).FindAndCount(&ads)
	return
}

func (i *IndexService) ListIndexProduct(query *entity.IndexReq) (ps []*gen.Product, total int64, err error) {
	ps = make([]*gen.Product, 0)
	session := mysql.Db.NewSession()
	defer session.Close()
	total, err = session.Where("index_show=?", state.IndexShow).Asc("update_time").Limit(query.Page.PageSize, query.Page.DbPageIndex()).Cols("id", "cover_url", "name", "summary").FindAndCount(&ps)
	return
}
