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

type SysSetting struct {
}

func init() {
	ss := &SysSetting{}
	router.AddRoute(
		"/admin/sys/setting/new",
		http.POST,
		ss.NewSetting,
		handler.CheckAdmin,
	)

	router.AddRoute(
		"/admin/sys/setting/modify",
		http.POST,
		ss.ModifySetting,
		handler.CheckAdmin,
	)

	router.AddRoute(
		"/admin/sys/setting",
		http.GET,
		ss.GetSetting,
		handler.CheckAdmin,
	)

}

//网站设置
func (ss *SysSetting) NewSetting(ctx *fasthttp.RequestCtx, req *entity.SystemSettingReq) (rsp *result.Result) {
	isExist, ok, err := service.SysSrv.InsertSystemSetting(req)
	if err != nil {
		rsp = result.DatabaseError
		return
	}
	if isExist {
		rsp = result.NewError(result.ParamEc, "网站设置不可以重复插入")
		return
	}
	if ok {
		rsp = result.NewSuccess("")
		return
	}
	rsp = result.NewError(result.UnKnowEc, "")
	return
}

//修改网站设置
func (ss *SysSetting) ModifySetting(ctx *fasthttp.RequestCtx, req *entity.SystemSettingReq) (rsp *result.Result) {
	ok, err := service.SysSrv.UpdateSystemSetting(req)
	if err != nil {
		return
	}
	if ok {
		rsp = result.NewSuccess("")
		return
	}
	rsp = result.NewError(result.UnKnowEc, "")
	return
}

//获取网站设置
func (ss *SysSetting) GetSetting(ctx *fasthttp.RequestCtx) (rsp *result.Result) {
	ok, webSetting, err := service.SysSrv.SelectSystemSetting()
	if err != nil {
		return
	}
	if ok {
		rsp = result.NewSuccess(&entity.GetSystemSettingRsp{Setting:webSetting})
		return
	}
	rsp = result.NewError(result.UnKnowEc, "")
	return
}