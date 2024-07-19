package assignments

import (
	"log"
	"sep_setting_mgr/internal/domain/models"
)

type AssignmentsService interface {
	ListAll() (models.Assignments, error)
	GetByAssignmentID(id int) (*models.Assignment, error)
	UpdateRoom(assignmentID, roomID int) (*models.Room, error)
	GetByTeacherID(teacherID int) (models.Assignments, error)
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

func (s service) ListAll() (models.Assignments, error) {
	assignments, err := s.assignments.All()
	if err != nil {
		return nil, err
	}
	return assignments, nil
}

func (s service) GetByAssignmentID(id int) (*models.Assignment, error) {
	return s.assignments.GetByAssignmentID(id)
}

func (s service) UpdateRoom(assignmentID, roomID int) (*models.Room, error) {
	var room *models.Room
	if roomID < 0 {
		log.Printf("Nullifying room for assignment %d", assignmentID)
		err := s.assignments.NullifyRoomIDByAssignmentID(assignmentID)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
	err := s.assignments.UpdateRoom(assignmentID, roomID)
	if err != nil {
		return nil, err
	}
	room, err = s.rooms.FindByID(roomID)
	if err != nil {
		return nil, err
	}
	return room, nil
}

func (s service) GetByTeacherID(teacherID int) (models.Assignments, error) {
	return s.assignments.GetByTeacherID(teacherID)
}
