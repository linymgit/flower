package service

import (
	"flower/entity"
	"flower/entity/gen"
	"flower/entity/state"
	"flower/mysql"
	"sync"
	"xorm.io/builder"
)

var ProdSrv = &ProdService{new(sync.Mutex)}

type ProdService struct {
	m *sync.Mutex
}

func (p *ProdService) ListProductCategory(query *entity.ListProductCategoryReq) (pcs []*gen.ProductCategory, total int64, err error) {
	pcs = make([]*gen.ProductCategory, 0)
	session := mysql.Db.NewSession()
	defer session.Close()
	cond := builder.NewCond()
	if query.Id > 0 {
		cond = cond.And(builder.Eq{"id": query.Id})
	} else {
		cond = cond.And(builder.Eq{"parent_id": query.ParentId})
	}
	if query.Page == nil {
		err = session.Where(cond).Asc("sort").Find(&pcs)
		if err != nil {
			//TODO
		}
	} else {
		total, err = session.Where(cond).Asc("sort").Limit(query.Page.PageSize, query.Page.DbPageIndex()).FindAndCount(&pcs)
		if err != nil {
			//TODO
		}
	}
	return
}

func (p *ProdService) GetProductCategoryTree() (tree *entity.ProdCategoryTree, err error) {
	rows, err := mysql.Db.Asc("sort").Rows(&gen.ProductCategory{})
	if err != nil {
	}
	defer rows.Close()
	vosMap := make(map[int][]*entity.ProdCategoryVo, 0)
	bean := new(gen.ProductCategory)
	for rows.Next() {
		err = rows.Scan(bean)
		if err != nil {
			return
		}
		vo := &entity.ProdCategoryVo{
			Id:       bean.Id,
			Name:     bean.Name,
			Desc:     bean.Desc,
			States:   bean.States,
			Level:    bean.Level,
			Sort:     bean.Sort,
			ParentId: bean.ParentId,
		}
		if vos, ok := vosMap[vo.Level]; ok {
			vos = append(vos, vo)
			vosMap[vo.Level] = vos
		} else {
			vos := make([]*entity.ProdCategoryVo, 0)
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
			tempSub := make([]*entity.ProdCategoryVo, 0)
			for s := range vos {
				if preVos[p].Id == vos[s].ParentId {
					tempSub = append(tempSub, vos[s])
				}
			}
			if len(tempSub) > 0 {
				preVos[p].Sub = tempSub
			}
		}
	}
	tree = &entity.ProdCategoryTree{Tree: vosMap[1]}
	return
}

func (p *ProdService) NewProductCategory(query *entity.NewProdCategoryReq) (isExistName, isExistParent, ok bool, err error) {
	p.m.Lock()
	defer p.m.Unlock()
	// 分类名称唯一校验
	isExistName, err = mysql.Db.Where("name = ?", query.Name).Cols("id").Exist(&gen.ProductCategory{})
	if err != nil {
		return
	}
	if isExistName {
		return
	}
	level := 1
	if query.ParentId != 0 {
		parent := &gen.ProductCategory{}
		isExistParent, err = mysql.Db.Id(query.ParentId).Get(parent)
		if !isExistParent {
			//TODO
			return
		}
		level = parent.Level + 1
	} else {
		isExistParent = true
	}
	affected, err := mysql.Db.Cols("parent_id", "name", "desc", "states", "level", "sort").InsertOne(&gen.ProductCategory{
		ParentId: query.ParentId,
		Name:     query.Name,
		Desc:     query.Desc,
		States:   state.ProdCategoryShow,
		Level:    level,
		Sort:     query.Sort,
	})
	if err != nil {
		//TODO
		return
	}
	ok = affected == 1
	return
}

func (p *ProdService) ChangeProcategoryState(id int) (ok bool, err error) {
	result, err := mysql.Db.Exec("	UPDATE `product_category` SET `states`=`states`^1 WHERE  `id`=?;", id)
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

func (p *ProdService) NewProduct(query *entity.NewProductReq) (ok bool, err error) {
	i, err := mysql.Db.Cols("name", "states", "heat", "intro", "summary", "index_show", "details_pic_url", "cover_url", "price", "category_id", "author_id").InsertOne(&gen.Product{
		Name:          query.Name,
		Intro:         query.Intro,
		Summary:       query.Summary,
		IndexShow:     query.IndexShow,
		DetailsPicUrl: query.DetailsPicUrl,
		CoverUrl:      query.CoverUrl,
		Price:         query.Price,
		CategoryId:    query.CategoryId,
		AuthorId:      query.AuthorId,
	})
	if err != nil {
		return
	}
	ok = i == 1
	return
}

func (p *ProdService) NewProductResutId(query *entity.NewProductReq) (productId int64, err error) {
	result, err := mysql.Db.Exec("INSERT INTO `product` (`name`,`intro`,`summary`,`states`,`index_show`,"+
		"`details_pic_url`,`cover_url`,`price`,`heat`,`category_id`,`author_id`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		query.Name, query.Intro, query.Summary, state.ProdOnline, query.IndexShow, query.DetailsPicUrl, query.CoverUrl,
		query.Price, 0, query.CategoryId, query.AuthorId)
	if err != nil {
		return
	}
	productId, err = result.LastInsertId()
	if err != nil {
		return
	}
	return
}

func (p *ProdService) ModifyProcategory(query *entity.ModifyCategoryReq) (affected int64, err error) {
	affected, err = updateProcategory(&gen.ProductCategory{
		Id:     query.Id,
		Name:   query.Name,
		Desc:   query.Desc,
		States: query.States,
		Sort:   query.Sort,
	}, "name", "desc", "states", "sort")
	return
}

func (p *ProdService) DeleteProcategoryById(id int) (isParent bool, affected int64, err error) {
	isParent, err = mysql.Db.Where("parent_id = ?", id).Cols("id").Exist(&gen.ProductCategory{})
	if err != nil {
		return
	}
	if isParent {
		return
	}
	affected, err = mysql.Db.Id(id).Delete(&gen.ProductCategory{})
	return
}

func (p *ProdService) ListProduct(query *entity.ListProductReq) (ps []*gen.Product, total int64, err error) {
	ps = make([]*gen.Product, 0)
	cond := builder.NewCond()
	if query.States != state.ProdAll {
		cond = cond.And(builder.Eq{"states": query.States})
	}
	if query.Name != "" {
		cond = cond.And(builder.Eq{"name": query.Name})
	}
	if query.CategoryId != 0 {
		cond = cond.And(builder.Eq{"category_id": query.CategoryId})
	}
	session := mysql.Db.NewSession()
	defer session.Close()
	total, err = session.Where(cond).Asc("save_time").Limit(query.Page.PageSize, query.Page.DbPageIndex()).FindAndCount(&ps)
	return
}

func updateProcategory(procategory *gen.ProductCategory, columns ...string) (affected int64, err error) {
	affected, err = mysql.Db.Id(procategory.Id).Cols(columns...).Update(procategory)
	return
}

func (p *ProdService) CategoryId2Name() (id2nameMap map[int]string, err error) {
	rows, err := mysql.Db.Cols("id", "name").Rows(&gen.ProductCategory{})
	bean := new(gen.ProductCategory)
	id2nameMap = make(map[int]string)
	for rows.Next() {
		err = rows.Scan(bean)
		if err != nil {
			return
		}
		id2nameMap[bean.Id] = bean.Name
	}
	return
}

func (p *ProdService) ChangeProductState(id int64) (ok bool, err error) {
	result, err := mysql.Db.Exec("UPDATE `product` SET `states`=`states`^1 WHERE  `id`=?;", id)
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

func (p *ProdService) ChangeProductIndexShow(id int64) (ok bool, err error) {
	result, err := mysql.Db.Exec("UPDATE `product` SET `index_show`=`index_show`^1 WHERE  `id`=?;", id)
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

func (p *ProdService) ModifyProduct(query *entity.ModifyProductReq) (ok bool, err error) {
	cols := []string{}
	if query.Name != "" {
		cols = append(cols, "name")
	}
	if query.Intro != "" {
		cols = append(cols, "intro")
	}
	if query.Summary != "" {
		cols = append(cols, "summary")
	}
	if query.DetailsPicUrl != "" {
		cols = append(cols, "details_pic_url")
	}
	if query.CoverUrl != "" {
		cols = append(cols, "cover_url")
	}
	if query.Price != "" {
		cols = append(cols, "price")
	}
	if query.CategoryId > 0 {
		cols = append(cols, "category_id")
	}
	if len(cols) <= 0 {
		ok = false
		return
	}
	affected, err := mysql.Db.ID(query.Id).Cols(cols...).Update(&gen.Product{
		Name:          query.Name,
		Intro:         query.Intro,
		Summary:       query.Summary,
		DetailsPicUrl: query.DetailsPicUrl,
		CoverUrl:      query.CoverUrl,
		Price:         query.Price,
		CategoryId:    query.CategoryId,
		AuthorId:      query.AuthorId,
	})
	if err != nil {
		return
	}
	ok = affected == 1
	return
}

func (p *ProdService) DeleteProductById(id int64) (ok bool, err error) {
	affected, err := mysql.Db.ID(id).Delete(&gen.Product{})
	if err != nil {
		return
	}
	ok = affected == 1
	return
}
