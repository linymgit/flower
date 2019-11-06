package gen

import (
	"time"
)

type ProductCategory struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	ParentId   int       `json:"parent_id" xorm:"not null INT(11)"`
	Name       string    `json:"name" xorm:"not null VARCHAR(50)"`
	Desc       string    `json:"desc" xorm:"not null VARCHAR(255)"`
	States     int       `json:"states" xorm:"not null comment('0显示 1不显示') TINYINT(4)"`
	Level      int       `json:"level" xorm:"not null comment('类目级别') INT(11)"`
	Weight     int       `json:"weight" xorm:"not null INT(11)"`
	SaveTime   time.Time `json:"_" xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	UpdateTime time.Time `json:"_" xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
