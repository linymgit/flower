package entity

import "flower/entity/gen"

type FrontListProductReq struct {
	IsIndexShow bool   `json:"is_index_show"`
	IsNew       bool   `json:"is_new"`
	IsHot       bool   `json:"is_hot"`
	Search      string `json:"search"`
	CategoryId  int    `json:"category_id"`
	Page        *Page  `json:"page"`
}

type FrontListProductVo struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Intro    string `json:"intro"`
	CoverUrl string `json:"cover_url"`
}

type FrontListProductRsp struct {
	Page *Page                 `json:"page"`
	Ps   []*FrontListProductVo `json:"ps"`
}

type FrontCategory struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type FrontListCategoryRsp struct {
	Categories []*FrontCategory `json:"categories"`
}

type FrontGetProductReq struct {
	Id int64 `json:"id"`
}

type FrontGetProductRsp struct {
	Product *gen.Product `json:"product"`
}
