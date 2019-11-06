package gen

import (
	"time"
)

type Crm struct {
	Id          int64     `json:"id" xorm:"pk autoincr BIGINT(20)"`
	Name        string    `json:"name" xorm:"VARCHAR(50)"`
	Email       string    `json:"email" xorm:"VARCHAR(255)"`
	Phone       string    `json:"phone" xorm:"not null VARCHAR(32)"`
	OfficialWeb string    `json:"official_web" xorm:"not null comment('官网') VARCHAR(255)"`
	Company     string    `json:"company" xorm:"not null VARCHAR(255)"`
	Deleted     int       `json:"deleted" xorm:"not null comment('删除状态 1删除 0未删除') TINYINT(4)"`
	Message     string    `json:"message" xorm:"not null comment('留言') VARCHAR(1024)"`
	SaveTime    time.Time `json:"_" xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	UpdateTime  time.Time `json:"_" xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
