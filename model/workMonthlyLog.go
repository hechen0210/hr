package model

type WorkMonthlyLog struct {
	Id          int
	Csid        int
	Cid         int
	Csno        string
	Uid         int
	WorkType    int
	WorkTime    int
	Quantity    int
	InvalidTime int
	ExtraTime   int
	CreatedAt   int64
	UpdatedAt   int64
}
