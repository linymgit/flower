package gen

type AdPosition struct {
	Id         int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Expression string `json:"expression" xorm:"not null default '0' comment('位置表达式') VARCHAR(255)"`
	Name       string `json:"name" xorm:"not null default '0' comment('广告位') VARCHAR(50)"`
}
