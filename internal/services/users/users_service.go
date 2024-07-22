package users

import (
	"sep_setting_mgr/internal/domain/models"
)

type UsersService interface {
	ListUsers() ([]*models.User, error)
	FindUserByID(id int) (*models.User, error)
	UpdateUser(user *models.User) error
	IsAdmin(userID int) bool
}

type service struct {
	users models.UserRepository
}

func NewService(users models.UserRepository) UsersService {
	return &service{users}
}

func (s service) IsAdmin(userID int) bool {
	user, err := s.users.FindByID(userID)
	if err != nil {
		return false
	}
	return user.Admin
}

func (s service) ListUsers() ([]*models.User, error) {
	users, err := s.users.All()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s service) FindUserByID(id int) (*models.User, error) {
	user, err := s.users.FindByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s service) UpdateUser(user *models.User) error {
	err := s.users.Update(user)
	if err != nil {
		return err
	}
	return nil
}
