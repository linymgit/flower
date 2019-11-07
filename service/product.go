package service

import (
	"flower/entity"
	"flower/entity/gen"
	"flower/mysql"
)

var ProdSrv = &ProdService{}

type ProdService struct {
}

func (p *ProdService) ListProductCategory(query *entity.ListProductCategoryReq) (pcs []*gen.ProductCategory, total int64, err error) {
	pcs = make([]*gen.ProductCategory, 0)
	session := mysql.Db.NewSession()
	defer session.Close()
	total, err = session.Asc("sort").Limit(query.Page.PageSize, query.Page.DbPageIndex()).FindAndCount(&pcs)
	return
}

func (p *ProdService) GetProductCategoryTree()  (err error){
	rows, err := mysql.Db.Rows(&gen.ProductCategory{})
	if err != nil {
	}
	defer rows.Close()
	bean := new(gen.ProductCategory)
	for rows.Next() {
		err = rows.Scan(bean)
		if err != nil {
			return
		}

	}
	return
}

func (p *ProdService) NewProductCategory(query *entity.NewProdCategoryReq) (isExistName bool, ok bool, err error) {
	p.GetProductCategoryTree()
	// 分类名称唯一校验
	isExistName, err = mysql.Db.Where("name = ?", query.Name).Exist(&gen.ProductCategory{})
	if err != nil {
		return
	}
	if isExistName {
		return
	}
	affected, err := mysql.Db.Cols("parent_id","name","desc","states","level","sort").InsertOne(&gen.ProductCategory{
		ParentId:   query.ParentId,
		Name:       query.Name,
		Desc:       query.Desc,
		Level:      0,
		Sort:       query.Sort,
	})
	if err != nil {
		//TODO
		return
	}
	ok = affected == 1
	return
}
