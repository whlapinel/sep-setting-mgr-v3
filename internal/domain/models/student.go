package models

type (
	Student struct {
		ID         int
		FirstName  string
		LastName   string
		Teacher    User
		Class      Class
		TestEvents []*TestEvent
		OneOnOne   bool
	}

	StudentRepository interface {
		Store(*Student) (int, error)
		All(classID int) ([]*Student, error)
		Delete(studentID int) error
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
