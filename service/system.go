package service

import (
	"flower/entity"
	"flower/entity/gen"
	"flower/mysql"
	"sync"
)

var SysSrv = &SystemService{new(sync.Mutex)}

type SystemService struct {
	m *sync.Mutex
}

// 插入网站设置
func (s *SystemService)InsertSystemSetting(setting *entity.SystemSettingReq) (isExist bool, ok bool, err error) {
	b, err := mysql.Db.IsTableEmpty(&gen.WebSetting{})
	if err != nil {
		return
	}
	if !b {
		isExist = true
		return
	}
	i, err := mysql.Db.Cols("id", "name", "url", "rectangle_logo",
		"square_logo", "address", "enterprise_email", "hotline", "icp").InsertOne(&gen.WebSetting{
		Name:            setting.Name,
		Url:             setting.Url,
		RectangleLogo:   setting.RectangleLogo,
		SquareLogo:      setting.SquareLogo,
		Address:         setting.Address,
		EnterpriseEmail: setting.EnterpriseEmail,
		Hotline:         setting.Hotline,
		Icp:             setting.Icp,
		})
	if err != nil {
		return
	}
	ok = i == 1
	return
}

func (s *SystemService) UpdateSystemSetting(setting *entity.SystemSettingReq) (ok bool, err error) {
	i, err := mysql.Db.Update(&gen.WebSetting{
		Name:            setting.Name,
		Url:             setting.Url,
		RectangleLogo:   setting.RectangleLogo,
		SquareLogo:      setting.SquareLogo,
		Address:         setting.Address,
		EnterpriseEmail: setting.EnterpriseEmail,
		Hotline:         setting.Hotline,
		Icp:             setting.Icp,
	})
	if err != nil {
		return
	}
	ok = i == 1
	return 
}

func (s *SystemService) SelectSystemSetting() (ok bool, webSetting *gen.WebSetting, err error) {
	webSetting = &gen.WebSetting{}
	ok, err = mysql.Db.Limit(1,0).Get(webSetting)
	if err != nil {
		return
	}
	return
}
