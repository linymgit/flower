package service

import (
	"flower/entity"
	"flower/entity/gen"
	"flower/mysql"
)

var FrontArticleSrv = &FrontArticleService{}

type FrontArticleService struct {
}

func (fA FrontArticleService) ListArticleType() (categories []*entity.FrontArticleType, err error){
	categories = make([]*entity.FrontArticleType, 0)
	err = mysql.Db.Table(&gen.ArticleType{}).Asc("sort").Cols("id","type_name").Find(&categories)
	return
}

