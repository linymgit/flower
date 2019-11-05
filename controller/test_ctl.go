package controller

import (
	"flower/http"
	"flower/router"
	"github.com/valyala/fasthttp"
)

type HH struct {
}

func init() {
	hh := &HH{}
	router.AddRoute("/test", http.POST, hh.Say)
	router.AddRoute("/test", http.GET, hh.Say)
}

type Hello struct {
	Name string `json:"name" validate:"required"`
}

func (h *HH) Say(ctx *fasthttp.RequestCtx, i *Hello) (j *Hello) {
	j = new(Hello)
	j.Name = "lym"
	return
}
