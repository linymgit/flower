package gen

import (
	"time"
)

type ProductUv struct {
	Id         int64     `json:"id" xorm:"pk autoincr BIGINT(20)"`
	PId        int64     `json:"p_id" xorm:"not null default 0 comment('产品id') index BIGINT(20)"`
	Ip         string    `json:"ip" xorm:"not null default '0' comment('客户端ip') VARCHAR(32)"`
	AccessTime time.Time `json:"access_time" xorm:"not null default 'CURRENT_TIMESTAMP' comment('访问时间') DATETIME"`
}
