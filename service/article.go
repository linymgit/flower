package service

import (
	"flower/entity"
	"flower/entity/gen"
	"flower/mysql"
	"sync"
	"time"
	"xorm.io/builder"
)

var ArticleSrv = &ArticleService{&sync.Mutex{}}

type ArticleService struct {
	m *sync.Mutex
}

func (ac *ArticleService) ListArticleType(query *entity.ListArticleTypeReq) (ats []*gen.ArticleType, total int64, err error) {
	ats = make([]*gen.ArticleType, 0)
	session := mysql.Db.NewSession()
	defer session.Close()
	if query.Page == nil {
		err = session.Where(builder.Eq{"parent_id": query.ParentId}).Asc("sort").Find(&ats)
		if err != nil {
			//TODO
		}
	} else {
		total, err = session.Where(builder.Eq{"parent_id": query.ParentId}).Asc("sort").Limit(query.Page.PageSize, query.Page.DbPageIndex()).FindAndCount(&ats)
		if err != nil {
			//TODO
		}
	}
	return
}

func (ac *ArticleService) GetArticleCategoryTree() (tree *entity.ArticleTypeTree, err error) {
	rows, err := mysql.Db.Asc("sort").Rows(&gen.ArticleType{})
	if err != nil {
	}
	defer rows.Close()
	vosMap := make(map[int][]*entity.ArticleTypeVo, 0)
	bean := new(gen.ArticleType)
	for rows.Next() {
		err = rows.Scan(bean)
		if err != nil {
			return
		}
		vo := &entity.ArticleTypeVo{
			Id:       bean.Id,
			TypeName: bean.TypeName,
			Sort:     bean.Sort,
			Level:    bean.Level,
			ParentId: bean.ParentId,
		}
		if vos, ok := vosMap[vo.Level]; ok {
			vos = append(vos, vo)
			vosMap[vo.Level] = vos
		} else {
			vos := make([]*entity.ArticleTypeVo, 0)
			vos = append(vos, vo)
			vosMap[vo.Level] = vos
		}
	}
	for k := range vosMap {
		if k == 1 {
			continue
		}
		vos := vosMap[k]
		preVos := vosMap[k-1]
		for p := range preVos {
			tempSub := make([]*entity.ArticleTypeVo, 0)
			for s := range vos {
				if preVos[p].Id == vos[s].ParentId {
					tempSub = append(tempSub, vos[s])
				}
			}
			if len(tempSub) > 0 {
				preVos[p].SubArticleType = tempSub
			}
		}
	}
	tree = &entity.ArticleTypeTree{Tree: vosMap[1]}
	return
}

func (ac *ArticleService) NewArticleType(query *entity.NewArticleTypeReq) (isExistName, isExistParent, ok bool, err error) {
	ac.m.Lock()
	defer ac.m.Unlock()
	// 分类名称唯一校验
	isExistName, err = mysql.Db.Where("type_name = ?", query.TypeName).Exist(&gen.ArticleType{})
	if err != nil {
		return
	}
	if isExistName {
		return
	}
	level := 1
	if query.ParentId != 0 {
		parent := &gen.ArticleType{}
		isExistParent, err = mysql.Db.Id(query.ParentId).Get(parent)
		if !isExistParent {
			//TODO
			return
		}
		level = query.ParentId + 1
	} else {
		isExistParent = true
	}
	affected, err := mysql.Db.Cols("type_name", "sort", "level", "parent_id").InsertOne(&gen.ArticleType{
		TypeName: query.TypeName,
		Sort:     query.Sort,
		Level:    level,
		ParentId: query.ParentId,
	})
	if err != nil {
		//TODO
		return
	}
	ok = affected == 1
	return
}

func (ac *ArticleService) EditArticle(query *entity.EditArticleTypeReq) (ok, existParent bool, err error) {
	ac.m.Lock()
	defer ac.m.Unlock()
	level := 1
	if query.ParentId != 0 {
		parent := &gen.ArticleType{}
		existParent, err = mysql.Db.Id(query.ParentId).Get(parent)
		if !existParent {
			//TODO
			return
		}
		level = query.ParentId + 1
	} else {
		existParent = true
	}
	i, err := mysql.Db.Id(query.Id).Cols("type_name", "sort", "level", "parent_id").Update(&gen.ArticleType{
		TypeName: query.TypeName,
		Sort:     query.Sort,
		Level:    level,
		ParentId: query.ParentId,
	})
	ok = i == 1
	return
}

func (ac *ArticleService) NewArticle(query *entity.NewArticleReq) (articleId int64, err error) {
	result, err := mysql.Db.Exec("INSERT INTO `article` (`type_id`,`title`,`author`,`source`,`source_url`,"+
		"`preview`,`key_word`,`summary`,`content`,`clicks`,`states`,`sort`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		query.TypeId, query.Title, query.Author, query.Source, query.SourceUrl, query.Preview, query.KeyWord, query.Summary,
		query.Content, 0, query.States, query.Sort)
	if err != nil {
		return
	}
	articleId, err = result.LastInsertId()
	if err != nil {
		return
	}
	return
}

func (ac *ArticleService) ListArticle(query *entity.ListArticleReq) (as []*gen.Article, total int64, err error) {
	as = make([]*gen.Article, 0)
	cond := builder.NewCond()
	if query.Title != "" {
		cond = cond.And(builder.Eq{"title": query.Title})
	}
	if query.TypeId != 0 {
		cond = cond.And(builder.Eq{"type_id": query.TypeId})
	}
	if query.PublishStartTime > 0 && query.PublishEndTime > 0 {
		start := time.Unix(query.PublishStartTime, 0).Add(-24 * time.Hour).Format(timeTemplate)
		end := time.Unix(query.PublishEndTime, 0).Add(24 * time.Hour).Format(timeTemplate)
		cond = cond.And(builder.Gt{"save_time": start}).And(builder.Lt{"save_time": end})
	}
	total, err = mysql.Db.Where(cond).Limit(query.Page.PageSize, query.Page.DbPageIndex()).FindAndCount(&as)
	return
}

func (ac *ArticleService) ChangeOnline(id int64) (ok bool, err error) {
	result, err := mysql.Db.Exec("UPDATE `article` SET `states`=`states`^1 WHERE  `id`=?;", id)
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

func (ac *ArticleService) DeleteArticle(id int64) (ok bool, err error) {
	affected, err := mysql.Db.ID(id).Delete(&gen.Article{})
	if err != nil {
		return
	}
	ok = affected == 1
	return
}

func (ac *ArticleService) ModifyArticle(query *entity.ModifyArticleReq) (ok bool, err error) {
	cols := []string{}
	if query.TypeId > 0 {
		cols = append(cols, "type_id")
	}
	if query.Title != "" {
		cols = append(cols, "title")
	}
	if query.Author != "" {
		cols = append(cols, "author")
	}
	if query.Source != "" {
		cols = append(cols, "source")
	}
	if query.SourceUrl != "" {
		cols = append(cols, "source_url")
	}
	if query.Preview != "" {
		cols = append(cols, "preview")
	}
	if query.KeyWord != "" {
		cols = append(cols, "key_word")
	}
	if query.Summary != "" {
		cols = append(cols, "summary")
	}
	if query.Content != "" {
		cols = append(cols, "content")
	}
	if query.Sort > 0 {
		cols = append(cols, "sort")
	}
	if len(cols) <= 0 {
		ok = false
		return
	}
	affected, err := mysql.Db.ID(query.Id).Cols(cols...).Update(&gen.Article{
		TypeId:    query.TypeId,
		Title:     query.Title,
		Author:    query.Author,
		Source:    query.Source,
		SourceUrl: query.SourceUrl,
		Preview:   query.Preview,
		KeyWord:   query.KeyWord,
		Summary:   query.Summary,
		Content:   query.Content,
		Sort:      query.Sort,
	})
	if err != nil {
		return
	}
	ok = affected == 1
	return
}

func (ac *ArticleService) DeleteAricleTypeById(id int)(isParent,ok bool, err error) {
	isParent, err = mysql.Db.Where("parent_id = ?", id).Cols("id").Exist(&gen.ArticleType{})
	if err != nil {
		return
	}
	if isParent {
		return
	}
	var affected int64
	affected, err = mysql.Db.Id(id).Delete(&gen.ArticleType{})
	ok = affected == 1
	return
}
