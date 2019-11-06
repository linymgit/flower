package service

import (
	"flower/entity"
	"flower/entity/gen"
	"flower/mysql"
)

var ProdSrv = &ProdService{}

type ProdService struct {
}

func (p *ProdService) ListProductCategory(query *entity.ProductCategoryReq) (pcs []*gen.ProductCategory, total int64, err error) {
	pcs = make([]*gen.ProductCategory, 0)
	session := mysql.Db.NewSession()
	defer session.Close()
	total, err = session.Desc("weight").Limit(query.Page.PageSize, query.Page.DbPageIndex()).FindAndCount(&pcs)
	return
}
