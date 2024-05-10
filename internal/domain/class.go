package domain

import "github.com/google/uuid"

type (
	Class struct {
		ID       uuid.UUID
		Name     string
		Block    int
		Students Students
		Teacher  User
	}

	Classes []*Class

	ClassService interface {
		Service
	}

	ClassRepository interface {
		Repository
		Add(name string, block int) (*Class, error)
		All() Classes
	}
)

func NewClass(name string, block int) *Class {
	return &Class{
		ID:      uuid.New(),
		Name:    name,
		Block:   block,
	}
}

func NewClasses() *Classes {
	return &Classes{}
	// return make(Classes, 0)
}

func (c *Classes) Add(name string, block int) (*Class, error) {
	class := NewClass(name, block)
	*c = append(*c, class)
	return class, nil
}

func (c *Classes) All() Classes {
	return *c
}
