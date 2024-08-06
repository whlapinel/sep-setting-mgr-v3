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
		Repository[Room]
		DeleteAll
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

func NewRoom(name, number string, maxCapacity, priority int) (*Room, error) {
	return &Room{
		Name:        name,
		Number:      number,
		MaxCapacity: maxCapacity,
		Priority:    priority,
	}, nil
}

func (rooms Rooms) GetNextPriority() int {
	if len(rooms) == 0 {
		return 1
	}
	max := 0
	for _, room := range rooms {
		if room.Priority > max {
			max = room.Priority
		}
	}
	return max + 1
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
