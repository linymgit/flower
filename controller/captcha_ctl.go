package controller

import (
	"flower/captcha"
	"flower/entity"
	"flower/http"
	"flower/result"
	"flower/router"
	"github.com/valyala/fasthttp"
)

type Cap struct {
}

func init() {
	c := &Cap{}
	router.AddRoute("/captcha", http.GET, c.get)
}

func (c *Cap) get(ctx *fasthttp.RequestCtx) (resp *result.Result) {
	id, base64Png := captcha.GetCaptcha()
	resp = result.NewSuccess(&entity.Capache{
		Id:        id,
		Base64Png: base64Png,
	})
	return
}
