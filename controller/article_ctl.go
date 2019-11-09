package controller

import (
	"flower/entity"
	"flower/handler"
	"flower/http"
	"flower/result"
	"flower/router"
	"flower/service"
	"github.com/valyala/fasthttp"
)

type Article struct {
}

func init() {

	ac := &Article{}

	// --------------文章分类--------------------

	router.AddRoute(
		"/admin/article/category/new",
		http.POST,
		ac.NewArticleType,
		handler.CheckAdmin,
	)

	router.AddRoute(
		"/admin/article/category/tree",
		http.GET,
		ac.GetActicleCategoryTree,
		handler.CheckAdmin,
	)

	router.AddRoute(
		"/admin/article/category/list",
		http.GET,
		ac.ListArticleType,
		handler.CheckAdmin,
	)

	//编辑文章分类
	router.AddRoute(
		"/admin/article/category/edit",
		http.POST,
		ac.EditArticleType,
		handler.CheckAdmin,
	)

	// --------------文章--------------------
	router.AddRoute(
		"/admin/article/new",
		http.POST,
		ac.NewArticle,
		handler.CheckAdmin,
	)

	router.AddRoute(
		"/admin/article/list",
		http.GET,
		ac.ListArticle,
		handler.CheckAdmin,
	)
}

//获取产品分类列表
func (ac *Article) ListArticleType(ctx *fasthttp.RequestCtx, req *entity.ListArticleTypeReq) (resp *result.Result) {
	ats, total, err := service.ArticleSrv.ListArticleType(req)
	if err != nil {
		resp = result.DatabaseError
		return
	}
	var page *entity.Page
	if req.Page != nil {
		page = &entity.Page{
			PageSize:  req.Page.PageSize,
			PageIndex: req.Page.PageIndex,
			Total:     total,
		}
	}
	resp = result.NewSuccess(
		&entity.ListArticleTypeRsp{
			Page: page,
			Ats:  ats,
		})
	return
}

func (ac *Article) NewArticleType(ctx *fasthttp.RequestCtx, req *entity.NewArticleTypeReq) (rsp *result.Result) {
	isExistName, isExistParent, ok, err := service.ArticleSrv.NewArticleType(req)
	if err != nil {
		rsp = result.DatabaseError
		return
	}
	if isExistName {
		rsp = result.NewError(result.RequestParamEc, "分类名称已存在")
		return
	}
	if !isExistParent {
		rsp = result.NewError(result.RequestParamEc, "上级id不存在")
		return
	}
	rsp = result.NewSuccess(ok)
	return
}

func (ac *Article) GetActicleCategoryTree(ctx *fasthttp.RequestCtx) (rsp *result.Result) {
	tree, err := service.ArticleSrv.GetArticleCategoryTree()
	if err != nil {
		rsp = result.DatabaseError
		return
	}
	rsp = result.NewSuccess(tree)
	return
}

func (ac *Article) EditArticleType(ctx *fasthttp.RequestCtx,req *entity.EditArticleTypeReq) (rsp *result.Result) {
	ok, existParent, err := service.ArticleSrv.EditArticle(req)
	if err != nil {
		rsp = result.DatabaseError
		return
	}
	if !existParent {
		rsp = result.NewError(result.RequestParamEc, "上级分类不存在")
		return
	}
	rsp = result.NewSuccess(ok)
	return
}

//创建文章
func (ac *Article) NewArticle(ctx *fasthttp.RequestCtx,req *entity.NewArticleReq) (rsp *result.Result){
	articleId, err := service.ArticleSrv.NewArticle(req)
	if err != nil {
		rsp = result.DatabaseError
		return
	}
	rsp = result.NewSuccess(&entity.NewArticleRsp{ArticleId:articleId})
	return
}

//获取文章
func (ac *Article) ListArticle(ctx *fasthttp.RequestCtx, req *entity.ListArticleReq) (resp *result.Result) {
	as, total, err := service.ArticleSrv.ListArticle(req)
	if err != nil {
		resp = result.DatabaseError
		return
	}
	resp = result.NewSuccess(
		&entity.ListArticleRsp{
			Page: &entity.Page{
				PageSize:  req.Page.PageSize,
				PageIndex: req.Page.PageIndex,
				Total:     total,
			},
			As:as,
		})
	return
}