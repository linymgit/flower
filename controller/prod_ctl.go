package controller

import (
	"flower/entity"
	"flower/handler"
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
		http.GET,
		pc.ListProdCategory,
		//handler.CheckAdmin,
	)

	router.AddRoute(
		"/admin/product/category/new",
		http.POST,
		pc.NewProdCategory,
		handler.CheckAdmin,
	)

	router.AddRoute(
		"/admin/product/category/tree",
		http.POST,
		pc.ProdCategoryTree,
		handler.CheckAdmin,
	)

	router.AddRoute(
		"/admin/product/category/state",
		http.POST,
		pc.ChangeProcategoryState,
		handler.CheckAdmin,
	)

	//新增商品
	router.AddRoute(
		"/admin/product/new",
		http.POST,
		pc.NewProduct,
		handler.CheckAdmin,
	)

}

//获取产品分类列表
func (pc *ProdCategory) ListProdCategory(ctx *fasthttp.RequestCtx, req *entity.ListProductCategoryReq) (resp *result.Result) {
	pcs, total, err := service.ProdSrv.ListProductCategory(req)
	if err != nil {
		resp = result.DatabaseError
		return
	}
	var page *entity.Page
	if req.Page != nil {
		page = &entity.Page{
			PageSize:  req.Page.PageSize,
			PageIndex: req.Page.PageIndex,
			Total:     total,
		}
	}
	resp = result.NewSuccess(
		&entity.ListProductCategoryResp{
			Page: page,
			Pcs: pcs,
		})
	return
}

//创建产品分类
func (pc *ProdCategory) NewProdCategory(ctx *fasthttp.RequestCtx, req *entity.NewProdCategoryReq) (resp *result.Result) {
	isExistName, isExistParent, ok, err := service.ProdSrv.NewProductCategory(req)
	if err != nil {
		resp = result.DatabaseError
		return
	}
	if isExistName {
		resp = result.NewError(result.RequestParamEc, "分类名称已存在")
		return
	}
	if !isExistParent {
		resp = result.NewError(result.RequestParamEc, "上级分类不存在")
		return
	}
	if ok {
		resp = result.NewSuccess("创建成功")
	} else {
		resp = result.NewSuccess("创建失败")
	}
	return
}

func (pc *ProdCategory) ProdCategoryTree(ctx *fasthttp.RequestCtx) (resp *result.Result) {
	tree, err := service.ProdSrv.GetProductCategoryTree()
	if err != nil {
		result.NewError(result.UnKnowEc, "未知错误")
		return
	}
	resp = result.NewSuccess(tree)
	return
}

func (pc *ProdCategory) ChangeProcategoryState(ctx *fasthttp.RequestCtx, req *entity.ProdCategoryStateReq) (resp *result.Result) {
	ok, err := service.ProdSrv.ChangeProcategoryState(req.Id)
	if err != nil {
		resp = result.DatabaseError
		return
	}
	if !ok {
		resp = result.NewError(result.ParamEc, "分类id 不存在")
		return
	}
	resp = result.NewSuccess(ok)
	return
}

func (pc *ProdCategory) NewProduct(ctx *fasthttp.RequestCtx, req *entity.NewProductReq) (rsp *result.Result) {
	id, ok := http.GetJwtId(ctx)
	if !ok {
		rsp = result.NewError(result.UnKnowEc, "未登录")
	}
	req.AuthorId = id
	productId, err := service.ProdSrv.NewProductResutId(req)
	if err != nil {
		rsp = result.DatabaseError
		return
	}
	rsp = result.NewSuccess(&entity.NewProductRsp{ProductId: productId})
	return
}
