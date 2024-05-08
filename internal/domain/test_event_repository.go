package domain

// TestEventRepository provides access a test event store
type TestEventRepository interface {
	Store(testEvent *TestEvent) error
	FindById(id string) (*TestEvent, error)
	FindAll() ([]*TestEvent, error)
	Delete(id string) error
	FindByTestClass(testClass int) ([]*TestEvent, error)
}
