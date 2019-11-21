package service

import (
	"flower/entity"
	"fmt"
	"strings"
	"testing"
)

func TestSort(t *testing.T) {
	navs := make([]entity.ArticleNav, 0)

	navs = append(navs, entity.ArticleNav{
		Time:  "2019-November",
		Count: 1,
	})
	navs = append(navs, entity.ArticleNav{
		Time:  "2018-November",
		Count: 1,
	})
	navs = append(navs, entity.ArticleNav{
		Time:  "2019-December",
		Count: 1,
	})

	entity.SortArticleNavs(navs, func(p, q *entity.ArticleNav)bool{
		pt := p.Time
		qt := q.Time
		spP := strings.Split(pt, "-")
		spQ := strings.Split(qt, "-")
		if spP[0] == spQ[0] {
			return entity.MonthSort[spP[1]]>entity.MonthSort[spQ[1]]
		}else{
			return spP[0]>spQ[0]
		}
		return false
	})

	for k := range navs {
		fmt.Printf("%#v\n", navs[k])
	}
	
}
