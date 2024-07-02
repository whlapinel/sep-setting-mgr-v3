package dashboard

import (
	"sep_setting_mgr/internal/domain"
)

type (
	Service interface {
		// List returns a copy of the todos list
		List(teacherID int) ([]*domain.Class, error)
		AddClass(name string, block int, teacherID int) (*domain.Class, error)
		DeleteClass(classID int) error
		AddStudent(firstName string, lastName string, classID int) (*domain.Student, error)
		DeleteStudent(studentID int) error
		ListStudents(classID int) ([]*domain.Student, error)
		FindClassByID(classID int) (*domain.Class, error)
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
