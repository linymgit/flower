package entity

import "flower/entity/gen"

type NewAdReq struct {
	Slogan    string `json:"slogan" validate:"required"`
	PicUrl    string `json:"pic_url" validate:"required"`
	State     int    `json:"state"`
	AdLink    string `json:"ad_link" validate:"required"`
	PostionId int    `json:"postion_id" validate:"required"`
	GotoType  int    `json:"goto_type"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
}

type NewAdRsp struct {
	AdId int64 `json:"ad_id"`
}

type GetAdsReq struct {
	Slogan     string `json:"slogan,omitempty"`
	PositionId int    `json:"position_id,omitempty"`
	StartTime  int64  `json:"start_time,omitempty"`
	EndTime    int64  `json:"end_time,omitempty"`
	Page       *Page  `json:"page"`
}

type GetAdsResp struct {
	Page *Page     `json:"page"`
	Ads  []*gen.Ad `json:"ads"`
}

type ChangeAdStateReq struct {
	Id int64 `json:"id" validate:"required"`
}

type DeleteAdStateReq struct {
	Id int64 `json:"id" validate:"required"`
}