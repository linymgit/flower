package gen

import (
	"time"
)

type Article struct {
	Id         int64     `json:"id" xorm:"pk autoincr BIGINT(20)"`
	TypeId     int       `json:"type_id" xorm:"not null INT(11)"`
	Title      string    `json:"title" xorm:"not null VARCHAR(255)"`
	Author     string    `json:"author" xorm:"not null VARCHAR(50)"`
	Source     string    `json:"source" xorm:"not null VARCHAR(255)"`
	SourceUrl  string    `json:"source_url" xorm:"not null VARCHAR(255)"`
	Preview    string    `json:"preview" xorm:"not null VARCHAR(50)"`
	KeyWord    string    `json:"key_word" xorm:"not null VARCHAR(50)"`
	Summary    string    `json:"summary" xorm:"not null VARCHAR(255)"`
	Content    string    `json:"content" xorm:"not null TEXT"`
	Clicks     int       `json:"clicks" xorm:"not null index INT(11)"`
	States     int       `json:"states" xorm:"not null TINYINT(4)"`
	Weight     int       `json:"weight" xorm:"not null INT(11)"`
	SaveTime   time.Time `json:"save_time" xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	UpdateTime time.Time `json:"update_time" xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
