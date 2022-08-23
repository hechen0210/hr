package model

type Company struct {
	Id        int
	Name      string
	Account   string
	Password  string
	Status    int
	Uuid      string
	CreatedAt int64
	UpdatedAt int64
}

func (Company) TableName() string {
	return "company"
}
