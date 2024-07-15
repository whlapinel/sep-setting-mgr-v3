package admin

import (
	"sep_setting_mgr/internal/domain/models"
)

type AdminService interface {
	IsAdmin(userID int) bool
}

type service struct {
	users models.UserRepository
}

func NewService(users models.UserRepository) AdminService {
	return &service{users}
}

func (s service) IsAdmin(userID int) bool {
	user, err := s.users.FindByID(userID)
	if err != nil {
		return false
	}
	return user.Admin
}
