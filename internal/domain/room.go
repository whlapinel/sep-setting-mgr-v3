package domain

type (
	Room struct {
		ID	int
		Name string
		Number string
		Priority int // defines which order to fill rooms with students
	}

	RoomService interface {
		RegisterNewRoom(name string, number string)
		UpdatePriority(priority int)
	}


)