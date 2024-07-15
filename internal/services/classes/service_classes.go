package classes

import (
	"log"
	"sep_setting_mgr/internal/domain/models"
)

type ClassesService interface {
	List(teacherID int) ([]*models.Class, error)
	AddClass(name string, block int, teacherID int) (*models.Class, error)
	DeleteClass(classID int) error
	UpdateClass(classID int, name string) (*models.Class, error)
	FindClassByID(classID int) (*models.Class, error)
}

type service struct {
	classes  models.ClassRepository
	students models.StudentRepository
}

func NewService(classes models.ClassRepository, students models.StudentRepository) ClassesService {
	return &service{
		classes:  classes,
		students: students,
	}
}

func (s service) DeleteClass(classID int) error {
	log.Println("Service: Deleting class from database")
	err := s.classes.Delete(classID)
	if err != nil {
		return err
	}
	return nil
}

func (s service) UpdateClass(classID int, name string) (*models.Class, error) {
	log.SetPrefix("Class Service: ")
	log.Println("Updating class in database")
	log.Println("Class ID: ", classID)
	class, err := s.classes.FindByID(classID)
	if err != nil {
		return nil, err
	}
	class.Name = name
	err = s.classes.Update(class)
	if err != nil {
		return nil, err
	}
	return class, nil
}

func (s service) AddClass(name string, block int, teacherID int) (*models.Class, error) {
	log.Println("Service: Adding class to database")
	class := models.NewClass(name, block, teacherID)
	id, err := s.classes.Store(class)
	if err != nil {
		return nil, err
	}
	class.ID = id
	return class, nil
}

func (s service) FindClassByID(classID int) (*models.Class, error) {
	class, err := s.classes.FindByID(classID)
	if err != nil {
		return nil, err
	}
	return class, nil
}

func (s service) List(teacherID int) ([]*models.Class, error) {
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
