package models

type (
	Room struct {
		ID	int
		Name string
		Number string
		Priority int // defines which order to fill rooms with students
	}
)