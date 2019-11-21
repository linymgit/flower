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

	//获取文章列表
	router.AddRoute(
		"/article/list",
		http.POST_AND_OPTIONS,
		frontArticle.ListArticles,
	)

	//获取文章详情
	router.AddRoute(
		"/article/get",
		http.POST_AND_OPTIONS,
		frontArticle.GetArticle,
	)

	//新闻时间发布导航条
	router.AddRoute(
		"/news/nav",
		http.POST_AND_OPTIONS,
		frontArticle.GetNewsNav,
	)

	router.AddRoute(
		"/news/titles/get",
		http.POST_AND_OPTIONS,
		frontArticle.GetNewsTitles,
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

//获取产品
func (fa *FrontArticle) GetArticle(ctx *fasthttp.RequestCtx, req *entity.FrontGetArticleReq) (rsp *result.Result) {
	p, ok, err := service.ArticleSrv.GetArticle(req.Id)
	if err != nil {
		rsp = result.DatabaseError
		return
	}
	if !ok {
		rsp = result.NewError(result.RequestParamEc, "文章不存在")
		return
	}
	rsp = result.NewSuccess(p)
	return
}

func (fa *FrontArticle) GetNewsNav(ctx *fasthttp.RequestCtx) (rsp *result.Result) {
	navs, err := service.ArticleSrv.GetNewsNav()
	if err != nil {
		rsp = result.DatabaseError
		return
	}
	rsp = result.NewSuccess(navs)
	return
}

func (fa *FrontArticle) GetNewsTitles(ctx *fasthttp.RequestCtx, req *entity.GetNewsTitlesReq) (rsp *result.Result) {
	title, err := service.ArticleSrv.GetNewsTitles(req)
	if err != nil {
		rsp = result.DatabaseError
		return
	}
	rsp = result.NewSuccess(title)
	return
}
