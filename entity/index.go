package entity

import "flower/entity/gen"

type IndexReq struct {
	Page *Page `json:"page" validate:"required"`
}

type IndexAdRsp struct {
	Page *Page     `json:"page"`
	Ad   []*gen.Ad `json:"ad"`
}

type IndexProduct struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Summary  string `json:"summary"`
	CoverUrl string `json:"cover_url"`
}

type IndexProductRsp struct {
	Page *Page           `json:"page"`
	Ps   []*IndexProduct `json:"ad"`
}

type IndexReqV2 struct {
	Page      *Page `json:"page" validate:"required"`
	PostionId int   `json:"postion_id"`
}
