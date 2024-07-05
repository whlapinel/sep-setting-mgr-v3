package models

type (
	Room struct {
		ID          int
		Name        string
		Number      string
		MaxCapacity int
		Priority    int // defines which order to fill rooms with students
	}

	RoomRepository interface {
		Store(*Room) (int, error)
		Delete(roomID int) error
		All() ([]*Room, error)
		FindByID(roomID int) (*Room, error)
		Update(*Room) error
	}
)
