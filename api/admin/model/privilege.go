package model

type Privilege struct {
	Id        int
	Admin     int
	Privilege string
	CreatedAt int64
	UpdatedAt int64
}

func (Privilege) TableName() string {
	return "admin_privilege"
}
