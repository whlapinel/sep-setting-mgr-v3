package assignments

import (
	"log"
	"sep_setting_mgr/internal/domain/models"
	"time"
)

type AssignmentsService interface {
	ListAll() (models.Assignments, error)
	GetAssignments(teacherID int, start, end time.Time) (models.Assignments, error)
	CreateAssignment(student *models.Student, testEvent *models.TestEvent) (*models.Assignment, error)
	DeleteByStudentID(studentID int) error
}

type service struct {
	assignments models.AssignmentRepository
	rooms       models.RoomRepository
	testEvents  models.TestEventRepository
}

func NewService(assignments models.AssignmentRepository, rooms models.RoomRepository, testEvents models.TestEventRepository) AssignmentsService {
	return &service{
		assignments,
		rooms,
		testEvents,
	}
}

func (s service) CreateAssignment(student *models.Student, testEvent *models.TestEvent) (*models.Assignment, error) {
	var assignment *models.Assignment

	return assignment, nil
}

func (s service) ListAll() (models.Assignments, error) {
	assignments, err := s.assignments.All()
	if err != nil {
		return nil, err
	}
	return assignments, nil
}

func (s service) GetAssignments(classID int, start, end time.Time) (models.Assignments, error) {
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

func (s service) DeleteByStudentID(studentID int) error {
	return s.assignments.DeleteByStudentID(studentID)
}
