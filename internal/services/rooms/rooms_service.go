package rooms

import "sep_setting_mgr/internal/domain/models"

type RoomsService interface {
	AddRoom(*models.Room) (id int, err error)
	ListRooms() ([]*models.Room, error)
	FindRoomByID(id int) (*models.Room, error)
}

type service struct {
	rooms models.RoomRepository
}

func NewService(rooms models.RoomRepository) RoomsService {
	return &service{rooms}
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
