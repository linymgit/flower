package controller

import (
	"flower/entity"
	"flower/handler"
	"flower/http"
	"flower/result"
	"flower/router"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"github.com/valyala/fasthttp"
)
type Resource struct {
}

func init() {
	r := &Resource{}

	router.AddRoute(
		"/admin/resource/token",
		http.GET,
		r.GetQNUploadToken,
		handler.CheckAdmin,
	)

}

func (r *Resource) GetQNUploadToken(ctx *fasthttp.RequestCtx)  (resp *result.Result) {
	upToken := getUploadToken()
	resp = result.NewSuccess(&entity.ResourceRsp{Token:upToken})
	return
}

//bucket accesskey secretkey 先写死之后再改为配置
func getUploadToken()(upToken string){
	bucket:="picture"
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac("aYpbxU9ziIThJUlbrf5M0-w61ouCi38yFiCcEhzH", "SkI35ZbYlq87xUvQyzDCpofj4HjRzCJbzlBb4M33")
	upToken = putPolicy.UploadToken(mac)
	return
}