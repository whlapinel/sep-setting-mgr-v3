package database

import "github.com/google/uuid"

type (
	studentTable struct {
		ID        int
		FirstName string
		LastName  string
		Teacher   uuid.UUID
	}
)
