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

type BusinessPartners struct {
}

func init() {
	bp := &BusinessPartners{}
	//后台
	router.AddRoute(
		"/admin/business/partners/list",
		http.POST,
		bp.AdminListBusinessPartners,
		handler.CheckAdmin,
	)
	router.AddRoute(
		"/admin/business/partners/delete",
		http.POST,
		bp.DeleteBusinessPartnersById,
		handler.CheckAdmin,
	)
	router.AddRoute(
		"/admin/business/partners/add",
		http.POST,
		bp.AddBusinessPartners,
		handler.CheckAdmin,
	)
	router.AddRoute(
		"/admin/business/partners/modify",
		http.POST,
		bp.ModifyBusinessPartners,
		handler.CheckAdmin,
	)

	//前台
	router.AddRoute(
		"/index/business/partners/list",
		http.POST_AND_OPTIONS,
		bp.FrontListBusinessPartners,
	)

}

func (bp *BusinessPartners) AdminListBusinessPartners(ctx *fasthttp.RequestCtx, req *entity.BusinessPartnersPageReq) (rsp *result.Result) {
	bps, total, err := service.BusinessPartnersSrv.AdminListBusinessPartners(req)
	if err != nil {
		rsp = result.DatabaseError
		log.ErrorF("AdminListBusinessPartners->[%v]", err)
		return
	}
	rsp = result.NewSuccess(
		&entity.AdminBusinessPartnersRsp{
			Page: &entity.Page{
				PageSize:  req.Page.PageSize,
				PageIndex: req.Page.PageIndex,
				Total:     total,
			},
			Bps: bps,
		})
	return
}
func (bp *BusinessPartners) DeleteBusinessPartnersById(ctx *fasthttp.RequestCtx, req *entity.BusinessPartnersIdReq) (rsp *result.Result) {
	ok, err := service.BusinessPartnersSrv.DeleteBusinessPartnersById(req.Id)
	if err != nil {
		rsp = result.DatabaseError
		log.ErrorF("DeleteBusinessPartnersById->[%v]", err)
		return
	}
	if !ok {
		rsp = result.NewError(result.RequestParamEc, "不存在这个id")
		log.WarnF("DeleteBusinessPartnersById->[不存在这个id]->[%v]", req)
		return
	}
	rsp = result.NewSuccess("删除成功")
	return
}
func (bp *BusinessPartners) AddBusinessPartners(ctx *fasthttp.RequestCtx, req *entity.BusinessPartnersReq) (rsp *result.Result) {
	ok, err := service.BusinessPartnersSrv.AddBusinessPartners(req)
	if err != nil {
		rsp = result.DatabaseError
		log.ErrorF("AddBusinessPartners->[%v]", err)
		return
	}
	if !ok {
		rsp = result.NewError(result.RequestParamEc, "创建失败")
		log.WarnF("AddBusinessPartners->[创建失败]->[%v]", req)
		return
	}
	rsp = result.NewSuccess("创建成功")
	return
}
func (bp *BusinessPartners) ModifyBusinessPartners(ctx *fasthttp.RequestCtx, req *entity.BusinessPartnersReq) (rsp *result.Result) {
	ok, err := service.BusinessPartnersSrv.ModifyBusinessPartners(req)
	if err != nil {
		rsp = result.DatabaseError
		log.ErrorF("ModifyBusinessPartners->[%v]", err)
		return
	}
	if !ok {
		rsp = result.NewError(result.RequestParamEc, "修改失败")
		log.WarnF("ModifyBusinessPartners->[修改失败]->[%v]", req)
		return
	}
	rsp = result.NewSuccess("修改成功")
	return
}
func (bp *BusinessPartners) FrontListBusinessPartners(ctx *fasthttp.RequestCtx, req *entity.BusinessPartnersPageReq) (rsp *result.Result) {
	bps, total, err := service.BusinessPartnersSrv.FrontListBusinessPartners(req)
	if err != nil {
		rsp = result.DatabaseError
		log.ErrorF("FrontListBusinessPartners->[%v]", err)
		return
	}
	iBps := make([]*entity.IndexBusinessPartners, total)
	for k := range iBps {
		iBps[k] = &entity.IndexBusinessPartners{
			Id:   bps[k].Id,
			Logo: bps[k].Logo,
		}
	}
	rsp = result.NewSuccess(
		&entity.IndexBusinessPartnersRsp{
			Page: &entity.Page{
				PageSize:  req.Page.PageSize,
				PageIndex: req.Page.PageIndex,
				Total:     total,
			},
			Bps: iBps,
		})
	return
}
