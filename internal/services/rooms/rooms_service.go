package rooms

import (
	"sep_setting_mgr/internal/domain/models"
	"sep_setting_mgr/internal/domain/services"
)

type RoomsService interface {
	AddRoom(*models.Room) (id int, err error)
	DeleteRoom(id int) error
	UpdateRoom(room *models.Room) error
	ListRooms() (models.Rooms, error)
	FindRoomByID(id int) (*models.Room, error)
}

type service struct {
	rooms     models.RoomRepository
	asService services.AssignmentsService
}

func NewService(rooms models.RoomRepository, asService services.AssignmentsService) RoomsService {
	return &service{rooms, asService}
}

func (s service) ListRooms() (models.Rooms, error) {
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

func (s service) DeleteRoom(id int) error {
	err := s.asService.NullifyRoomID(id)
	if err != nil {
		return err
	}
	err = s.rooms.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (s service) UpdateRoom(room *models.Room) error {
	err := s.rooms.Update(room)
	if err != nil {
		return err
	}
	return nil
}

func (s service) FindRoomByID(id int) (*models.Room, error) {
	room, err := s.rooms.FindByID(id)
	if err != nil {
		return nil, err
	}
	return room, nil
}
