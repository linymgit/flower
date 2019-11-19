package controller

import (
	"flower/entity"
	"flower/http"
	"flower/result"
	"flower/router"
	"flower/service"
	"github.com/valyala/fasthttp"
)

type Index struct {
}

func init() {
	index := &Index{}

	router.AddRoute(
		"/index/ad/list",
		http.POST_AND_OPTIONS,
		index.ListIndexAd,
	)

	router.AddRoute(
		"/index/product/list",
		http.POST_AND_OPTIONS,
		index.ListIndexProduct,
	)
}

func (i *Index) ListIndexAd(ctx *fasthttp.RequestCtx, req *entity.IndexReq) (rsp *result.Result) {
	ads, total, err := service.IndexSrv.ListIndexAd(req)
	if err != nil {
		rsp = result.DatabaseError
		return
	}
	rsp = result.NewSuccess(
		&entity.IndexAdRsp{
			Page: &entity.Page{
				PageSize:  req.Page.PageSize,
				PageIndex: req.Page.PageIndex,
				Total:     total,
			},
			Ad: ads,
		})
	return
}

func (i *Index) ListIndexProduct(ctx *fasthttp.RequestCtx, req *entity.IndexReq) (rsp *result.Result) {
	ps, total, err := service.IndexSrv.ListIndexProduct(req)
	if err != nil {
		rsp = result.DatabaseError
		return
	}
	ip := make([]*entity.IndexProduct, total)
	for k := range ps {
		ip[k] = &entity.IndexProduct{
			Id:       ps[k].Id,
			Name:     ps[k].Name,
			Summary:  ps[k].Summary,
			CoverUrl: ps[k].CoverUrl,
		}
	}
	rsp = result.NewSuccess(
		&entity.IndexProductRsp{
			Page: &entity.Page{
				PageSize:  req.Page.PageSize,
				PageIndex: req.Page.PageIndex,
				Total:     total,
			},
			Ps: ip,
		})
	return
}
