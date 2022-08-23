package model

import "github.com/shopspring/decimal"

type UserSalary struct {
	Id             int
	Uid            int
	Cid            int
	Csid           int
	TotalSalry     decimal.Decimal
	TotalDeduction decimal.Decimal
	RealSalary     decimal.Decimal
	Status         int
	CreatedAt      int64
	UpdatedAt      int64
}
