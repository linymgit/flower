package entity

type FrontArticleType struct {
	Id       int    `json:"id"`
	TypeName string `json:"type_name"`
}

type FrontArticleTypeRsp struct {
	Categories []*FrontArticleType `json:"categories"`
}

type FrontGetArticleReq struct {
	Id int64 `json:"id"`
}
