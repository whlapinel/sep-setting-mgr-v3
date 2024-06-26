package dashboard

import (
	"log"
	"sep_setting_mgr/internal/domain"
)

type (
	Service interface {
		// List returns a copy of the todos list
		List(teacherID int) ([]*domain.Class, error)
		AddClass(name string, block int, teacherID int) (*domain.Class, error)
		DeleteClass(classID int) error
		AddStudent(firstName string, lastName string, classID int) (*domain.Student, error)
		FindClassByID(classID string) (*domain.Class, error)
	}

	service struct {
		classes  domain.ClassRepository
		users    domain.UserRepository
		students domain.StudentRepository
	}
)

func NewService(classes domain.ClassRepository, users domain.UserRepository, students domain.StudentRepository) Service {
	return &service{
		classes:  classes,
		users:    users,
		students: students,
	}
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

func (s service) AddStudent(firstName string, lastName string, classID int) (*domain.Student, error) {
	log.Println("Service: Adding student to database")
	// find teacher id
	student, err := domain.NewStudent(firstName, lastName, classID)
	if err != nil {
		return nil, err
	}
	id, err := s.students.Store(student)
	if err != nil {
		return nil, err
	}
	student.ID = id
	return student, nil
}

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

func (s service) FindClassByID(classID string) (*domain.Class, error) {
	class, err := s.classes.FindByID(classID)
	if err != nil {
		return nil, err
	}
	return class, nil
}
