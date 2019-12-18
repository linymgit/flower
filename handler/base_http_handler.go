package handler

import (
	"encoding/json"
	"flower/http"
	"flower/log"
	error2 "flower/result"
	"github.com/valyala/fasthttp"
	"reflect"
	"strings"
)

func BaseHttpHandler(handler interface{}) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		defer func() {
			if err := recover(); err != nil {
				log.ErrorF("BaseHttpHandler recover error[%v]", err)
				//e := err.(result)
				print(err)
				bytes, err := json.Marshal(error2.NewError(error2.ServerEc, "server internal result"))
				if err != nil {
					log.ErrorF("BaseHttpHandler->json.Marshal(error2.NewError(error2.ServerEc, \"server internal result\"))->%v", err)
				}
				ctx.Response.SetBodyRaw(bytes)
			}
		}()
		value := ctx.UserValue(http.CheckResultFastHttpKey)
		if value != nil {
			bytes, e := json.Marshal(value)
			if e != nil {
				log.ErrorF("BaseHttpHandler -> json.Marshal(value) -> [%v]", e)
			}
			ctx.Response.SetBodyRaw(bytes)
			return
		}
		ctx.Response.Header.SetContentType("application/json")
		fv := reflect.ValueOf(handler)
		ft := fv.Type()

		var params []reflect.Value
		numIn := ft.NumIn()
		// 请求参数处理
		if numIn > 1 {

			contentType := string(ctx.Request.Header.ContentType())
			if strings.Contains(contentType, http.Multipart) {
				fd, e := ctx.Request.MultipartForm()
				if e != nil {
					// TODO
				} else {
					fdMap := make(map[string]interface{}, len(fd.Value))
					for k := range fd.Value {
						fdMap[k] = fd.Value[k]
					}
					if rValue, ok := FillFieldValue(fdMap, ft.In(1)); ok {
						params = []reflect.Value{reflect.ValueOf(ctx), rValue}
					}
				}

			}
			if strings.EqualFold(contentType, http.Json) {
				tt := ft.In(1).Elem()
				vv := reflect.New(tt)
				err := json.Unmarshal(ctx.Request.Body(), vv.Interface())
				if err != nil {
					log.ErrorF("BaseHttpHandler -> json.Unmarshal(ctx.Request.Body(), vv.Interface())  ->[%v]", err)
				}
				params = []reflect.Value{reflect.ValueOf(ctx), vv}
			}
			if strings.EqualFold(contentType, http.Normal) || strings.EqualFold(contentType, "") {
				queryArgs := ctx.Request.URI().QueryArgs()
				if rValue, ok := FillFieldValueByQueryArgs(queryArgs, ft.In(1)); ok {
					params = []reflect.Value{reflect.ValueOf(ctx), rValue}
				}
			}

			if strings.EqualFold(contentType, http.NormalWithCharset) {
				//body := ctx.Request.Body()
				//values, e := url.ParseQuery(string(body))
				//if e != nil {
				//}
				//for k := range values {
				//	println(k)
				//	println(values[k])
				//}
				//fmt.Print("%#v", values)
				//TODO
				bytes, err := json.Marshal(error2.NewError(error2.ParamEc, "目前还不支持这种类型的请求，联系18316471919解决！！！"))
				if err != nil {
					// TODO
				}
				ctx.Response.SetBodyRaw(bytes)
				return
			}

			// 校验参数
			err := http.ValidateReq(params[1].Interface())
			if err != nil {
				bytes, err := json.Marshal(error2.NewError(error2.ParamEc, err.Error()))
				if err != nil {
					// TODO
				}
				ctx.Response.SetBodyRaw(bytes)
				return
			}

		} else {
			params = []reflect.Value{reflect.ValueOf(ctx)}
		}

		// 执行业务
		rspResult := fv.Call(params)

		for i := range rspResult {
			if i == 0 {
				bytes, e := json.Marshal(rspResult[i].Interface())
				if e != nil {
					log.ErrorF("BaseHttpHandler-> json.Marshal(rspResult[i].Interface())  ->[%v]", e)
					return
				}
				ctx.Response.SetBodyRaw(bytes)
			}
		}

	})
}
