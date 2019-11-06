package entity

type Page struct {
	PageSize  int   `json:"page_size"`
	PageIndex int   `json:"page_index"`
	Total     int64 `json:"total"`
}

func (p *Page) DbPageIndex() (index int) {
	index = (p.PageIndex - 1) * p.PageSize
	return
}
