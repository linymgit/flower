package service

import (
	"flower/entity"
	"flower/entity/gen"
	"flower/entity/state"
	"flower/mysql"
	"xorm.io/builder"
)

var FrontProdSrv = &FrontProdService{}

type FrontProdService struct {
}

func (fP FrontProdService) ListCategory() (categories []*entity.FrontCategory, err error) {
	categories = make([]*entity.FrontCategory, 0)

	session := mysql.Db.Table(&gen.ProductCategory{}).Asc("id").Where(builder.Eq{"states": state.ProdCategoryShow})

	pIds := []int{}
	rows, err := mysql.Db.Where("parent_id>?", 0).Cols("parent_id").Rows(&gen.ProductCategory{})
	if err != nil {
	}
	defer rows.Close()
	bean := new(gen.ProductCategory)
	for rows.Next() {
		err = rows.Scan(bean)
		if err != nil {
			return
		}
		pIds = append(pIds, bean.ParentId)
	}
	if len(pIds) > 0 {
		session = session.NotIn("id", pIds)
	}

	err = session.Cols("id", "name").Find(&categories)
	return
}

func (fP FrontProdService) ListProduct(query *entity.FrontListProductReq) (ps []*gen.Product, total int64, err error) {
	session := mysql.Db.NewSession()
	defer session.Close()
	ps = make([]*gen.Product, 0)
	cond := builder.NewCond()
	desc := []string{}
	if query.IsHot {
		desc = append(desc, "heat")
	}
	if query.IsNew {
		desc = append(desc, "save_time")
	}
	if len(desc) <= 0 {
		desc = append(desc, "update_time")
	}
	if query.IsIndexShow {
		cond = cond.And(builder.Eq{"index_show": state.IndexShow})
	}
	if query.CategoryId > 0 {
		cond = cond.And(builder.Eq{"category_id": query.CategoryId})
	}
	// 搜索TODO
	total, err = session.Where(cond).Desc(desc...).Limit(query.Page.PageSize, query.Page.DbPageIndex()).FindAndCount(&ps)
	return
}

func (fP FrontProdService) GetProduct(id int64) (p *gen.Product, ok bool, err error) {
	session := mysql.Db.NewSession()
	defer session.Close()
	p = &gen.Product{}
	ok, err = mysql.Db.ID(id).Get(p)
	if err != nil {
		return
	}
	return
}
