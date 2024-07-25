package applications

import (
	"sep_setting_mgr/internal/domain/models"
)

type ApplicationsService interface {
	ApplyForRole(userID int, role string) error
}

type service struct {
	applications models.ApplicationRepository
	users        models.UserRepository
}

func NewService(applications models.ApplicationRepository, users models.UserRepository) ApplicationsService {
	return &service{applications, users}
}

func (s service) ApplyForRole(userId int, roleString string) error {
	role, err := models.GetRole(roleString)
	if err != nil {
		return err
	}
	user, err := s.users.FindByID(userId)
	if err != nil {
		return err
	}
	application, err := models.NewApplication(userId, user.FirstName, user.LastName, user.Email, role)
	if err != nil {
		return err
	}
	return s.applications.Store(application)
}
