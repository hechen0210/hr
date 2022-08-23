package model

type Route struct {
	Id        int
	Name      string
	Type      int
	Url       string
	CreatedAt int64
	UpdatedAt int64
}

func (Route) TableName() string {
	return "admin_route"
}
