package handler

import (
	"encoding/json"
	"flower/result"
	"github.com/valyala/fasthttp"
)

var Options = CROS(func(ctx *fasthttp.RequestCtx) {
	success := result.NewSuccess("")
	bytes, e := json.Marshal(success)
	if e != nil {
		//todo
	}
	ctx.Response.SetBodyRaw(bytes)
})
