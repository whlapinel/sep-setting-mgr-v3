package services

import (
	"errors"
	"log"
	"sep_setting_mgr/internal/domain/models"
	"sep_setting_mgr/internal/util"
	"time"
)

type AssignmentsService struct {
	asRepo     models.AssignmentRepository
	rooms      models.RoomRepository
	testEvents models.TestEventRepository
	classes    models.ClassRepository
	students   models.StudentRepository
}

func NewAssignmentsService(
	asRepo models.AssignmentRepository,
	rooms models.RoomRepository,
	testEvents models.TestEventRepository,
	classes models.ClassRepository,
	students models.StudentRepository) *AssignmentsService {

	return &AssignmentsService{
		asRepo,
		rooms,
		testEvents,
		classes,
		students,
	}
}

func (s *AssignmentsService) CreateAssignment(student *models.Student, testEvent *models.TestEvent) (*models.Assignment, error) {
	assignment := models.NewAssignment(student, nil, testEvent)
	s.asRepo.Store(assignment)
	log.Println("Assignment created")
	return assignment, nil
}

func (s *AssignmentsService) CreateAssignmentsForStudent(student *models.Student) error {
	// get test events for class and filter for future events
	log.Println("Service: Creating assignments for student")
	testEvents, err := s.testEvents.FindByClass(student.Class.ID)
	if err != nil {
		return err
	}
	var notAssignedErr error = nil
	for _, event := range testEvents {
		log.Println("creating assignment for event")
		event.Class = &student.Class
		assignment, err := s.CreateAssignment(student, event)
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

func (s *AssignmentsService) NullifyRoomID(roomID int) error {
	return s.asRepo.NullifyRoomID(roomID)
}

func (s *AssignmentsService) UpdateStudentAssignments(studentID int) error {
	student, err := s.students.FindByID(studentID)
	if err != nil {
		return err
	}
	err = s.DeleteAssignmentsForStudent(studentID)
	if err != nil {
		return err
	}
	err = s.CreateAssignmentsForStudent(student)
	if err != nil {
		return err
	}
	return nil
}

func (s *AssignmentsService) DeleteAssignmentsForStudent(studentID int) error {
	return s.asRepo.DeleteByStudentID(studentID)
}

func (s *AssignmentsService) DeleteAssignmentsForTestEvent(testEventID int) error {
	return s.asRepo.DeleteByEventID(testEventID)
}

func (s *AssignmentsService) CreateAssignmentsForTestEvent(testEventID int) error {
	testEvent, err := s.testEvents.FindByID(testEventID)
	if err != nil {
		return err
	}
	students, err := s.students.AllInClass(testEvent.Class.ID)
	if err != nil {
		return err
	}
	for _, student := range students {
		_, err := s.CreateAssignment(student, testEvent)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *AssignmentsService) GetAllAssignments() (models.Assignments, error) {
	return s.asRepo.All()
}

func (s *AssignmentsService) GetAssignments(classID int, start, end time.Time) (models.Assignments, error) {
	log.SetPrefix("Service: GetAssignments()")
	var assignments models.Assignments = make(models.Assignments, 0)
	testEvents, err := s.testEvents.FindByClass(classID)
	if err != nil {
		return nil, err
	}
	for _, event := range testEvents {
		if event.TestDate.After(start) && event.TestDate.Before(end) {
			eventAssignments, err := s.asRepo.FindByEventID(event.ID)
			if err != nil {
				return nil, err
			}
			assignments = append(assignments, eventAssignments...)
		}
	}
	return assignments, nil
}

func (s *AssignmentsService) GetAssignmentsByTeacherID(teacherID int) (models.Assignments, error) {
	return s.asRepo.GetByTeacherID(teacherID)
}

func (s *AssignmentsService) AutoAssign(assignments models.Assignments, date time.Time) error {
	// get all rooms
	rooms, err := s.rooms.All()
	if err != nil {
		log.Println("Error getting rooms")
		return nil
	}
	for _, a := range assignments {
		oneOnOne := a.Student.OneOnOne
		err := s.autoAssign(a, date, rooms, oneOnOne)
		if err != nil {
			log.Println("Error auto assigning")
			log.Println(err)
		}
	}
	return nil
}

func (s *AssignmentsService) autoAssign(a *models.Assignment, date time.Time, rooms models.Rooms, oneOnOne bool) error {
	var ErrNoRoomAvailable = errors.New("no room available")
	for _, r := range rooms {
		roomCount, err := s.asRepo.CountInRoomOnDate(r.ID, date)
		if err != nil {
			log.Println("Error counting assignments in room")
			return err
		}
		if oneOnOne {
			if roomCount == 0 {
				a.Room = r
				return nil
			}
		} else {
			if roomCount < r.MaxCapacity {
				a.Room = r
				return nil
			}
		}
	}
	return ErrNoRoomAvailable
}
