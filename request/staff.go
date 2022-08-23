package request

import "hr/repository"

type StaffRequest struct {
}

func NewStaffRequest() StaffRequest {
	return StaffRequest{}
}

func (sr *StaffRequest) Search(query map[string]string) []repository.Condition {
	return []repository.Condition{}
}
