package gen

import (
	"time"
)

type BusinessPartners struct {
	Id           int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	BusinessName string    `json:"business_name" xorm:"not null default '0' comment('企业名称') VARCHAR(64)"`
	Intro        string    `json:"intro" xorm:"not null default '0' comment('企业介绍') VARCHAR(256)"`
	Sort         int       `json:"sort" xorm:"not null default 0 comment('排序') INT(11)"`
	SaveTime     time.Time `json:"save_time" xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME"`
	UpdateTime   time.Time `json:"update_time" xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME"`
}
