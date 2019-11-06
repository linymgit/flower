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

type Crm struct {
}

func init() {
	crm := &Crm{}

	router.AddRoute(
		"/admin/crm/list",
		http.POST,
		crm.ListCrm,
		handler.CheckAdmin,
	)

	router.AddRoute(
		"/admin/crm/delete",
		http.POST,
		crm.DeleteCrmById,
		handler.CheckAdmin,
	)

}

// 获取客户列表
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

//根据id删除客户信息 逻辑上删除物理并无删除
func (c *Crm) DeleteCrmById(ctx *fasthttp.RequestCtx, req *entity.CrmDeleteReq) (resp *result.Result) {
	ok, err := service.CrmSrv.DeleteCrmById(req.Id)
	if err != nil {
		resp = result.DatabaseError
		return
	}
	resp = result.NewSuccess(ok)
	return
}
