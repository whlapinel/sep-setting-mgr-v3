package models

import (
	"log"
	"time"
)

// testEvent represents a test event
type TestEvent struct {
	ID       int
	TestName string
	Class    *Class
	TestDate *time.Time
	Block    int
}
type TestEvents []*TestEvent

// TestEventRepository provides access a test event store
type TestEventRepository interface {
	Repository[TestEvent]
	DeleteAll
	FindByClass(classID int) (TestEvents, error)
	FindByTeacherID(teacherID int) (TestEvents, error)
}

func (t *TestEvent) Update(testName string, testDate *time.Time) {
	t.TestName = testName
	t.TestDate = testDate
}

func (t TestEvents) Len() int {
	return len(t)
}

func (t TestEvents) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t TestEvents) Less(i, j int) bool {
	if t[i].TestDate.Equal(*t[j].TestDate) {
		return t[i].Block < t[j].Block
	}
	return t[i].TestDate.Before(*t[j].TestDate)
}

// NewTestEvent creates a new test event
func NewTestEvent(testName string, class *Class, testDate *time.Time) (*TestEvent, error) {
	log.SetPrefix("Domain: ")
	log.Println("Creating new test event")
	log.Println("Class ID: ", class.ID)

	// check for students in class

	// assign room(s)

	return &TestEvent{
		TestName: testName,
		Class:    class,
		TestDate: testDate,
		Block:    class.Block,
	}, nil
}
