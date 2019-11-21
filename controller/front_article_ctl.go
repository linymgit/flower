package controller

import (
	"flower/entity"
	"flower/http"
	"flower/result"
	"flower/router"
	"flower/service"
	"github.com/valyala/fasthttp"
)

type FrontArticle struct {
}

func init() {
	frontArticle := &FrontArticle{}

	router.AddRoute(
		"/article/category/list",
		http.POST_AND_OPTIONS,
		frontArticle.ListArticlesType,
	)

	router.AddRoute(
		"/article/list",
		http.POST_AND_OPTIONS,
		frontArticle.ListArticles,
	)
}

func (fA *FrontArticle) ListArticlesType(ctx *fasthttp.RequestCtx) (rsp *result.Result) {
	categories, err := service.FrontArticleSrv.ListArticleType()
	if err != nil {
		rsp = result.DatabaseError
		return
	}
	rsp = result.NewSuccess(&entity.FrontArticleTypeRsp{Categories: categories})
	return
}

func (fA *FrontArticle) ListArticles(ctx *fasthttp.RequestCtx, req *entity.FrontArticleListReq) (rsp *result.Result) {
	al, total, err := service.FrontArticleSrv.ListArticles(req)
	if err != nil {
		rsp = result.DatabaseError
		return
	}
	vos := make([]*entity.FrontArticleIntro, len(al))
	for k := range al {
		vos[k] = &entity.FrontArticleIntro{
			Id:      al[k].Id,
			Title:   al[k].Title,
			Author:  al[k].Author,
			Preview: al[k].Preview,
			Summary: al[k].Summary,
		}
	}
	rsp = result.NewSuccess(
		&entity.FrontArticleListRsp{
			Page: &entity.Page{
				PageSize:  req.Page.PageSize,
				PageIndex: req.Page.PageIndex,
				Total:     total,
			},
			List:vos,
		})
	return
}
