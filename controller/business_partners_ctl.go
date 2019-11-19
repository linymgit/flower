package controller

import (
	"flower/http"
	"flower/result"
	"flower/router"
	"github.com/valyala/fasthttp"
)

type BusinessPartners struct {
}

func init() {
	bp := &BusinessPartners{}
	//后台
	router.AddRoute(
		"/admin/business/partners/list",
		http.POST,
		bp.ListBusinessPartners,
	)
	router.AddRoute(
		"/admin/business/partners/delete",
		http.POST,
		bp.ListBusinessPartners,
	)
	router.AddRoute(
		"/admin/business/partners/add",
		http.POST,
		bp.ListBusinessPartners,
	)
	router.AddRoute(
		"/admin/business/partners/modify",
		http.POST,
		bp.ListBusinessPartners,
	)

	//前台
	router.AddRoute(
		"/business/partners/list",
		http.POST,
		bp.ListBusinessPartners,
	)

}

func (bp *BusinessPartners) ListBusinessPartners(ctx *fasthttp.RequestCtx) (resp *result.Result) {
	return
}