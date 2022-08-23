package request

import "hr/repository"

type CompanySalaryRequest struct{

}

func NewCompanySalaryRequest() CompanySalaryRequest {
	return CompanySalaryRequest{}
}

func (sr *CompanySalaryRequest) Search(data map[string]string) []repository.Condition {
	return []repository.Condition{}
}