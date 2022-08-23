package repository

import "hr/model"

type UserProfile struct {
	UPModel model.UserProfile
}

func NewUserProfileRepo() *UserProfile {
	return &UserProfile{
		UPModel: model.UserProfile{},
	}
}

func (up *UserProfile) GetUserProfileById(id int) {

}

func (up *UserProfile) UpdateUserProfile() {

}
