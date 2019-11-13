package gen

import (
	"time"
)

type Account struct {
	Id         int64     `json:"id" xorm:"pk autoincr BIGINT(20)"`
	AvatarUrl  string    `json:"avatar_url" xorm:"default '0' comment('头像') VARCHAR(255)"`
	Name       string    `json:"name" xorm:"not null default '0' unique VARCHAR(50)"`
	Password   string    `json:"password" xorm:"not null default '0' VARCHAR(50)"`
	RoleId     int       `json:"role_id" xorm:"not null comment('角色id') INT(11)"`
	SaveTime   time.Time `json:"-" xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	UpdateTime time.Time `json:"-" xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
