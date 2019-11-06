package entity

import "flower/entity/gen"

type ProductCategoryReq struct {
	Page *Page `json:"page" validate:"required"`
}
type ProductCategoryResp struct {
	Page *Page                  `json:"page"`
	Pcs  []*gen.ProductCategory `json:"crms"`
}
