package handler

import (
	"flower/entity/role"
	"flower/http"
	"flower/result"
	"github.com/valyala/fasthttp"
)

// CheckAdmin 检验是否为管理员
func CheckAdmin(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	checkToken := CheckToken(func(ctx *fasthttp.RequestCtx) {
		value := ctx.UserValue(http.CheckResultFastHttpKey)
		if value != nil {
			next(ctx)
			return
		}
		roleId, ok := http.GetRoleId(ctx)
		if !ok || roleId&role.Admin <= 0 {
			ctx.SetUserValue(http.CheckResultFastHttpKey, result.ForbiddenError)
		}
		next(ctx)
	})
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		checkToken(ctx)
	})
}
