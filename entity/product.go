package entity

import (
	"flower/entity/gen"
)

type ListProductCategoryReq struct {
	ParentId int   `json:"parent_id"`
	Id       int   `json:"id"`
	Page     *Page `json:"page"`
}

type ListProductCategoryResp struct {
	Page *Page                  `json:"page"`
	Pcs  []*gen.ProductCategory `json:"pcs"`
}

type NewProdCategoryReq struct {
	Name     string `json:"name" validate:"required"`
	ParentId int    `json:"parent_id"`
	Sort     int    `json:"sort" validate:"required"`
	Desc     string `json:"desc"`
}

type ProdCategoryVo struct {
	Id       int               `json:"id"`
	Name     string            `json:"name"`
	Desc     string            `json:"desc"`
	States   int               `json:"states"`
	Level    int               `json:"level"`
	Sort     int               `json:"sort"`
	ParentId int               `json:"-"`
	Sub      []*ProdCategoryVo `json:"sub"`
}

type ProdCategoryTree struct {
	Tree []*ProdCategoryVo `json:"tree"`
}

type ProdCategoryStateReq struct {
	Id int `json:"id"`
}

type ModifyCategoryReq struct {
	Id     int    `json:"id" validate:"required"`
	Name   string `json:"name" validate:"required"`
	Desc   string `json:"desc"`
	States int    `json:"states"`
	Sort   int    `json:"sort"`
}

type DeleteProdCategoryReq struct {
	Id int `json:"id"`
}

// --------------商品--------------------

type NewProductReq struct {
	Name    string `json:"name" validate:"required"`
	Intro   string `json:"intro" validate:"required"`
	Summary string `json:"summary" validate:"required"`
	//States        int    `json:"states"`
	IndexShow     int    `json:"index_show" validate:"required"`
	DetailsPicUrl string `json:"details_pic_url" validate:"required"`
	CoverUrl      string `json:"cover_url" validate:"required"`
	Price         string `json:"price"validate:"required"`
	//Heat          int     `json:"heat"`
	CategoryId int   `json:"category_id" validate:"required"`
	AuthorId   int64 `json:"-"`
}

type NewProductRsp struct {
	ProductId int64 `json:"product_id"`
}

type ListProductReq struct {
	Name       string `json:"name"`
	States     int    `json:"states"`
	CategoryId int    `json:"category_id"`
	Page       *Page  `json:"page" validate:"required"`
}

type ListProductRsp struct {
	Page *Page          `json:"page"`
	Ps   []*gen.Product `json:"ps"`
}

type ChangeProductStateReq struct {
	Id int64 `json:"id" validate:"required"`
}

type ChangeProductIndexShowReq struct {
	Id int64 `json:"id" validate:"required"`
}

type ModifyProductReq struct {
	Id            int64  `json:"id" validate:"required"`
	Name          string `json:"name"`
	Intro         string `json:"intro"`
	Summary       string `json:"summary"`
	DetailsPicUrl string `json:"details_pic_url"`
	CoverUrl      string `json:"cover_url"`
	Price         string `json:"price"`
	CategoryId    int    `json:"category_id"`
	AuthorId      int64  `json:"-"`
}

type DeleteProductByIdReq struct {
	Id int64 `json:"id" validate:"required"`
}
