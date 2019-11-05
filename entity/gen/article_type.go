package gen

import (
	"time"
)

type ArticleType struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	TypeName   string    `json:"type_name" xorm:"VARCHAR(50)"`
	Weight     int       `json:"weight" xorm:"INT(11)"`
	ParentId   int64     `json:"parent_id" xorm:"BIGINT(20)"`
	SaveTime   time.Time `json:"save_time" xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	UpdateTime time.Time `json:"update_time" xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
