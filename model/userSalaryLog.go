package model

import "github.com/shopspring/decimal"

type UserSalaryLog struct {
	Id                 int
	Uid                int
	Cid                int
	Csid               int
	Csno               string
	Date               string
	BaseSalary         decimal.Decimal // 基本工资
	OvertimeSalary     decimal.Decimal // 加班工资
	Bonus              decimal.Decimal // 奖金
	PercentageSalary   decimal.Decimal // 提成
	ExtraSalary        decimal.Decimal // 其它工资
	ExtraSalaryDesc    string          // 其它工资描述
	TaxDeduction       decimal.Decimal // 扣税
	LateDeduction      decimal.Decimal // 迟到扣除
	LeaveDeduction     decimal.Decimal // 早退扣除
	SocialDeduction    decimal.Decimal // 社保
	FundsDeduction     decimal.Decimal // 公积金
	ExtraDeduction     decimal.Decimal // 其它扣除
	ExtraDeductionDesc string          // 其它扣除说明
	CreatedAt          int64
	UpdatedAt          int64
}
