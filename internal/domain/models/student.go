package models

type StudentTestEvent struct {
	Room  Room
	Event TestEvent
}

type (
	Student struct {
		ID         int
		FirstName  string
		LastName   string
		Teacher    User
		Class      Class
		TestEvents StudentTestEvent
		OneOnOne   bool
	}

	StudentRepository interface {
		Repository[Student]
		DeleteAll
		AllInClass(classID int) ([]*Student, error)
	}
)

func NewStudent(firstName string, lastName string, classID int, oneOnOne bool) (*Student, error) {
	return &Student{
		FirstName: firstName,
		LastName:  lastName,
		Class:     Class{ID: classID},
		OneOnOne:  oneOnOne,
	}, nil
}
