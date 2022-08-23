package model

type UserProfile struct {
	Id        int
	Uid       int
	Name      string
	Gender    int
	Birthday  string
	Education string
	School    string
	Major     string
	CreatedAt int64
	UpdatedAt int64
}

func (UserProfile) TableName() string {
	return "user_profile"
}
