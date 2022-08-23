package model

type CompanyStaff struct {
	Id        int
	Cid       int
	No        string
	Uid       int
	Status    int
	CreatedAt int64
	UpdatedAt int64
}

func (CompanyStaff) TableName() string {
	return "company_staff"
}