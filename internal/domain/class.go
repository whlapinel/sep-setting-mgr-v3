package domain

type (
	Class struct {
		ID       int
		Name     string
		Block    int
		Students []*Student
		Teacher  User
	}

	ClassRepository interface {
		Store(*Class) error
		All() []*Class
		FindByID(classID string) (*Class, error)
	}
)

func NewClass(name string, block int) *Class {
	return &Class{
		Name:  name,
		Block: block,
	}
}
