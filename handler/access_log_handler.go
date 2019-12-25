package handler

import (
	"flower/log"
	"fmt"
	"github.com/valyala/fasthttp"
	"time"
)

func AccessLog(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		now := time.Now()
		next(ctx)
		log.AccessF(fmt.Sprintf("uri: %s>ip: %s>use: %d ms", string(ctx.RequestURI()),clientIP(ctx), time.Since(now).Milliseconds()))
	})
}