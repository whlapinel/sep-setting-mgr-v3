package domain

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
		Store(*Class) (int, error)
		Delete(classID int) error
		All(teacherID int) ([]*Class, error)
		FindByID(classID string) (*Class, error)
	}
)

func NewClass(name string, block int, teacherID int) *Class {

	teacher := User{
		ID: teacherID,
	}

	return &Class{
		Name:    name,
		Block:   block,
		Teacher: teacher,
	}
}
