package users

import (
	"sep_setting_mgr/internal/domain/models"
	"sep_setting_mgr/internal/domain/pages"
)


type UsersService interface {
	ListUsers() ([]*models.User, error)
	ListRooms() ([]*models.Room, error)
	IsAdmin(userID int) bool
}


type service struct {
	users       models.UserRepository
	rooms       models.RoomRepository
	assignments models.AssignmentRepository
}

func NewService(users models.UserRepository, rooms models.RoomRepository, assignments models.AssignmentRepository) pages.AdminService {
	return &service{users, rooms, assignments}
}

func (s service) IsAdmin(userID int) bool {
	user, err := s.users.FindByID(userID)
	if err != nil {
		return false
	}
	return user.Admin
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

func (s service) FindRoomByID(id int) (*models.Room, error) {
	room, err := s.rooms.FindByID(id)
	if err != nil {
		return nil, err
	}
	return room, nil
}

func (s service) GetAllAssignments() (models.Assignments, error) {
	assignments, err := s.assignments.All()
	if err != nil {
		return nil, err
	}
	return assignments, nil
}
