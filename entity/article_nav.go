package entity

import "sort"

type ArticleNavSort struct {
	ArtNav []ArticleNav
	by     func(p, q *ArticleNav) bool
}

type SortBy func(p, q *ArticleNav) bool

func (ans ArticleNavSort) Len() int {
	return len(ans.ArtNav)
}

func (ans ArticleNavSort) Swap(i, j int) {
	ans.ArtNav[i], ans.ArtNav[j] = ans.ArtNav[j], ans.ArtNav[i]
}
func (ans ArticleNavSort) Less(i, j int) bool {
	return ans.by(&ans.ArtNav[i], &ans.ArtNav[j])
}

func SortArticleNavs(ArtNav []ArticleNav, by SortBy){
	sort.Sort(ArticleNavSort{ArtNav, by})
}

// --------------v2.0------------------
type ArticleNavSortV2 struct {
	ArtNav []ArticleNavV2
	by     func(p, q *ArticleNavV2) bool
}

type SortByV2 func(p, q *ArticleNavV2) bool

func (ans ArticleNavSortV2) Len() int {
	return len(ans.ArtNav)
}

func (ans ArticleNavSortV2) Swap(i, j int) {
	ans.ArtNav[i], ans.ArtNav[j] = ans.ArtNav[j], ans.ArtNav[i]
}
func (ans ArticleNavSortV2) Less(i, j int) bool {
	return ans.by(&ans.ArtNav[i], &ans.ArtNav[j])
}

func SortArticleNavsV2(ArtNav []ArticleNavV2, by SortByV2){
	sort.Sort(ArticleNavSortV2{ArtNav, by})
}

var MonthSort = map[string]int{
	"January":   1,
	"February":  2,
	"March":     3,
	"April":     4,
	"May":       5,
	"June":      6,
	"July":      7,
	"August":    8,
	"September": 9,
	"October":   10,
	"November":  11,
	"December":  12,
}
