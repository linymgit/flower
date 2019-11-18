package http

import (
	"github.com/valyala/fasthttp"
)

type MothodType int32

const (
	// 请求类型
	Multipart         = "multipart/form-data;"
	Json              = "application/json"
	Normal            = "application/x-www-form-urlencoded"
	NormalWithCharset = "application/x-www-form-urlencoded; charset=UTF-8"
	// 请求header
	TokenKey = "f-token"
	// jwt key
	JwtIdKey  = "id"
	JwtRoleId = "rid"
	JwtExp    = "exp"
	// 检验权限 fasthttp context的key
	CheckResultFastHttpKey = "resp"
)

// 请求方法
const (
	GET MothodType = iota
	POST
	DELETE
	OPTIONS
	POST_AND_OPTIONS
)

func GetJwtId(ctx *fasthttp.RequestCtx) (id int64, ok bool) {
	peek := ctx.UserValue(JwtIdKey)
	if peek == nil {
		return
	}
	switch peek.(type) {
	case float64:
		id = int64(peek.(float64))
	case float32:
		id = int64(peek.(float32))
	case int:
		id = int64(peek.(int))
	case int64:
		id = peek.(int64)
	case int32:
		id = int64(peek.(int32))
	case uint:
		id = int64(peek.(uint))
	case uint64:
		id = int64(peek.(uint64))
	case uint32:
		id = int64(peek.(uint32))
	case uint16:
		id = int64(peek.(uint16))
	case uint8:
		id = int64(peek.(uint8))
	default:
		// TODO
		return
	}
	ok = true
	return
}
func GetRoleId(ctx *fasthttp.RequestCtx) (roleId int, ok bool) {
	peek := ctx.UserValue(JwtRoleId)
	if peek == nil {
		return
	}
	switch peek.(type) {
	case float64:
		roleId = int(peek.(float64))
	case float32:
		roleId = int(peek.(float32))
	case int:
		roleId = peek.(int)
	case int64:
		roleId = int(peek.(int64))
	case int32:
		roleId = int(peek.(int32))
	case uint:
		roleId = int(peek.(uint))
	case uint64:
		roleId = int(peek.(uint64))
	case uint32:
		roleId = int(peek.(uint32))
	case uint16:
		roleId = int(peek.(uint16))
	case uint8:
		roleId = int(peek.(uint8))
	default:
		// TODO
		return
	}
	ok = true
	return
}
