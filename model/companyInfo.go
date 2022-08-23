package model

type CompanyInfo struct {
	Id   int
	Cid  int
	Logo string
	Desc string
}

func (CompanyInfo) TableName() string {
	return "company_info"
}