package model

import "github.com/shopspring/decimal"

type CompanyDemand struct {
	Id        int
	Cid       int
	Name      string
	Unit      decimal.Decimal
	Desc      string
	Flag      int
	Published int
	CreatedAt int64
	UpdatedAt int64
}

func (CompanyDemand) TableName() string {
	return "company_demand"
}
