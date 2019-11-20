package controller

import (
	"flower/entity"
	"flower/http"
	"flower/result"
	"flower/router"
	"flower/service"
	"github.com/valyala/fasthttp"
)

type FrontCrm struct {
}

func init() {
	frontCrm := &FrontCrm{}

	router.AddRoute(
		"/crm/add",
		http.POST_AND_OPTIONS,
		frontCrm.InsertCrm,
	)

}

func (fC *FrontCrm) InsertCrm(ctx *fasthttp.RequestCtx, req *entity.AddCrmReq) (rsp *result.Result) {
	ok, err := service.CrmSrv.InsertCrm(req)
	if err != nil {
		rsp = result.DatabaseError
		return
	}
	if !ok {
		rsp = result.NewError(result.RequestParamEc, "插入失败检查参数")
		return
	}
	rsp = result.NewSuccess("提交成功")
	return
}
