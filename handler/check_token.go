package handler

import (
	"flower/http"
	"flower/jwt"
	"flower/result"
	"github.com/valyala/fasthttp"
)

// 检验是否登录
func CheckToken(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		peek := ctx.Request.Header.Peek(http.TokenKey)
		if peek == nil {
			ctx.SetUserValue(http.CheckResultFastHttpKey, result.UnauthorizedError)
		} else {
			if claims, ok, err := jwt.ParseJwt(string(peek)); !ok || err != nil {
				ctx.SetUserValue(http.CheckResultFastHttpKey, result.InvalidTokenError)
			} else {
				id := claims[http.JwtIdKey]
				roleId := claims[http.JwtRoleId]
				if id == nil || roleId == nil {
					ctx.SetUserValue(http.CheckResultFastHttpKey, result.ForbiddenError)
				} else {
					ctx.SetUserValue(http.JwtIdKey, id)
					ctx.SetUserValue(http.JwtRoleId, roleId)
				}
			}
		}
		next(ctx)
	})
}
