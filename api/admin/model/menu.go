package model

type Menu struct {
	Id        int
	Name      string
	Icon      string
	Parent    int
	Level     int
	Sort      int
	Show      int
	ShowType  int
	Url       string
	Api       string
	CreatedAt int64
	UpdatedAt int64
}

func (Menu) TableName() string {
	return "admin_menu"
}