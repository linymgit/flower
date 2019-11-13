package gen

import (
	"time"
)

type Product struct {
	Id            int64     `json:"id" xorm:"pk autoincr BIGINT(20)"`
	Name          string    `json:"name" xorm:"not null comment('商品名称') VARCHAR(255)"`
	Intro         string    `json:"intro" xorm:"not null comment('简介') VARCHAR(255)"`
	Summary       string    `json:"summary" xorm:"not null VARCHAR(1024)"`
	States        int       `json:"states" xorm:"not null comment('0在线 1下线') TINYINT(4)"`
	IndexShow     int       `json:"index_show" xorm:"not null comment('首页推荐 1是 0否') TINYINT(4)"`
	DetailsPicUrl string    `json:"details_pic_url" xorm:"not null VARCHAR(255)"`
	CoverUrl      string    `json:"cover_url" xorm:"not null VARCHAR(255)"`
	Price         string    `json:"price" xorm:"not null DECIMAL(10)"`
	Heat          int       `json:"heat" xorm:"not null comment('热度') INT(11)"`
	CategoryId    int       `json:"category_id" xorm:"not null comment('类目id') INT(11)"`
	AuthorId      int64     `json:"author_id" xorm:"not null comment('作者id') BIGINT(20)"`
	UpdateTime    time.Time `json:"-" xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	SaveTime      time.Time `json:"-" xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	CategoryName  string    `json:"category_name" xorm:"-"`
}
