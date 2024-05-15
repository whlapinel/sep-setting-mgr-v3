package dashboard

import (
	"log"
	"sep_setting_mgr/internal/domain"
)

type (
	Service interface {
		// List returns a copy of the todos list
		List() ([]*domain.Class, error)
		AddClass(name string, block int, teacherID int) (*domain.Class, error)
		FindClassByID(classID string) (*domain.Class, error)
	}

	service struct {
		classes domain.ClassRepository
		users   domain.UserRepository
	}
)

func NewService(classes domain.ClassRepository, users domain.UserRepository) Service {
	return &service{
		classes: classes,
		users:   users,
	}
}

func (s service) List() ([]*domain.Class, error) {
	return s.classes.All(), nil
}

func (s service) AddClass(name string, block int, teacherID int) (*domain.Class, error) {
	log.Println("Service: Adding class to database")
	// find teacher id
	class := domain.NewClass(name, block, teacherID)
	err := s.classes.Store(class)
	if err != nil {
		return nil, err
	}
	log.Println("Teacher email:", class.Teacher.Email)
	return class, nil
}

func (s service) FindClassByID(classID string) (*domain.Class, error) {
	class, err := s.classes.FindByID(classID)
	if err != nil {
		return nil, err
	}
	return class, nil
}
