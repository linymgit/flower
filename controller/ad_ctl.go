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

type Ad struct {
}

func init() {
	ad := &Ad{}

	router.AddRoute(
		"/admin/ad",
		http.POST,
		ad.NewGd,
		handler.CheckAdmin,
	)

	router.AddRoute(
		"/admin/ad",
		http.GET,
		ad.GetGds,
		handler.CheckAdmin,
	)

	router.AddRoute(
		"/admin/ad/state",
		http.POST,
		ad.ChangeAdState,
		handler.CheckAdmin,
	)

	router.AddRoute(
		"/admin/ad/delete",
		http.POST,
		ad.DeleteAdById,
		handler.CheckAdmin,
	)
}

func (ad *Ad) NewGd(ctx *fasthttp.RequestCtx, req *entity.NewAdReq) (rsp *result.Result) {
	adId, err := service.Adsrv.NewAd(req)
	if err != nil {
		rsp = result.DatabaseError
		return
	}
	rsp = result.NewSuccess(&entity.NewAdRsp{AdId:adId})
	return
}

func (ad *Ad) GetGds(ctx *fasthttp.RequestCtx, req *entity.GetAdsReq ) (rsp *result.Result){
	ads, total, err := service.Adsrv.GetAds(req)
	if err != nil {
		rsp = result.DatabaseError
		return
	}
	rsp = result.NewSuccess(
		&entity.GetAdsResp{
			Page: &entity.Page{
				PageSize:  req.Page.PageSize,
				PageIndex: req.Page.PageIndex,
				Total:     total,
			},
			Ads: ads,
		})
	return
}

func (ad *Ad)  ChangeAdState(ctx *fasthttp.RequestCtx, req *entity.ChangeAdStateReq ) (rsp *result.Result){
	ok, err := service.Adsrv.ChangeAdState(req.Id)
	if err != nil {
		rsp = result.DatabaseError
		return
	}
	rsp = result.NewSuccess(ok)
	return
}

func (ad *Ad)  DeleteAdById(ctx *fasthttp.RequestCtx, req *entity.ChangeAdStateReq ) (rsp *result.Result){
	ok, err := service.Adsrv.DeleteAd(req.Id)
	if err != nil {
		rsp = result.DatabaseError
		return
	}
	rsp = result.NewSuccess(ok)
	return
}
