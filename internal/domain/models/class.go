package models

type (
	Class struct {
		ID          int
		Name        string
		Block       int
		Periodicity Periodicity
		Students    []*Student
		TestEvents  []*TestEvent
		Teacher     User
	}

	Periodicity string

	ClassRepository interface {
		Repository[Class]
		DeleteAll
		AllByTeacherID(teacherID int) ([]*Class, error)
	}
)

const (
	AB Periodicity = "AB"
	A  Periodicity = "A"
	B  Periodicity = "B"
)

func NewClass(name string, block int, periodicity Periodicity, teacherID int) (*Class, error) {

	teacher := User{
		ID: teacherID,
	}

	return &Class{
		Name:        name,
		Block:       block,
		Periodicity: periodicity,
		Teacher:     teacher,
	}, nil
}
