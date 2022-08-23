package model

type Group struct {
	Id        int
	Name      string
	Parent    int
	Privilege string
	CreatedAt int64
	UpdatedAt int64
}

func (Group) TableName() string {
	return "admin_group"
}