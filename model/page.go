package model

type Page struct {
	Books       []*Book // slice存储每页显示的图书
	PageNo      int64   // 当前页
	PageSize    int64   // 每页显示的条数
	TotalPageNo int64   // 总页数
	TotalRecord int64   // 总的记录数
	MinPrice    string
	MaxPrice    string
	IsLogin     bool
	Username    string
}

// IsHasPrev 判断是否有上一页
func (p *Page) IsHasPrev() bool {
	return p.PageNo > 1
}

// IsHasNext 判断是否有下一页
func (p *Page) IsHasNext() bool {
	return p.PageNo < p.TotalPageNo
}

// GetPrevPageNo 获取上一页
func (p *Page) GetPrevPageNo() int64 {
	if p.IsHasPrev() {
		return p.PageNo - 1
	}
	return 1
}

// GetNextPageNo 获取下一页
func (p *Page) GetNextPageNo() int64 {
	if p.IsHasNext() {
		return p.PageNo + 1
	}
	return p.TotalPageNo
}
