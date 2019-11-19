package entity

import "flower/entity/gen"

type BusinessPartnersPageReq struct {
	Page *Page `json:"page" validate:"required"`
}

type BusinessPartnersIdReq struct {
	Id int `json:"id"`
}

type BusinessPartnersReq struct {
	Id           int       `json:"id"`
	Logo         string    `json:"logo"`
	BusinessName string    `json:"business_name"`
	Intro        string    `json:"intro"`
	Sort         int       `json:"sort"`
}

type AdminBusinessPartnersRsp struct {
	Page *Page                   `json:"page"`
	Bps  []*gen.BusinessPartners `json:"bps"`
}

type IndexBusinessPartners struct {
	Id   int    `json:"id"`
	Logo string `json:"logo"`
}

type IndexBusinessPartnersRsp struct {
	Page *Page                   `json:"page"`
	Bps  []*IndexBusinessPartners `json:"bps"`
}
