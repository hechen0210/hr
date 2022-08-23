package model

import "github.com/shopspring/decimal"

type CompanySalary struct {
	Id             int
	Cid            int
	Date           string
	StaffNum       string
	TotalSalary    decimal.Decimal
	TotalDeduction decimal.Decimal
	RealSalary     decimal.Decimal
	Status         int
	CreatedAt      int64
	UpdatedAt      int64
}

func (CompanySalary) TableName() string {
	return "company_salary"
}