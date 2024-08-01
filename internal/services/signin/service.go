package signin

import (
	"errors"
	"sep_setting_mgr/internal/domain/models"
	domain "sep_setting_mgr/internal/domain/models"
)

type SigninService interface {
	GetUser(email string) (*models.User, error)
}

type service struct {
	users domain.UserRepository
}

func NewService(users domain.UserRepository) SigninService {
	return &service{users: users}
}

func (s service) GetUser(email string) (*models.User, error) {
	user, err := s.users.Find(email)
	if err != nil {
		return nil, err
	}
	if user.Email == "" {
		return nil, errors.New("user not found")
	}
	return user, nil
}
