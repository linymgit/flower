package entity

import "flower/entity/gen"

type CrmListReq struct {
	Page      *Page  `json:"page" validate:"required"`
	Name      string `json:"name,omitempty"`
	BeginTime int64  `json:"begin_time"`
	EndTime   int64  `json:"end_time"`
}

type CrmListResp struct {
	Page *Page      `json:"page"`
	Crms []*gen.Crm `json:"crms"`
}

type CrmDeleteReq struct {
	Id int64 `json:"id" validate:"required"`
}
