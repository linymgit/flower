package gen

import (
	"time"
)

type WebSetting struct {
	Id              int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Name            string    `json:"name" xorm:"not null comment('公司名称') VARCHAR(255)"`
	Url             string    `json:"url" xorm:"not null comment('官网地址') VARCHAR(255)"`
	RectangleLogo   string    `json:"rectangle_logo" xorm:"not null comment('长方形的logo') VARCHAR(255)"`
	SquareLogo      string    `json:"square_logo" xorm:"not null comment('正方形的logo') VARCHAR(255)"`
	Address         string    `json:"address" xorm:"not null comment('公司地址') VARCHAR(255)"`
	EnterpriseEmail string    `json:"enterprise_email" xorm:"not null comment('企业邮箱') VARCHAR(255)"`
	Hotline         string    `json:"hotline" xorm:"not null comment('服务热线') VARCHAR(32)"`
	Icp             string    `json:"icp" xorm:"not null comment('网站ICP备案号') VARCHAR(32)"`
	SaveTime        time.Time `json:"_" xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	UpdateTime      time.Time `json:"_" xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
