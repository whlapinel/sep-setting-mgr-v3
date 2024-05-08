package domain

import (
	"time"
)

// testEvent represents a test event
type TestEvent struct {
	TestName  string    // name of the test
	Id        string    // id of the test event
	TestClass int       // id of the test class
	TestDate  time.Time // date of the test
}

// NewTestEvent creates a new test event
func NewTestEvent(testName string, testClass int, testDate time.Time) *TestEvent {
	return &TestEvent{
		TestName:  testName,
		TestClass: testClass,
		TestDate:  testDate,
	}
}
