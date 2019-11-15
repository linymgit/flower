package service

import (
	"flower/entity"
	"flower/entity/gen"
	"flower/mysql"
	"sync"
	"time"
	"xorm.io/builder"
)

var Adsrv = &AdService{&sync.Mutex{}}

type AdService struct {
	m *sync.Mutex
}

func (ad *AdService) GetAds(query *entity.GetAdsReq) (ads []*gen.Ad, total int64, err error) {
	ads = make([]*gen.Ad, 0)
	cond := builder.NewCond()
	if query.Slogan != "" {
		cond = cond.And(builder.Eq{"slogan": query.Slogan})
	}
	if query.PositionId != 0 {
		cond = cond.And(builder.Eq{"postion_id": query.PositionId})
	}
	if query.EndTime > 0 && query.StartTime > 0 {
		begin := time.Unix(query.StartTime, 0).Add(-24 * time.Hour).Format(timeTemplate)
		end := time.Unix(query.EndTime, 0).Add(24 * time.Hour).Format(timeTemplate)
		cond = cond.And(builder.Lt{"end_time": &end})
		cond = cond.And(builder.Gt{"start_time": &begin})
	}
	session := mysql.Db.NewSession()
	defer session.Close()
	total, err = session.Where(cond).Desc("weight").Limit(query.Page.PageSize, query.Page.DbPageIndex()).FindAndCount(&ads)
	return
}

func (ad *AdService) NewAd(query *entity.NewAdReq) (adId int64, err error) {
	result, err := mysql.Db.Exec("INSERT INTO `ad` (`slogan`,`pic_url`,`postion_id`,`goto_type`,`ad_link`,`state`,"+
		"`clicks`,`weight`,`start_time`,`end_time`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		query.Slogan, query.PicUrl, query.PostionId, query.GotoType, query.AdLink, query.State, 0, 0,
		time.Unix(query.StartTime, 0).Format(timeTemplate), time.Unix(query.EndTime, 0).Format(timeTemplate))
	if err != nil {
		return
	}
	adId, err = result.LastInsertId()
	if err != nil {
		return
	}
	return
}

func (ad *AdService) ChangeAdState(adId int64) (ok bool, err error) {
	ad.m.Lock()
	defer ad.m.Unlock()
	result, err := mysql.Db.Exec("	UPDATE `ad` SET `state`=`state`^1 WHERE  `id`=?;", adId)
	if err != nil {
		return
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return
	}
	ok = affected == 1
	return
}

func (ad *AdService) DeleteAd(adId int64) (ok bool, err error) {
	ad.m.Lock()
	defer ad.m.Unlock()
	i, err := mysql.Db.ID(adId).Delete(&gen.Ad{})
	if err != nil {
		return
	}
	ok = i == 1
	return
}

func (ad *AdService) Modify(query *entity.ModifyAdReq) (ok bool, err error) {
	cols := []string{"state","goto_type"}
	if query.AdLink != "" {
		cols = append(cols, "ad_link")
	}
	if query.PicUrl != "" {
		cols = append(cols, "pic_url")
	}
	if query.Slogan != "" {
		cols = append(cols, "slogan")
	}
	if query.PostionId>0 {
		cols = append(cols, "postion_id")
	}
	if query.StartTime > 0 {
		cols = append(cols, "start_time")
	}
	if query.EndTime > 0 {
		cols = append(cols, "end_time")
	}
	affected, err := mysql.Db.ID(query.Id).Cols(cols...).Update(&gen.Ad{
		Slogan:    query.Slogan,
		PicUrl:    query.PicUrl,
		PostionId: query.PostionId,
		GotoType:  query.GotoType,
		AdLink:    query.AdLink,
		State:     query.State,
		StartTime: time.Unix(query.StartTime, 0),
		EndTime:   time.Unix(query.EndTime, 0),
	})
	ok = affected == 1
	return
}
