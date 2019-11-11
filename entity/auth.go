package entity

type LoginReq struct {
	Id          string `json:"id"` //验证码的标识id
	VerifyValue string `json:"verify_value"`
	Name        string `json:"name" validate:"required" validate:"required"`
	Password    string `json:"password" validate:"required"`
}

type LoginResp struct {
	Token string `json:"token"`
}

type ModifyPasswordReq struct {
	Id          string `json:"id"` //验证码的标识id
	VerifyValue string `json:"verify_value"`
	Password    string `json:"password" validate:"required"`
}
