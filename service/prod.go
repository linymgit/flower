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
	if query.Page == nil {
		err = session.Where(builder.Eq{"parent_id": query.ParentId}).Asc("sort").Find(&pcs)
		if err != nil {
			//TODO
		}
	}else{
		total, err = session.Where(builder.Eq{"parent_id": query.ParentId}).Asc("sort").Limit(query.Page.PageSize, query.Page.DbPageIndex()).FindAndCount(&pcs)
		if err != nil {
			//TODO
		}
	}
	return
}

func (p *ProdService) GetProductCategoryTree() (tree *entity.ProdCategoryTree, err error) {
	rows, err := mysql.Db.Rows(&gen.ProductCategory{})
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
	isExistName, err = mysql.Db.Where("name = ?", query.Name).Exist(&gen.ProductCategory{})
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
		level = query.ParentId+1
	}else{
		isExistParent = true
	}
	affected, err := mysql.Db.Cols("parent_id", "name", "desc", "states", "level", "sort").InsertOne(&gen.ProductCategory{
		ParentId: query.ParentId,
		Name:     query.Name,
		Desc:     query.Desc,
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

func (p *ProdService) ChangeProcategoryState(id int) (ok bool, err error){
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
	i, err := mysql.Db.Cols("name", "states", "heat","intro", "summary", "index_show", "details_pic_url", "cover_url", "price", "category_id", "author_id").InsertOne(&gen.Product{
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
	result, err := mysql.Db.Exec("INSERT INTO `product` (`name`,`intro`,`summary`,`states`,`index_show`," +
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
