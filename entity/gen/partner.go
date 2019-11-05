package gen

import (
	"time"
)

type Partner struct {
	Id             int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Logo           string    `json:"logo" xorm:"not null VARCHAR(255)"`
	EnterpriseName string    `json:"enterprise_name" xorm:"not null comment('企业名称') VARCHAR(255)"`
	Intro          string    `json:"intro" xorm:"not null VARCHAR(255)"`
	Weight         int       `json:"weight" xorm:"not null INT(11)"`
	SaveTime       time.Time `json:"save_time" xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	UpdateTime     time.Time `json:"update_time" xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
