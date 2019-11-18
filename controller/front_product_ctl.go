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
		http.POST,
		fP.ListProduct,
	)

	router.AddRoute(
		"/product/list",
		http.OPTIONS,
		fP.ListProduct,
	)

	router.AddRoute(
		"/product/category/list",
		http.POST,
		fP.ListCategory,
	)

	//根据产品详情
	router.AddRoute(
		"/product/get",
		http.POST,
		fP.GetProduct,
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
	rsp = result.NewSuccess(&entity.FrontListCategoryRsp{Categories: categories})
	return
}

//获取产品
func (fP *FrontProduct) GetProduct(ctx *fasthttp.RequestCtx, req *entity.FrontGetProductReq) (rsp *result.Result) {
	p, ok, err := service.FrontProdSrv.GetProduct(req.Id)
	if err != nil {
		rsp = result.DatabaseError
		return
	}
	if !ok {
		rsp = result.NewError(result.RequestParamEc, "产品不存在")
		return
	}
	rsp = result.NewSuccess(&entity.FrontGetProductRsp{Product: p})
	return
}
