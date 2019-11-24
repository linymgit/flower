package service

import (
	"flower/config"
	"flower/entity"
	"flower/entity/gen"
	"flower/mysql"
	"xorm.io/builder"
)

var FrontArticleSrv = &FrontArticleService{}

type FrontArticleService struct {
}

func (fA FrontArticleService) ListArticleType() (categories []*entity.FrontArticleType, err error) {
	categories = make([]*entity.FrontArticleType, 0)

	session := mysql.Db.Table(&gen.ArticleType{}).Asc("id").Cols("id", "type_name")

	pIds := []int{}
	rows, err := mysql.Db.Where("parent_id>?", 0).Distinct("parent_id").Cols("parent_id").Rows(&gen.ArticleType{})
	bean := new(gen.ArticleType)
	for rows.Next() {
		err = rows.Scan(bean)
		if err != nil {
			return
		}
		pIds = append(pIds, bean.ParentId)
	}
	defer rows.Close()
	if len(pIds) == 1 {
		session = session.Where(builder.Neq{"id": pIds[0]})
	}else if len(pIds) > 0 {
		session = session.NotIn("id", pIds)
	}

	err = session.And("id !=?",config.News_Type_Id).Find(&categories)
	return
}

func (fA FrontArticleService) ListArticles(req *entity.FrontArticleListReq) (al []*gen.Article, total int64, err error) {
	session := mysql.Db.NewSession()
	defer session.Close()
	al = make([]*gen.Article, 0)
	if req.TypeId > 0 {
		session = session.Where(builder.Eq{"type_id":req.TypeId})
	}else{
		session = session.Where(builder.Neq{"type_id":config.News_Type_Id})
	}
	// 搜索TODO
	total, err = session.Desc("update_time").Limit(req.Page.PageSize, req.Page.DbPageIndex()).FindAndCount(&al)
	return
}
