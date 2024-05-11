package domain

import (
	"time"
)

// testEvent represents a test event
type TestEvent struct {
	ID       int
	TestName string
	Class    Class
	TestDate time.Time
	Block    int
	Room     Room
}
type TestEvents []*TestEvent

type TestEventService interface {
	RegisterNewTestEvent(testName string, class Class, testDate time.Time, block int) (*TestEvent, error)
	ListAll() *TestEvents
}

// TestEventRepository provIDes access a test event store
type TestEventRepository interface {
	Store(testEvent *TestEvent) error
	// FindByID(ID string) (*TestEvent, error)
	FindAll() (*TestEvents, error)
	// Delete(ID string) error
	// FindByTestClass(testClass int) ([]*TestEvent, error)
}

func (t *TestEvent) Update(testName string, testDate time.Time) {
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
	if t[i].TestDate.Equal(t[j].TestDate) {
		return t[i].Block < t[j].Block
	}
	return t[i].TestDate.Before(t[j].TestDate)
}

// NewTestEvent creates a new test event
func NewTestEvent(testName string, class Class, testDate time.Time, block int) (*TestEvent, error) {
	return &TestEvent{
		TestName: testName,
		Class:    class,
		TestDate: testDate,
		Block:    block,
	}, nil
}
