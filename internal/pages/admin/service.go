package admin

import (
	"sep_setting_mgr/internal/domain/models"
	"sep_setting_mgr/internal/domain/pages"
)

type service struct {
	users       models.UserRepository
	rooms       models.RoomRepository
	assignments models.AssignmentRepository
}

func NewService(users models.UserRepository, rooms models.RoomRepository, assignments models.AssignmentRepository) pages.AdminService {
	return &service{users, rooms, assignments}
}

func (s service) ListUsers() ([]*models.User, error) {
	users, err := s.users.All()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s service) ListRooms() ([]*models.Room, error) {
	rooms, err := s.rooms.All()
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

func (s service) AddRoom(room *models.Room) (id int, err error) {
	id, err = s.rooms.Store(room)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s service) GetAllAssignments() (models.Assignments, error) {
	assignments, err := s.assignments.All()
	if err != nil {
		return nil, err
	}
	return assignments, nil
}
