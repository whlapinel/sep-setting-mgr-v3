package models

type (
	Class struct {
		ID         int
		Name       string
		Block      int
		Students   []*Student
		TestEvents []*TestEvent
		Teacher    User
	}

	ClassRepository interface {
		Repository[Class]
		DeleteAll
		AllByTeacherID(teacherID int) ([]*Class, error)
	}
)

func NewClass(name string, block int, teacherID int) (*Class, error) {

	teacher := User{
		ID: teacherID,
	}

	return &Class{
		Name:    name,
		Block:   block,
		Teacher: teacher,
	}, nil
}
