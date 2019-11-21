package router

import (
	"flower/handler"
	"flower/http"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

var addRouterFunc []func()

var router = fasthttprouter.New()

var routes []string

var optionsRoutes = map[string]struct{}{}

func Do() {
	for _, f := range addRouterFunc {
		f()
	}
}

func AddRoute(path string, mothodType http.MothodType, handlerFunc interface{}, otherPreHandler ...func(n fasthttp.RequestHandler) fasthttp.RequestHandler) {
	addRoute(func() {
		// 基本请求的封装handler
		httpHandler := handler.BaseHttpHandler(handlerFunc)

		// 其他预先处理的handler
		for _, oPh := range otherPreHandler {
			httpHandler = oPh(httpHandler)
		}
		// 跨域处理
		httpHandler = handler.CROS(httpHandler)

		// 路由注册
		if mothodType == http.GET {
			router.GET(path, httpHandler)
		}
		if mothodType == http.POST {
			router.POST(path, httpHandler)
		}
		if mothodType == http.OPTIONS {
			router.OPTIONS(path, httpHandler)
		}
		if mothodType == http.DELETE {
			router.DELETE(path, httpHandler)
		}
		if mothodType == http.POST_AND_OPTIONS {
			router.POST(path, httpHandler)
			router.OPTIONS(path, httpHandler)
		}
		if mothodType == http.GET_AND_OPTIONS{
			router.GET(path, httpHandler)
			router.OPTIONS(path, httpHandler)
		}
	})
}

func Handler(ctx *fasthttp.RequestCtx) {
	router.Handler(ctx)
}

func addRoute(f func()) {
	addRouterFunc = append(addRouterFunc, f)
}

type CommonHandler func()
