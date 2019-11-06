package controller

import (
	"flower/entity"
	"flower/http"
	"flower/result"
	"flower/router"
	"flower/service"
	"github.com/valyala/fasthttp"
)

type ProdCategory struct {
}

func init() {
	pc := &ProdCategory{}

	router.AddRoute(
		"/admin/product/category/list",
		http.POST,
		pc.ListProdCategory,
		//handler.CheckAdmin,
	)

}

// 获取客户列表
func (pc *ProdCategory) ListProdCategory(ctx *fasthttp.RequestCtx, req *entity.ProductCategoryReq) (resp *result.Result) {
	pcs, total, err := service.ProdSrv.ListProductCategory(req)
	if err != nil {
		//TODO
		resp = result.DatabaseError
		return
	}
	resp = result.NewSuccess(
		&entity.ProductCategoryResp{
			Page: &entity.Page{
				PageSize:  req.Page.PageSize,
				PageIndex: req.Page.PageIndex,
				Total:     total,
			},
			Pcs: pcs,
		})
	return
}
