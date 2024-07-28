package rooms

import (
	"errors"
	"sep_setting_mgr/internal/domain/models"
	"sep_setting_mgr/internal/domain/services"
)

type RoomsService interface {
	AddRoom(*models.Room) (id int, err error)
	DeleteRoom(id int) error
	UpdateRoom(room *models.Room) error
	ListRooms() (models.Rooms, error)
	FindRoomByID(id int) (*models.Room, error)
	PromoteRoom(id int) error
}

type service struct {
	rooms     models.RoomRepository
	asService services.AssignmentsService
}

func NewService(rooms models.RoomRepository, asService services.AssignmentsService) RoomsService {
	return &service{rooms, asService}
}

// should swap priority with next higher priority room
func (s service) PromoteRoom(id int) error {
	room, err := s.rooms.FindByID(id)
	if err != nil {
		return err
	}
	if room.Priority < 2 {
		return errors.New("room cannot be promoted due to current priority being less than 2")
	}
	nextRoom, err := s.rooms.FindByPriority(room.Priority - 1)
	if err != nil {
		return err
	}
	if nextRoom == nil || nextRoom.ID == 0 {
		room.Priority--
		s.rooms.Update(room)
		return nil
	}
	nextRoom.Priority++
	err = s.rooms.Update(nextRoom)
	if err != nil {
		return err
	}
	room.Priority--
	err = s.rooms.Update(room)
	if err != nil {
		return err
	}
	return nil
}

func (s service) ListRooms() (models.Rooms, error) {
	rooms, err := s.rooms.All()
	if err != nil {
		return nil, err
	}
	return rooms.SortByRoomPriority(), nil
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
