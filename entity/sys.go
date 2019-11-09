package entity

import "flower/entity/gen"

type SystemSettingReq struct {
	Name            string `json:"name" validate:"required"`
	Url             string `json:"url" validate:"required"`
	RectangleLogo   string `json:"rectangle_logo" validate:"required"`
	SquareLogo      string `json:"square_logo" validate:"required"`
	Address         string `json:"address" validate:"required"`
	EnterpriseEmail string `json:"enterprise_email" validate:"required"`
	Hotline         string `json:"hotline" validate:"required"`
	Icp             string `json:"icp" validate:"required"`
}

type GetSystemSettingRsp struct {
	Setting *gen.WebSetting `json:"setting"`
}