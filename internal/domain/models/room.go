package models

import "time"

type (
	Room struct {
		ID          int
		Name        string
		Number      string
		MaxCapacity int
		Priority    int // defines which order to fill rooms with students
	}

	RoomRepository interface {
		Store(*Room) (int, error)
		Delete(roomID int) error
		All() ([]*Room, error)
		FindByID(roomID int) (*Room, error)
		Update(room *Room) error
		FindByPriority(priority int) (*Room, error)
		GetRoomAssignments(room *Room, block int, date time.Time) (Assignments, error)
	}
)
