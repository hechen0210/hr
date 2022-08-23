package util

import (
	"math"

	"github.com/jinzhu/gorm"
)

type Page struct {
	Total       int         `json:"total"`
	PerPage     int         `json:"per_page"`
	TotalPage   int         `json:"total_page"`
	CurrentPage int         `json:"current_page"`
	List        interface{} `json:"list"`
}

func NewPage(total, perPage, currentPage int, list interface{}) Page {
	return Page{
		Total:       total,
		PerPage:     perPage,
		TotalPage:   int(math.Ceil(float64(total) / float64(perPage))),
		CurrentPage: currentPage,
		List:        list,
	}
}

func Paginate(page, perPage int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset((page - 1) * perPage).Limit(perPage)
	}
}
