package assignments

import (
	"sep_setting_mgr/internal/domain/models"
)

type AssignmentsService interface {
	All() (models.Assignments, error)
	FindByID(id int) (*models.Assignment, error)
	UpdateRoom(assignmentID, newRoomID int) error
	FindByTeacherID(teacherID int) (models.Assignments, error)
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

func (s service) All() (models.Assignments, error) {
	assignments, err := s.assignments.All()
	if err != nil {
		return nil, err
	}
	return assignments, nil
}

func (s service) FindByID(id int) (*models.Assignment, error) {
	return s.assignments.FindByID(id)
}

func (s service) UpdateRoom(assignmentID, newRoomID int) error {
	var room *models.Room
	if newRoomID < 0 {
		room = &models.Unassigned
	} else {
		storedRoom, err := s.rooms.FindByID(newRoomID)
		if err != nil {
			return err
		}
		room = storedRoom
	}
	assignment, err := s.assignments.FindByID(assignmentID)
	if err != nil {
		return err
	}
	assignment.Room = room
	err = s.assignments.Update(assignment)
	if err != nil {
		return err
	}
	return nil
}

func (s service) FindByTeacherID(teacherID int) (models.Assignments, error) {
	return s.assignments.GetByTeacherID(teacherID)
}
