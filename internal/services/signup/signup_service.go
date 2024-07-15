package signup

import "sep_setting_mgr/internal/domain/models"

type SignupService interface {
	CreateUser(username string, password string) (bool, error)
}

type service struct {
	users models.UserRepository
}

func NewService(users models.UserRepository) SignupService {
	return &service{users: users}
}

func (s service) CreateUser(email string, password string) (bool, error) {
	user, err := models.NewUser(email, password)
	if err != nil {
		return false, err
	}
	err = s.users.Store(user)
	if err != nil {
		return false, err
	}
	return true, nil
}
