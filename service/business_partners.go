package service

import (
	"flower/entity"
	"flower/entity/gen"
	"flower/mysql"
)

var BusinessPartnersSrv = &businessPartnersService{}

type businessPartnersService struct {
}

func (service businessPartnersService) AdminListBusinessPartners(req *entity.BusinessPartnersPageReq) (bps []*gen.BusinessPartners, total int64, err error) {
	bps = make([]*gen.BusinessPartners, 0)
	session := mysql.Db.NewSession()
	defer session.Close()
	total, err = session.Asc("sort", "id").Limit(req.Page.PageSize, req.Page.DbPageIndex()).FindAndCount(&bps)
	return
}

func (service businessPartnersService) DeleteBusinessPartnersById(id int) (ok bool, err error) {
	affected, err := mysql.Db.ID(id).Delete(&gen.BusinessPartners{})
	if err != nil {
		return
	}
	ok = affected == 1
	return
}

func (service businessPartnersService) AddBusinessPartners(req *entity.BusinessPartnersReq) (ok bool, err error) {
	affected, err := mysql.Db.Cols("logo", "business_name", "intro", "sort").InsertOne(&gen.BusinessPartners{
		Logo:         req.Logo,
		BusinessName: req.BusinessName,
		Intro:        req.Intro,
		Sort:         req.Sort,
	})
	if err != nil {
		return
	}
	ok = affected == 1
	return
}

func (service businessPartnersService) ModifyBusinessPartners(req *entity.BusinessPartnersReq) (ok bool, err error){
	cols := []string{}
	if req.Logo != "" {
		cols = append(cols, "logo")
	}
	if req.Intro != "" {
		cols = append(cols, "intro")
	}
	if req.BusinessName != "" {
		cols = append(cols, "business_name")
	}
	if req.Sort > 0 {
		cols = append(cols, "sort")
	}
	if len(cols) <= 0 {
		return
	}
	affected, err := mysql.Db.Id(req.Id).Cols(cols...).Update(&gen.BusinessPartners{
		Logo:         req.Logo,
		BusinessName: req.BusinessName,
		Intro:        req.Intro,
		Sort:         req.Sort,
	})
	if err != nil {
		return
	}
	ok = affected == 1
	return
}

func (service businessPartnersService) FrontListBusinessPartners(req *entity.BusinessPartnersPageReq) (bps []*gen.BusinessPartners, total int64, err error) {
	bps = make([]*gen.BusinessPartners, 0)
	session := mysql.Db.NewSession()
	defer session.Close()
	total, err = session.Asc("sort", "id").Limit(req.Page.PageSize, req.Page.DbPageIndex()).Cols("id","logo").FindAndCount(&bps)
	return
}
