package models

// Assignment represents a student's assignment to a room.
type Assignment struct {
	ID          int
	TestEventID int
	StudentID   int
	RoomID      int
}

// AssignmentRepository provides access to the assignment storage.
type AssignmentRepository interface {
	Store(assignment *Assignment) (int, error)
	Delete(assignment *Assignment) error
	Update(assignment *Assignment) error
	FindByTestEventID(testEventID int) ([]*Assignment, error)
}

func NewAssignment(testEventID, studentID, roomID int) *Assignment {
	return &Assignment{
		TestEventID: testEventID,
		StudentID:   studentID,
		RoomID:      roomID,
	}
}


