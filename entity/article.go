package entity

import "flower/entity/gen"

type ArticleTypeVo struct {
	Id             int              `json:"id"`
	TypeName       string           `json:"type_name"`
	Sort           int              `json:"sort"`
	Level          int              `json:"-"`
	ParentId       int              `json:"-"`
	SubArticleType []*ArticleTypeVo `json:"sub_article_type"`
}

type ArticleTypeTree struct {
	Tree []*ArticleTypeVo `json:"tree"`
}

type NewArticleTypeReq struct {
	TypeName string `json:"type_name" validate:"required"`
	Sort     int    `json:"sort" validate:"required"`
	ParentId int    `json:"parent_id"`
}

type ListArticleTypeReq struct {
	ParentId int   `json:"parent_id"`
	Id       int   `json:"id"`
	Page     *Page `json:"page"`
}

type ListArticleTypeRsp struct {
	Page *Page              `json:"page,omitempty"`
	Ats  []*gen.ArticleType `json:"ats"`
}

type EditArticleTypeReq struct {
	Id       int    `json:"id"`
	TypeName string `json:"type_name" validate:"required"`
	Sort     int    `json:"sort" validate:"required"`
	ParentId int    `json:"parent_id"`
}

type NewArticleReq struct {
	TypeId    int    `json:"type_id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Source    string `json:"source"`
	SourceUrl string `json:"source_url"`
	Preview   string `json:"preview"`
	KeyWord   string `json:"key_word"`
	Summary   string `json:"summary"`
	Content   string `json:"content"`
	States    int    `json:"states"`
	Sort      int    `json:"sort" xorm:"not null INT(11)"`
}

type NewArticleRsp struct {
	ArticleId int64 `json:"article_id"`
}

type ListArticleReq struct {
	Title            string `json:"title"`
	TypeId           int    `json:"type_id"`
	PublishStartTime int64  `json:"publish_start_time"`
	PublishEndTime   int64  `json:"publish_end_time"`
	Page             *Page  `json:"page" validate:"required"`
}

type ListArticleRsp struct {
	Page *Page          `json:"page"`
	As   []*gen.Article `json:"as"`
}

type ChangeOnlineReq struct {
	Id int64 `json:"id" validate:"required"`
}

type DeleteArticleReq struct {
	Id int64 `json:"id" validate:"required"`
}

type ModifyArticleReq struct {
	Id        int64  `json:"id" validate:"required"`
	TypeId    int    `json:"type_id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Source    string `json:"source"`
	SourceUrl string `json:"source_url"`
	Preview   string `json:"preview"`
	KeyWord   string `json:"key_word"`
	Summary   string `json:"summary"`
	Content   string `json:"content"`
	Sort      int    `json:"sort" xorm:"not null INT(11)"`
}

type DeleteAricleTypeByIdReq struct {
	Id int `json:"id" validate:"required"`
}
