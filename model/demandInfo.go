package model

type DemandInfo struct {
	Id         int
	Cid        int
	Cdid       int
	Type       int
	Quantity   int
	SignTime   string
	Education  int
	Gender     int
	Age        string
	Expire     int
	ExpireTime string
	CreatedAt  int64
	UpdatedAt  int64
}

func (DemandInfo) TableName() string {
	return "demand_info"
}
