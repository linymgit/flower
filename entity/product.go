package entity

import (
	"flower/entity/gen"
)

type ListProductCategoryReq struct {
	ParentId int   `json:"parent_id"`
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

type NewProductReq struct {
	Name          string `json:"name" validate:"required"`
	Intro         string `json:"intro" validate:"required"`
	Summary       string `json:"summary" validate:"required"`
	//States        int    `json:"states"`
	IndexShow     int    `json:"index_show" validate:"required"`
	DetailsPicUrl string `json:"details_pic_url" validate:"required"`
	CoverUrl      string `json:"cover_url" validate:"required"`
	Price         string `json:"price"validate:"required"`
	//Heat          int     `json:"heat"`
	CategoryId int `json:"category_id" validate:"required"`
	AuthorId   int64 `json:"-"`
}

type NewProductRsp struct {
	ProductId int64 `json:"product_id"`
}

