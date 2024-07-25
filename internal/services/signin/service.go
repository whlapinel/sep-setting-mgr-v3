package signin

import (
	"errors"
	domain "sep_setting_mgr/internal/domain/models"
)

type SigninService interface {
	GetUserID(email string) (int, error)
}

type service struct {
	users domain.UserRepository
}

func NewService(users domain.UserRepository) SigninService {
	return &service{users: users}
}

func (s service) GetUserID(email string) (int, error) {
	user, err := s.users.Find(email)
	if err != nil {
		return 0, err
	}
	if user.Email == "" {
		return 0, errors.New("user not found")
	}
	return user.ID, nil
}
