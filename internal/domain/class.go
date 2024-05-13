package domain

type (
	Class struct {
		ID       int
		Name     string
		Block    int
		Students Students
		Teacher  User
	}

	ClassRepository interface {
		Store(*Class) error
		All() Classes
		FindByID(classID string) (*Class, error)
	}

	Classes []*Class
)

func NewClass(name string, block int) *Class {
	return &Class{
		Name:  name,
		Block: block,
	}
}
