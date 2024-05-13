package database

import (
	"time"
)

type (
	studentTable struct {
		ID        int
		FirstName string
		LastName  string
		Teacher   time.Time
	}
)
