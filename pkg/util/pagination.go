package util

import "math"

type Pagination struct {
	Model      string
	Count      int64
	Page       int64
	TotalPages int64
	NextPage   int64
	PrevPage   int64
}

func GetTotalPagesSize(count int64, limit int64) int64 {
	pages := float64(count) / float64(limit)
	return int64(math.Ceil(pages))
}

func ProcessPagination(model string, count int64, page int64, limit int64) Pagination {
	var pagination Pagination

	pagination.Model = model
	pagination.Count = count
	pagination.Page = page
	pagination.TotalPages = GetTotalPagesSize(count, limit)
	pagination.NextPage = page + 1
	pagination.PrevPage = page - 1

	return pagination
}
