package service

import (
	"flower/entity"
	"flower/entity/gen"
	"flower/mysql"
)

var FrontArticleSrv = &FrontArticleService{}

type FrontArticleService struct {
}

func (fA FrontArticleService) ListArticleType() (categories []*entity.FrontArticleType, err error) {
	categories = make([]*entity.FrontArticleType, 0)

	session := mysql.Db.Table(&gen.ArticleType{}).Asc("id").Cols("id", "type_name")

	pIds := []int{}
	rows, err := mysql.Db.Where("parent_id>?", 0).Cols("parent_id").Rows(&gen.ArticleType{})
	bean := new(gen.ArticleType)
	for rows.Next() {
		err = rows.Scan(bean)
		if err != nil {
			return
		}
		pIds = append(pIds, bean.ParentId)
	}
	defer rows.Close()
	if len(pIds) > 0 {
		session = session.NotIn("id", pIds)
	}

	err = session.Find(&categories)
	return
}
