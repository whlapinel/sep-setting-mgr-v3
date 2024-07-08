package dashboard

import (
	"log"
	"sep_setting_mgr/internal/domain/models"
	"time"
)

func (s service) AddStudent(firstName string, lastName string, classID int, oneOnOne bool) (*models.Student, error) {
	log.SetPrefix("Student Service: ")
	log.Println("Service: Adding student to database")
	// find teacher id
	student, err := models.NewStudent(firstName, lastName, classID, oneOnOne)
	if err != nil {
		return nil, err
	}
	log.Println("new student created")
	id, err := s.students.Store(student)
	if err != nil {
		return nil, err
	}
	log.Println("new student stored")
	student.ID = id
	testEvents, err := s.testEvents.FindByClass(classID)
	if err != nil {
		return nil, err
	}
	log.Println("test events retrieved")
	var futureEvents models.TestEvents
	for _, event := range testEvents {
		if event.TestDate.After(time.Now()) {
			futureEvents = append(futureEvents, event)
		}
	}
	log.Println("test events filtered for future events")
	for _, futureEvent := range futureEvents {
		log.Println("creating assignment for event")
		_, err := s.CreateAssignment(student, futureEvent)
		if err != nil {
			return nil, err
		}
	}
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

func (s service) ListStudents(classID int) ([]*models.Student, error) {
	students, err := s.students.All(classID)
	if err != nil {
		return nil, err
	}
	return students, nil
}
