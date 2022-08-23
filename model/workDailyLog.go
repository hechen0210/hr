package model

type WorkDailyLog struct {
	Id          int
	Csid        int
	Cid         int
	Csno        string
	Uid         int
	WorkType    int
	SignTime    int
	WorkTime    int
	Quantity    int
	InvalidTime int
	ExtraTime   int
	CreatedAt   int64
	UpdatedAt   int64
}
