package services

import (
	"log"
	"sep_setting_mgr/internal/domain/models"
	"sep_setting_mgr/internal/util"
	"time"
)

type AssignmentsService struct {
	assignments models.AssignmentRepository
	rooms       models.RoomRepository
	testEvents  models.TestEventRepository
	classes     models.ClassRepository
	students    models.StudentRepository
}

func NewAssignmentsService(
	assignments models.AssignmentRepository,
	rooms models.RoomRepository,
	testEvents models.TestEventRepository,
	classes models.ClassRepository,
	students models.StudentRepository) *AssignmentsService {

	return &AssignmentsService{
		assignments,
		rooms,
		testEvents,
		classes,
		students,
	}
}

func (s *AssignmentsService) CreateAssignment(student *models.Student, testEvent *models.TestEvent) (*models.Assignment, error) {
	var assignment *models.Assignment
	priority := 1
	foundOpenRoom := false
	var max int
	for !foundOpenRoom {
		room, err := s.rooms.FindByPriority(priority)
		if err != nil {
			return nil, err
		}
		if room == nil {
			return nil, nil
		}
		max = room.MaxCapacity
		if student.OneOnOne {
			max = 1
		}
		roomAssignments, err := s.rooms.GetRoomAssignments(room, testEvent.Class.Block, *testEvent.TestDate)
		if err != nil {
			return nil, err
		}
		if len(roomAssignments) < max {
			foundOpenRoom = true
			assignment = models.NewAssignment(student, room, testEvent)
			err := s.assignments.Store(assignment)
			if err != nil {
				return nil, err
			}
		}
		priority++
	}
	return assignment, nil
}

func (s *AssignmentsService) CreateAssignmentsForStudent(student *models.Student) error {
	// get test events for class and filter for future events
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
	return s.assignments.DeleteByStudentID(studentID)
}

func (s *AssignmentsService) CreateAssignmentsForTestEvent(testEventID int) error {
	testEvent, err := s.testEvents.FindByID(testEventID)
	if err != nil {
		return err
	}
	students, err := s.students.All(testEvent.Class.ID)
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
	return s.assignments.All()
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
			eventAssignments, err := s.assignments.GetByEventID(event.ID)
			if err != nil {
				return nil, err
			}
			assignments = append(assignments, eventAssignments...)
		}
	}
	return assignments, nil
}
