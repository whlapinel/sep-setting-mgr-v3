package models

import (
	"log"
	"time"
)

// Assignment represents a student's assignment to a room.
type Assignment struct {
	ID        int
	Student   *Student
	Room      *Room
	TestEvent *TestEvent
	Block     int
}

type AssignmentRepository interface {
	Repository[Assignment]
	DeleteAll
	FindByID(id int) (*Assignment, error)
	FindByEventID(eventID int) (Assignments, error)
	DeleteByStudentID(studentID int) error
	DeleteByEventID(eventID int) error
	NullifyRoomID(roomID int) error
	GetByTeacherID(teacherID int) (Assignments, error)
}

type Assignments []*Assignment

type AssignmentsByRoom map[int]Assignments

type AssignmentsByBlock map[int]AssignmentsByRoom

type AssignmentsByDate map[string]AssignmentsByBlock

func (a Assignments) MapForCalendar() AssignmentsByDate {
	assignmentsMap := make(AssignmentsByDate)
	for _, assignment := range a {
		if assignment.TestEvent == nil || assignment.TestEvent.TestDate == nil {
			continue
		}
		normalizedDate := NormalizeDate(*assignment.TestEvent.TestDate)
		if _, ok := assignmentsMap[normalizedDate]; !ok {
			assignmentsMap[normalizedDate] = make(AssignmentsByBlock)
		}
		if _, ok := assignmentsMap[normalizedDate][assignment.Block]; !ok {
			assignmentsMap[normalizedDate][assignment.Block] = make(AssignmentsByRoom)
		}
		if _, ok := assignmentsMap[normalizedDate][assignment.Block][assignment.Room.ID]; !ok {
			assignmentsMap[normalizedDate][assignment.Block][assignment.Room.ID] = make(Assignments, 0)
		}
		assignmentsMap[normalizedDate][assignment.Block][assignment.Room.ID] = append(assignmentsMap[normalizedDate][assignment.Block][assignment.Room.ID], assignment)
	}
	return assignmentsMap
}

func NormalizeDate(t time.Time) string {
	return t.Format("2006-01-02")
}

func (a AssignmentsByRoom) GetRoomList() []*Room {
	rooms := make([]*Room, 0)
	for roomID := range a {
		rooms = append(rooms, a[roomID][0].Room)
	}
	// sort rooms by room.Priority
	for i := 0; i < len(rooms); i++ {
		for j := i + 1; j < len(rooms); j++ {
			if rooms[i].Priority > rooms[j].Priority {
				rooms[i], rooms[j] = rooms[j], rooms[i]
			}
		}
	}
	return rooms
}

func ParseDate(date string) time.Time {
	t, _ := time.Parse("2006-01-02", date)
	return t
}

// FilterByDates returns a new Assignments slice with only the assignments that fall between the start and end dates.
func (a Assignments) FilterByDates(start, end *time.Time) Assignments {
	var filtered Assignments
	for _, assignment := range a {
		if assignment.TestEvent.TestDate.After(*start) && assignment.TestEvent.TestDate.Before(*end) {
			filtered = append(filtered, assignment)
		}
	}
	return filtered
}

// NewAssignment creates a new assignment.
func NewAssignment(student *Student, room *Room, testEvent *TestEvent) *Assignment {
	return &Assignment{
		Student:   student,
		Room:      room,
		TestEvent: testEvent,
	}
}

func (a Assignments) FilterAssignmentsByDate(date time.Time) Assignments {
	var filtered = make(Assignments, 0)
	log.Printf("Filtering assignments by date: %v", date)
	for _, assignment := range a {
		if assignment.TestEvent.TestDate.Year() == date.Year() &&
			assignment.TestEvent.TestDate.Month() == date.Month() &&
			assignment.TestEvent.TestDate.Day() == date.Day() {
			filtered = append(filtered, assignment)
		}
	}
	log.Printf("Filtered assignments by date: %v", filtered)
	return filtered
}

func (a Assignments) FilterAssignmentsByBlock(block int) Assignments {
	var filtered = make(Assignments, 0)
	for _, assignment := range a {
		if assignment.Block != block {
			continue
		}
		filtered = append(filtered, assignment)
	}
	log.Printf("Filtered assignments by block: %v", filtered)
	return filtered
}

func (a Assignments) FilterAssignmentsByRoom(roomID int) Assignments {
	var filtered = make(Assignments, 0)
	for _, assignment := range a {
		if assignment.Room == nil {
			if roomID == -1 {
				filtered = append(filtered, assignment)
			}
			continue
		}
		if assignment.Room.ID != roomID {
			continue
		}
		filtered = append(filtered, assignment)
	}
	log.Printf("Filtered assignments by room: %v", filtered)
	return filtered
}

func (a Assignments) SortByRoomPriority() Assignments {
	var sorted Assignments
	var unassigned Assignments
	for _, assignment := range a {
		if assignment.Room == nil {
			unassigned = append(unassigned, assignment)
			sorted = append(sorted, assignment)
			continue
		}
		if len(sorted) == 0 {
			sorted = append(sorted, assignment)
			continue
		}
		inserted := false
		for i, sortedAssignment := range sorted {
			if assignment.Room.Priority < sortedAssignment.Room.Priority {
				sorted = append(sorted[:i], append(Assignments{assignment}, sorted[i:]...)...)
				inserted = true
				break
			}
		}
		if !inserted {
			sorted = append(sorted, assignment)
		}
	}
	sorted = append(sorted, unassigned...)
	return sorted
}

func (a Assignments) GetRoomList() []*Room {
	rooms := make([]*Room, 0)
	for _, assignment := range a {
		if assignment.Room == nil {
			assignment.Room = &Room{
				ID:     -1,
				Number: "Unassigned",
				Name:   "Unassigned",
			}
			continue
		}
		rooms = append(rooms, assignment.Room)
	}
	return rooms
}
