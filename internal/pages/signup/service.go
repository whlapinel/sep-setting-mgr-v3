package signup

import (
	domain "sep_setting_mgr/internal/domain/models"
	pages "sep_setting_mgr/internal/domain/pages"
)

type service struct {
	users domain.UserRepository
}

func NewService(users domain.UserRepository) pages.SignupService {
	return &service{users: users}
}

func (s service) CreateUser(email string, password string) (bool, error) {
	user, err := domain.NewUser(email, password)
	if err != nil {
		return false, err
	}
	err = s.users.Store(user)
	if err != nil {
		return false, err
	}
	return true, nil
}
