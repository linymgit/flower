package controller

import (
	"flower/entity"
	"flower/http"
	"flower/log"
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
		"/index/ad/list/v2",
		http.POST_AND_OPTIONS,
		index.ListIndexAdV2,
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
		log.ErrorF("ListIndexAd->[%v]", err)
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

func (i *Index) ListIndexAdV2(ctx *fasthttp.RequestCtx, req *entity.IndexReqV2) (rsp *result.Result) {
	ads, total, err := service.IndexSrv.ListIndexAdV2(req)
	if err != nil {
		rsp = result.DatabaseError
		log.ErrorF("ListIndexAdv2->[%v]", err)
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
		log.ErrorF("ListIndexProduct->[%v]", err)
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
