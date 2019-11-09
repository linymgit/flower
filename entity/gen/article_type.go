package gen

import (
	"time"
)

type ArticleType struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	TypeName   string    `json:"type_name" xorm:"VARCHAR(50)"`
	Sort       int       `json:"sort" xorm:"INT(11)"`
	Level      int       `json:"level" xorm:"INT(11)"`
	ParentId   int       `json:"parent_id" xorm:"INT(11)"`
	SaveTime   time.Time `json:"_" xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	UpdateTime time.Time `json:"_" xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
