package handler

import (
	"encoding/json"
	"flower/entity/gen"
	"flower/log"
	"flower/mysql"
	"flower/result"
	"github.com/valyala/fasthttp"
	"strings"
)

func Pv4Product(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		//fmt.Println("pv log start" + clientIP(ctx))
		next(ctx)
		body := ctx.Request.Body()
		//fmt.Println(string(body))
		pvLog := &PvLog{}
		err := json.Unmarshal(body, pvLog)
		if err != nil {
			log.ErrorF("Pv4Product->json.Unmarshal(body, pvLog)->[%v]", err)
			return
		}
		rsp := ctx.Response.Body()
		r := &result.Result{}
		err = json.Unmarshal(rsp, r)
		if err != nil {
			log.ErrorF("Pv4Product->json.Unmarshal(rsp, r)->[%v]", err)
			return
		}
		if r.Code != result.SUCCESS {
			log.InfoF("Pv4Product->产品详情接口无数据")
			return
		}
		//加pv
		_, err = mysql.Db.Exec("UPDATE `product` SET `heat`=`heat`+1 WHERE  `id`=?;", pvLog.Id)
		if err != nil {
			log.ErrorF("Pv4Product->添加产品pv数据库出错->[%v]", err)
		}
		//加uv
		affected, err := mysql.Db.Cols("p_id", "ip").InsertOne(&gen.ProductUv{
			PId: pvLog.Id,
			Ip:  clientIP(ctx),
		})
		if err != nil {
			log.ErrorF("Pv4Product->添加产品uv数据库出错->[%v]", err)
		}
		if affected != 1 {
			log.WarnF("Pv4Product->添加产品uv数据库影响行数不等于1")
		}
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
