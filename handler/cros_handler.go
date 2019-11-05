package handler

import (
	"github.com/valyala/fasthttp"
)

const (
	corsAllowHeaders     = "Content-Type, Content-Length, Accept-Encoding, x-token, Authorization, accept, origin, Cache-Control, x-req-id"
	corsAllowMethods     = "HEAD,GET,POST,PUT,DELETE,OPTIONS"
	corsAllowOrigin      = "*"
	corsAllowCredentials = "true"
)

func CROS(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {

		ctx.Response.Header.Set("Access-Control-Allow-Credentials", corsAllowCredentials)
		ctx.Response.Header.Set("Access-Control-Allow-Headers", corsAllowHeaders)
		ctx.Response.Header.Set("Access-Control-Allow-Methods", corsAllowMethods)
		ctx.Response.Header.Set("Access-Control-Allow-Origin", corsAllowOrigin)

		next(ctx)
	})
}
