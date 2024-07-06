package dashboard

import (
	"log"
	domain "sep_setting_mgr/internal/domain/models"
)

func (s service) AddStudent(firstName string, lastName string, classID int, oneOnOne bool) (*domain.Student, error) {
	log.Println("Service: Adding student to database")
	// find teacher id
	student, err := domain.NewStudent(firstName, lastName, classID, oneOnOne)
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

func (s service) DeleteStudent(studentID int) error {
	log.Println("Service: Deleting student from database")
	err := s.students.Delete(studentID)
	if err != nil {
		return err
	}
	return nil
}

func (s service) ListStudents(classID int) ([]*domain.Student, error) {
	students, err := s.students.All(classID)
	if err != nil {
		return nil, err
	}
	return students, nil
}
