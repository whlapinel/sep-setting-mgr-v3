package dashboard

import (
	"log"
	"sep_setting_mgr/internal/domain"
)

func (s service) DeleteClass(classID int) error {
	log.Println("Service: Deleting class from database")
	err := s.classes.Delete(classID)
	if err != nil {
		return err
	}
	return nil
}

func (s service) AddClass(name string, block int, teacherID int) (*domain.Class, error) {
	log.Println("Service: Adding class to database")
	// find teacher id
	class := domain.NewClass(name, block, teacherID)
	id, err := s.classes.Store(class)
	if err != nil {
		return nil, err
	}
	class.ID = id
	return class, nil
}

func (s service) FindClassByID(classID int) (*domain.Class, error) {
	class, err := s.classes.FindByID(classID)
	if err != nil {
		return nil, err
	}
	return class, nil
}

func (s service) List(teacherID int) ([]*domain.Class, error) {
	classes, err := s.classes.All(teacherID)
	if err != nil {
		return nil, err
	}
	for _, class := range classes {
		students, err := s.students.All(class.ID)
		if err != nil {
			return nil, err
		}
		class.Students = students
	}
	return classes, nil
}
