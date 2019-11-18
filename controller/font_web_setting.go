package controller

import (
	"flower/entity"
	"flower/http"
	"flower/result"
	"flower/router"
	"flower/service"
	"github.com/valyala/fasthttp"
)

type FrontWebSetting struct {
}

func init() {
	setting := &FrontWebSetting{}

	router.AddRoute(
		"/index/Setting/get",
		http.POST_AND_OPTIONS,
		setting.GetSetting,
	)

}

//获取网站设置
func (s *FrontWebSetting) GetSetting(ctx *fasthttp.RequestCtx) (rsp *result.Result) {
	ok, webSetting, err := service.SysSrv.SelectSystemSetting()
	if err != nil {
		return
	}
	if ok {
		rsp = result.NewSuccess(&entity.GetSystemSettingRsp{Setting: webSetting})
		return
	}
	rsp = result.NewError(result.UnKnowEc, "请检查网站设置")
	return
}
