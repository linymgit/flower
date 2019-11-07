package entity

import "flower/entity/gen"

type ListProductCategoryReq struct {
	Page *Page `json:"page" validate:"required"`
}

type ListProductCategoryResp struct {
	Page *Page                  `json:"page"`
	Pcs  []*gen.ProductCategory `json:"crms"`
}

type NewProdCategoryReq struct {
	Name string `json:"name" validate:"required"`
	ParentId int `json:"parent_id"`
	Sort  int `json:"sort" validate:"required"`
	Desc string `json:"desc"`
}

type ProdCategoryTree struct {

}