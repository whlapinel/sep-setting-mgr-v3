package admin

import (
	"sep_setting_mgr/internal/domain/models"
)

type AdminService interface {
	IsAdmin(userID int) bool
	ListAllApplications() (models.Applications, error)
}

type service struct {
	users        models.UserRepository
	applications models.ApplicationRepository
}

func NewService(users models.UserRepository, applications models.ApplicationRepository) AdminService {
	return &service{users, applications}
}

func (s service) IsAdmin(userID int) bool {
	user, err := s.users.FindByID(userID)
	if err != nil {
		return false
	}
	return user.Admin
}

func (s service) ListAllApplications() (models.Applications, error) {
	return s.applications.All()
}
