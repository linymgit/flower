package gen

import (
	"time"
)

type Ad struct {
	Id         int64     `json:"id" xorm:"pk autoincr BIGINT(20)"`
	Slogan     string    `json:"slogan" xorm:"not null default '0' comment('广告语') VARCHAR(255)"`
	PicUrl     string    `json:"pic_url" xorm:"not null default '0' VARCHAR(255)"`
	PostionId  int       `json:"postion_id" xorm:"not null default 0 comment('广告位置') INT(11)"`
	GotoType   int       `json:"goto_type" xorm:"not null default 0 comment('跳转类型 0:url 1:to product') TINYINT(4)"`
	State      int       `json:"state" xorm:"not null default 0 comment('0在线 1下线') TINYINT(4)"`
	Clicks     int       `json:"clicks" xorm:"not null default 0 comment('点击数') INT(11)"`
	Weight     int       `json:"weight" xorm:"not null default 0 comment('权重用于排序，数值越大权重越大') INT(11)"`
	StartTime  time.Time `json:"start_time" xorm:"not null TIMESTAMP"`
	EndTime    time.Time `json:"end_time" xorm:"not null TIMESTAMP"`
	SaveTime   time.Time `json:"_" xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	UpdateTime time.Time `json:"_" xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
