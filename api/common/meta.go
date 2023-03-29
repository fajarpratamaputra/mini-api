package common

import "math"

type Pagination struct {
	TotalData   int64 `json:"total"`
	PageSize    int   `json:"per_page"`
	TotalPage   int64 `json:"total_page"`
	CurrentPage int   `json:"current_page"`
}

type MetaData struct {
	Pagination *Pagination `json:"pagination"`
	ImagePath  string      `json:"image_path"`
	VideoPath  string      `json:"video_path"`
}

type Status struct {
	Code          int    `json:"code"`
	MessageServer string `json:"message_server"`
	MessageClient string `json:"message_client"`
}

func (p *Pagination) BuildPagination(pageSize, currentPage int) {
	p.PageSize = pageSize
	p.CurrentPage = currentPage
	p.TotalPage = int64(math.Ceil(float64(p.TotalData) / float64(pageSize)))
	if p.TotalPage < 1 {
		p.TotalPage = 1
	}
}
