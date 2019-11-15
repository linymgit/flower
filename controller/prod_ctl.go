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
		http.POST,
		pc.ListProdCategory,
		handler.CheckAdmin,
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

	router.AddRoute(
		"/admin/product/category/modify",
		http.POST,
		pc.ModifyProcategory,
		handler.CheckAdmin,
	)

	router.AddRoute(
		"/admin/product/category/delete",
		http.POST,
		pc.DeleteProcategory,
		handler.CheckAdmin,
	)

	// --------------商品--------------------

	//新增商品
	router.AddRoute(
		"/admin/product/new",
		http.POST,
		pc.NewProduct,
		handler.CheckAdmin,
	)

	//获取商品
	router.AddRoute(
		"/admin/product/list",
		http.POST,
		pc.ListProduct,
		handler.CheckAdmin,
	)

	//上下架商品
	router.AddRoute(
		"/admin/product/state/change",
		http.POST,
		pc.ChangeProductState,
		handler.CheckAdmin,
	)

	//设置是否首页推荐
	router.AddRoute(
		"/admin/product/indexshow/change",
		http.POST,
		pc.ChangeProductIndexShow,
		handler.CheckAdmin,
	)

	//修改商品
	router.AddRoute(
		"/admin/product/modify",
		http.POST,
		pc.ModifyProduct,
		handler.CheckAdmin,
	)

	//删除商品
	router.AddRoute(
		"/admin/product/delete",
		http.POST,
		pc.DeleteProductById,
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
			Pcs:  pcs,
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
		resp = result.DatabaseError
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

func (pc *ProdCategory) ModifyProcategory(ctx *fasthttp.RequestCtx, req *entity.ModifyCategoryReq) (rsp *result.Result) {
	affected, err := service.ProdSrv.ModifyProcategory(req)
	if err != nil {
		rsp = result.DatabaseError
		return
	}
	if affected != 1 {
		rsp = result.NewError(result.RequestParamEc, "无修改的数据")
		return
	}
	rsp = result.NewSuccess("修改成功")
	return
}

func (pc *ProdCategory) DeleteProcategory(ctx *fasthttp.RequestCtx, req *entity.DeleteProdCategoryReq) (rsp *result.Result) {
	isParent, affected, err := service.ProdSrv.DeleteProcategoryById(req.Id)
	if err != nil {
		rsp = result.DatabaseError
		return
	}
	if isParent {
		rsp = result.NewError(result.ParamEc, "是上级分类,要删除这个分类要先删除他的所有下级分类")
		return
	}
	if affected != 1 {
		rsp = result.NewError(result.RequestParamEc, "id不存在")
		return
	}
	rsp = result.NewSuccess("删除成功")
	return
}

// --------------商品--------------------

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

func (pc *ProdCategory) ListProduct(ctx *fasthttp.RequestCtx, req *entity.ListProductReq) (rsp *result.Result) {
	ps, total, err := service.ProdSrv.ListProduct(req)
	if err != nil {
		rsp = result.DatabaseError
		return
	}
	id2nameMap, err := service.ProdSrv.CategoryId2Name()
	if err != nil {
		rsp = result.DatabaseError
		return
	}
	for k := range ps {
		ps[k].CategoryName = id2nameMap[ps[k].CategoryId]
	}
	rsp = result.NewSuccess(
		&entity.ListProductRsp{
			Page: &entity.Page{
				PageSize:  req.Page.PageSize,
				PageIndex: req.Page.PageIndex,
				Total:     total,
			},
			Ps: ps,
		})
	return
}

func (pc *ProdCategory) ChangeProductState(ctx *fasthttp.RequestCtx, req *entity.ChangeProductStateReq) (rsp *result.Result) {
	ok, err := service.ProdSrv.ChangeProductState(req.Id)
	if err != nil {
		rsp = result.DatabaseError
		return
	}
	if !ok {
		rsp = result.NewError(result.RequestParamEc, "产品id不存在")
		return
	}
	rsp = result.NewSuccess("修改成功")
	return
}

func (pc *ProdCategory) ChangeProductIndexShow(ctx *fasthttp.RequestCtx, req *entity.ChangeProductIndexShowReq) (rsp *result.Result) {
	ok, err := service.ProdSrv.ChangeProductIndexShow(req.Id)
	if err != nil {
		rsp = result.DatabaseError
		return
	}
	if !ok {
		rsp = result.NewError(result.RequestParamEc, "产品id不存在")
		return
	}
	rsp = result.NewSuccess("修改成功")
	return
}

func (pc *ProdCategory) ModifyProduct(ctx *fasthttp.RequestCtx, req *entity.ModifyProductReq) (rsp *result.Result) {
	ok, err := service.ProdSrv.ModifyProduct(req)
	if err != nil {
		rsp = result.DatabaseError
		return
	}
	if !ok {
		rsp = result.NewError(result.RequestParamEc, "无修改的数据")
		return
	}
	rsp = result.NewSuccess("修改成功")
	return
}

func (pc *ProdCategory) DeleteProductById(ctx *fasthttp.RequestCtx, req *entity.DeleteProductByIdReq) (rsp *result.Result) {
	ok, err := service.ProdSrv.DeleteProductById(req.Id)
	if err != nil {
		rsp = result.DatabaseError
		return
	}
	if !ok {
		rsp = result.NewError(result.RequestParamEc, "产品id不存在")
		return
	}
	rsp = result.NewSuccess("删除成功")
	return
}
