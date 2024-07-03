package signin

import (
	"sep_setting_mgr/internal/domain"
)

type service struct {
	users domain.UserRepository
}

func NewService(users domain.UserRepository) domain.SigninService {
	return &service{users: users}
}

func (s service) VerifyCredentials(email string, password string) (bool, error) {
	user, err := s.users.Find(email)
	if err != nil {
		return false, err
	}
	return user.Password == password, nil
}

func (s service) GetUserID(email string) int {
	user, _ := s.users.Find(email)
	return user.ID
}
