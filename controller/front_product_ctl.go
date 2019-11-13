package controller

import (
	"flower/entity"
	"flower/http"
	"flower/result"
	"flower/router"
	"flower/service"
	"github.com/valyala/fasthttp"
)

type FrontProduct struct {
}

func init() {
	fP := &FrontProduct{}

	router.AddRoute(
		"/product/list",
		http.GET,
		fP.ListProduct,
	)

	router.AddRoute(
		"/product/category/list",
		http.GET,
		fP.ListCategory,
	)
}

//获取产品分类列表
func (fP *FrontProduct) ListProduct(ctx *fasthttp.RequestCtx, req *entity.FrontListProductReq) (rsp *result.Result) {
	ps, total, err := service.FrontProdSrv.ListProduct(req)
	if err != nil {
		rsp = result.DatabaseError
		return
	}
	rsp = result.NewSuccess(
		&entity.FrontListProductRsp{
			Page: &entity.Page{
				PageSize:  req.Page.PageSize,
				PageIndex: req.Page.PageIndex,
				Total:     total,
			},
			Ps: ps,
		})
	return
}

//获取产品列表
func (fP *FrontProduct) ListCategory(ctx *fasthttp.RequestCtx) (rsp *result.Result) {
	categories, err := service.FrontProdSrv.ListCategory()
	if err != nil {
		rsp = result.DatabaseError
		return
	}
	rsp = result.NewSuccess(&entity.FrontListCategoryRsp{Categories:categories})
	return
}