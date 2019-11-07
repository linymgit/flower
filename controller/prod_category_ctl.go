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

	router.AddRoute(
		"/admin/product/category/new",
		http.POST,
		pc.NewCategory,
		//handler.CheckAdmin,
	)

}

//获取产品分类列表
func (pc *ProdCategory) ListProdCategory(ctx *fasthttp.RequestCtx, req *entity.ListProductCategoryReq) (resp *result.Result) {
	pcs, total, err := service.ProdSrv.ListProductCategory(req)
	if err != nil {
		resp = result.DatabaseError
		return
	}
	resp = result.NewSuccess(
		&entity.ListProductCategoryResp{
			Page: &entity.Page{
				PageSize:  req.Page.PageSize,
				PageIndex: req.Page.PageIndex,
				Total:     total,
			},
			Pcs: pcs,
		})
	return
}

//创建产品分类
func (pc *ProdCategory) NewCategory(ctx *fasthttp.RequestCtx, req *entity.NewProdCategoryReq) (resp *result.Result) {
	isExistName, ok, err := service.ProdSrv.NewProductCategory(req)
	if err != nil {
		resp = result.DatabaseError
		return
	}
	if isExistName {
		resp = result.NewError(result.RequestParamEc, "分类名称已存在")
		return
	}
	if ok {
		resp = result.NewSuccess("创建成功")
	}else{
		resp = result.NewSuccess("创建失败")
	}
	return
}
