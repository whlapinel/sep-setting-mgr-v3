package demodata

import "sep_setting_mgr/internal/domain/models"

func (ds *demoDataService) createDemoRooms() ([]*models.Room, error) {
	roomNumbers := []string{
		"Room 101",
		"Room 102",
		"Room 103",
		"Room 104",
		"Room 105",
		"Room 106",
		"Room 107",
		"Room 108",
		"Room 109",
		"Room 110",
	}
	for i, number := range roomNumbers {
		room, err := models.NewRoom("", number, 12, i+1)
		if err != nil {
			return nil, err
		}
		err = ds.roomsRepo.Store(room)
		if err != nil {
			return nil, err
		}
	}
	rooms, err := ds.roomsRepo.All()
	if err != nil {
		return nil, err
	}
	return rooms, nil
}
