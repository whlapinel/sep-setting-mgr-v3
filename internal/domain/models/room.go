package models

import "time"

type (
	Room struct {
		ID          int
		Name        string
		Number      string
		MaxCapacity int
		Priority    int // defines which order to display rooms and auto-fill with students
	}

	Rooms []*Room

	RoomRepository interface {
		Store(*Room) (int, error)
		Delete(roomID int) error
		All() (Rooms, error)
		FindByID(roomID int) (*Room, error)
		Update(room *Room) error
		FindByPriority(priority int) (*Room, error)
		GetRoomAssignments(room *Room, block int, date time.Time) (Assignments, error)
	}
)

var Unassigned = Room{
	ID:          -1,
	Name:        "Unassigned",
	Number:      "Unassigned",
	MaxCapacity: 1000000000,
	Priority:    -1,
}

func (rooms Rooms) SortByRoomPriority() Rooms {
	var sorted Rooms
	var unsorted Rooms
	for _, room := range rooms {
		if len(sorted) == 0 {
			sorted = append(sorted, room)
			continue
		}
		inserted := false
		for i, sortedRoom := range sorted {
			if room.Priority < sortedRoom.Priority {
				sorted = append(sorted[:i], append(Rooms{room}, sorted[i:]...)...)
				inserted = true
				break
			}
		}
		if !inserted {
			sorted = append(sorted, room)
		}
	}
	sorted = append(sorted, unsorted...)
	return sorted
}
