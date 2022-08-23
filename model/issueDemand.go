package model

import "github.com/shopspring/decimal"

type IssueDemand struct {
	Id        int
	Cid       int
	Cdid      int
	Name      string
	Desc      string
	Unit      decimal.Decimal
	Status    int
	CreatedAt int64
	UpdatedAt int64
}
