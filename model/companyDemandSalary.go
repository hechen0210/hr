package model

import "github.com/shopspring/decimal"

type CompanyDemandSalary struct {
	Id             int
	Cid            int
	Cdid           int
	Date           string
	StaffNum       int
	TotalSalary    decimal.Decimal
	TotalDeduction decimal.Decimal
	RealSalary     decimal.Decimal
	Status         int
	CreatedAt      int64
	UpdatedAt      int64
}
