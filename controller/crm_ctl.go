package controller

import (
	"flower/entity"
	"flower/http"
	"flower/result"
	"flower/router"
	"flower/service"
	"github.com/valyala/fasthttp"
)

type Crm struct {
}

func init() {
	crm := &Crm{}

	router.AddRoute(
		"/admin/crm/list",
		http.POST,
		crm.ListCrm,
	)
}

func (c *Crm) ListCrm(ctx *fasthttp.RequestCtx, req *entity.CrmListReq) (resp *result.Result) {
	if req.BeginTime > req.EndTime {
		resp = result.NewError(result.RequestParamEc, "开始时间大于结束时间")
		return
	}
	crms, total, err := service.CrmSrv.ListCrm(req)
	if err != nil {
		//TODO
		resp = result.DatabaseError
		return
	}
	resp = result.NewSuccess(
		&entity.CrmListResp{
			Page: &entity.Page{
				PageSize:  req.Page.PageSize,
				PageIndex: req.Page.PageIndex,
				Total:     total,
			},
			Crms: crms,
		})
	return
}
