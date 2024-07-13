package dashboard

import (
	"log"
	"sep_setting_mgr/internal/domain/models"
	"sep_setting_mgr/internal/util"
	"time"
)

func (s service) CreateAssignmentsForStudent(student *models.Student) error {
	log.Println("Service: Creating assignments for student")
	classID := student.Class.ID
	class, err := s.classes.FindByID(classID)
	if err != nil {
		return err
	}
	testEvents, err := s.testEvents.FindByClass(classID)
	if err != nil {
		return err
	}
	log.Println("test events retrieved")
	// TODO: this should be a models.TestEvents method
	var futureEvents models.TestEvents
	for _, event := range testEvents {
		if event.TestDate.After(time.Now()) {
			futureEvents = append(futureEvents, event)
		}
	}
	log.Println("test events filtered for future events")
	var notAssignedErr error = nil
	for _, futureEvent := range futureEvents {
		log.Println("creating assignment for event")
		futureEvent.Class = class
		assignment, err := s.CreateAssignment(student, futureEvent)
		if err != nil {
			return err
		}
		if assignment == nil {
			log.Println("Assignment not created")
			notAssignedErr = util.ErrNotAssigned
		} else {
			log.Println("Assignment created")
		}
	}
	return notAssignedErr
}

func (s service) AddStudent(firstName string, lastName string, classID int, oneOnOne bool) (*models.Student, error) {
	log.SetPrefix("Student Service: ")
	log.Println("Service: Adding student to database")
	// find teacher id
	student, err := models.NewStudent(firstName, lastName, classID, oneOnOne)
	if err != nil {
		return nil, err
	}
	class, err := s.classes.FindByID(classID)
	if err != nil {
		return nil, err
	}
	student.Class = *class
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
	// TODO: this should be a models.TestEvents method
	var futureEvents models.TestEvents
	for _, event := range testEvents {
		if event.TestDate.After(time.Now()) {
			futureEvents = append(futureEvents, event)
		}
	}
	log.Println("test events filtered for future events")
	var notAssignedErr error = nil
	for _, futureEvent := range futureEvents {
		log.Println("creating assignment for event")
		futureEvent.Class = class
		assignment, err := s.CreateAssignment(student, futureEvent)
		if err != nil {
			return student, err
		}
		if assignment == nil {
			log.Println("Assignment not created")
			notAssignedErr = util.ErrNotAssigned
		} else {
			log.Println("Assignment created")
		}
	}
	return student, notAssignedErr
}

func (s service) UpdateStudent(firstName string, lastName string, oneOnOne bool, studentID int) (*models.Student, error) {
	log.Println("Service: Updating student in database")
	student, err := s.students.FindByID(studentID)
	if err != nil {
		return nil, err
	}
	oneOnOneChanged := student.OneOnOne != oneOnOne
	// delete all assignments for student
	student.FirstName = firstName
	student.LastName = lastName
	student.OneOnOne = oneOnOne
	err = s.students.Update(student)
	if err != nil {
		log.Println("Failed to update student: ", err)
		return nil, err
	}
	var notAssignedErr error = nil
	if oneOnOneChanged {
		err := s.DeleteAssignments(studentID)
		if err != nil {
			return nil, err
		}
		err = s.CreateAssignmentsForStudent(student)
		if err != nil {
			if err == util.ErrNotAssigned {
				notAssignedErr = err
			} else {
				return nil, err
			}
		}
	}
	return student, notAssignedErr
}

func (s service) FindStudentByID(studentID int) (*models.Student, error) {
	student, err := s.students.FindByID(studentID)
	if err != nil {
		return nil, err
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
