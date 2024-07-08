package models

import (
	"sort"
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
	Store(*Assignment) error
	GetByEventID(eventID int) (Assignments, error)
}

type Assignments []*Assignment

func (a Assignments) Len() int {
	return len(a)
}

func (a Assignments) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a Assignments) Less(i, j int) bool {
	i_date := *a[i].TestEvent.TestDate
	j_date := *a[j].TestEvent.TestDate
	return i_date.Before(j_date)
}

func (a Assignments) Sort() {
	sort.Sort(a)
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

func CreateAssignment(student *Student, testEvent *TestEvent) *Assignment {

	return &Assignment{
		Student:   student,
		TestEvent: testEvent,
	}
}
