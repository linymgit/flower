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
		http.POST,
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
		http.POST,
		ac.ListArticle,
		handler.CheckAdmin,
	)

	//上下线状态设置
	router.AddRoute(
		"/admin/article/online/change",
		http.POST,
		ac.ChangeOnline,
		handler.CheckAdmin,
	)

	//删除文章
	router.AddRoute(
		"/admin/article/delete",
		http.POST,
		ac.DeleteArticle,
		handler.CheckAdmin,
	)

	//修改（编辑）文章
	router.AddRoute(
		"/admin/article/modify",
		http.POST,
		ac.ModifyArticle,
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

func (ac *Article) EditArticleType(ctx *fasthttp.RequestCtx, req *entity.EditArticleTypeReq) (rsp *result.Result) {
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
func (ac *Article) NewArticle(ctx *fasthttp.RequestCtx, req *entity.NewArticleReq) (rsp *result.Result) {
	articleId, err := service.ArticleSrv.NewArticle(req)
	if err != nil {
		rsp = result.DatabaseError
		return
	}
	rsp = result.NewSuccess(&entity.NewArticleRsp{ArticleId: articleId})
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
			As: as,
		})
	return
}

func (ac *Article) ChangeOnline(ctx *fasthttp.RequestCtx, req *entity.ChangeOnlineReq) (rsp *result.Result) {
	ok, err := service.ArticleSrv.ChangeOnline(req.Id)
	if err != nil {
		rsp = result.DatabaseError
		return
	}
	if !ok {
		rsp = result.NewError(result.RequestParamEc, "无修改的数据")
		return
	}
	rsp = result.NewSuccess("修改成功")
	return
}

func (ac *Article) DeleteArticle(ctx *fasthttp.RequestCtx, req *entity.DeleteArticleReq) (rsp *result.Result) {
	ok, err := service.ArticleSrv.DeleteArticle(req.Id)
	if err != nil {
		rsp = result.DatabaseError
		return
	}
	if !ok {
		rsp = result.NewError(result.RequestParamEc, "产品id不存在")
		return
	}
	rsp = result.NewSuccess("删除成功")
	return
}

func (ac *Article) ModifyArticle(ctx *fasthttp.RequestCtx, req *entity.ModifyArticleReq) (rsp *result.Result) {
	ok, err := service.ArticleSrv.ModifyArticle(req)
	if err != nil {
		rsp = result.DatabaseError
		return
	}
	if !ok {
		rsp = result.NewError(result.RequestParamEc, "产品id不存在")
		return
	}
	rsp = result.NewSuccess("修改成功")
	return
}
