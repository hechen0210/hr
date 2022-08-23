package model

type CompanyDemandLog struct {
	Id        int
	Cdid      int
	Status    int
	PreStatus int
	Content   string
	CreatedAt int64
}
