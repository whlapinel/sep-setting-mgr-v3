package signup

import "sep_setting_mgr/internal/domain"

type service struct {
	users domain.UserRepository
}

func NewService(users domain.UserRepository) domain.SignupService {
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
