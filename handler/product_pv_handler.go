package handler

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"strings"
)

func Pv4Product(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		//fmt.Println("pv log start" + clientIP(ctx))
		next(ctx)
		body := ctx.Request.Body()
		//fmt.Println(string(body))
		log := &PvLog{}
		err := json.Unmarshal(body, log)
		if err != nil {
			return
		}
		//加pv
		fmt.Println(log.Id)
	})
}

//获取真实的IP  1.1.1.1, 2.2.2.2, 3.3.3.3
func clientIP(ctx *fasthttp.RequestCtx) string {
	clientIP := string(ctx.Request.Header.Peek("X-Forwarded-For"))
	if index := strings.IndexByte(clientIP, ','); index >= 0 {
		clientIP = clientIP[0:index]
		//获取最开始的一个 即 1.1.1.1
	}
	clientIP = strings.TrimSpace(clientIP)
	if len(clientIP) > 0 {
		return clientIP
	}
	clientIP = strings.TrimSpace(string(ctx.Request.Header.Peek("X-Real-Ip")))
	if len(clientIP) > 0 {
		return clientIP
	}
	return ctx.RemoteIP().String()
}

type PvLog struct {
	Id int64 `json:"id"`
}
