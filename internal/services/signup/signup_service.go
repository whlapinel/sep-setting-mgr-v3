package signup

import "sep_setting_mgr/internal/domain/models"

type SignupService interface {
	IsDuplicate(email string) (bool, error)
	CreateUser(first, last, email, picture string) (*models.User, error)
}

type service struct {
	users models.UserRepository
}

func NewService(users models.UserRepository) SignupService {
	return &service{users: users}
}

func (s service) CreateUser(first, last, email, picture string) (*models.User, error) {
	user, err := models.NewUser(first, last, email, picture)
	if err != nil {
		return nil, err
	}
	err = s.users.Store(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s service) IsDuplicate(email string) (bool, error) {
	user, err := s.users.Find(email)
	if err != nil {
		return false, err
	}
	if user.Email == email {
		return true, nil
	}
	return false, nil
}
