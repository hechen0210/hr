package model

type Admin struct {
	Id          int
	Account     string
	Name        string
	Password    string
	Phone       string
	Email       string
	Group       int
	Status      int
	Uuid        string
	CreatedAt   int64
	UpdatedAt   int64
	LastLoginAt int64
	LastLoginIp int
}

func (Admin) TableName() string {
	return "admin"
}
