package controller

import (
	"flower/entity"
	"flower/handler"
	"flower/http"
	"flower/log"
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
		http.POST_AND_OPTIONS,
		fP.ListProduct,
	)

	router.AddRoute(
		"/product/category/list",
		http.POST_AND_OPTIONS,
		fP.ListCategory,
	)

	//根据产品详情
	router.AddRoute(
		"/product/get",
		http.POST_AND_OPTIONS,
		fP.GetProduct,
		handler.Pv4Product,
	)
}

//获取产品分类列表
func (fP *FrontProduct) ListProduct(ctx *fasthttp.RequestCtx, req *entity.FrontListProductReq) (rsp *result.Result) {
	ps, total, err := service.FrontProdSrv.ListProduct(req)
	if err != nil {
		rsp = result.DatabaseError
		log.ErrorF("FrontProduct.ListProduct->[%v]", err)
		return
	}
	vos := make([]*entity.FrontListProductVo, len(ps))
	for k := range ps {
		vos[k] = &entity.FrontListProductVo{
			Id:       ps[k].Id,
			Name:     ps[k].Name,
			Intro:    ps[k].Intro,
			CoverUrl: ps[k].CoverUrl,
		}
	}
	rsp = result.NewSuccess(
		&entity.FrontListProductRsp{
			Page: &entity.Page{
				PageSize:  req.Page.PageSize,
				PageIndex: req.Page.PageIndex,
				Total:     total,
			},
			Ps: vos,
		})
	return
}

//获取产品列表
func (fP *FrontProduct) ListCategory(ctx *fasthttp.RequestCtx) (rsp *result.Result) {
	categories, err := service.FrontProdSrv.ListCategory()
	if err != nil {
		rsp = result.DatabaseError
		log.ErrorF("FrontProduct.ListCategory->[%v]", err)
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
		log.ErrorF("FrontProduct.GetProduct->[%v]", err)
		return
	}
	if !ok {
		rsp = result.NewError(result.RequestParamEc, "产品不存在")
		log.WarnF("FrontProduct.GetProduct->[产品不存在]->[%v]", req)
		return
	}
	i2sMap, err := service.FrontProdSrv.GetHotTop6Product()
	if err == nil {
		p.HeatSort = i2sMap[p.Id]
	}
	rsp = result.NewSuccess(&entity.FrontGetProductRsp{Product: p})
	return
}
