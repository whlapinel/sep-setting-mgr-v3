package applications

import (
	"errors"
	"log"
	"sep_setting_mgr/internal/domain/models"
)

type ApplicationsService interface {
	ApplyForRole(userID int, role models.Role) error
	HasRole(userID int, role models.Role) (bool, error)
	HasApplied(userID int, role models.Role) (bool, error)
	AdjudicateApplication(appID int, action models.Action) error
	All() (models.Applications, error)
}

type service struct {
	applications models.ApplicationRepository
	users        models.UserRepository
}

func NewService(applications models.ApplicationRepository, users models.UserRepository) ApplicationsService {
	return &service{applications, users}
}

func (s service) ApplyForRole(userId int, role models.Role) error {
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

func (s service) HasApplied(userID int, role models.Role) (bool, error) {
	log.SetPrefix("Application Service: HasApplied() ")
	log.Println("Role: ", role)
	applications, err := s.applications.FindByUserID(userID)
	if err != nil {
		log.Println(err)
		return false, err
	}
	log.Printf("Iterating through %v Applications: ", len(applications))
	for _, application := range applications {
		log.Println("Application Role: ", application.Role)
		if application.Role == role {
			return true, nil
		}
	}
	return false, nil
}

func (s service) HasRole(userID int, role models.Role) (bool, error) {
	user, err := s.users.FindByID(userID)
	if err != nil {
		return false, nil
	}
	if role == models.AdminRole && user.Admin {
		return true, nil
	}
	if role == models.TeacherRole && user.Teacher {
		return true, nil
	}
	return false, nil
}

var ErrApplicationNotFound = errors.New("application not found")
var ErrSearchError = errors.New("error searching for application")

func (s service) AdjudicateApplication(appID int, action models.Action) error {
	if action != models.Approve && action != models.Deny {
		return models.ErrInvalidAction
	}
	app, err := s.applications.FindByID(appID)
	if err != nil {
		return ErrSearchError
	}
	if app == nil {
		return ErrApplicationNotFound
	}
	if action == models.Approve {
		user, err := s.users.FindByID(app.UserID)
		if err != nil {
			return err
		}
		if app.Role == models.AdminRole {
			user.Admin = true
		}
		if app.Role == models.TeacherRole {
			user.Teacher = true
		}
		err = s.users.Update(user)
		if err != nil {
			return err
		}
		return s.applications.Delete(app.ID)
	}
	return s.applications.Delete(app.ID)
}

func (s service) All() (models.Applications, error) {
	return s.applications.All()
}
