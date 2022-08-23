package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Condition struct {
	Name   string
	Symbol string
	Value  interface{}
}

// 查询条件
func GetBy(condition []Condition) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		for _, item := range condition {
			wSql := ""
			if item.Symbol == "" || item.Symbol == "=" {
				wSql = fmt.Sprintf("%s = ? ", item.Name)
			} else {
				wSql = fmt.Sprintf("%s %s (?)", item.Name, item.Symbol)
			}
			db.Where(wSql, item.Value)
		}
		return db
	}
}
