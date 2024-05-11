package domain

import (
	"fmt"

	"github.com/google/uuid"
)

type (
	Class struct {
		ID       uuid.UUID
		Name     string
		Block    int
		Students Students
		Teacher  User
	}

	ClassRepository interface {
		Add(name string, block int) (*Class, error)
		All() Classes
		FindByID(classID string) (*Class, error)
	}

	Classes []*Class
)

func NewClass(name string, block int) *Class {
	return &Class{
		ID:    uuid.New(),
		Name:  name,
		Block: block,
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

func (c *Classes) FindByID(classID string) (*Class, error) {
	for _, class := range *c {
		if class.ID.String() == classID {
			return class, nil
		}
	}
	return nil, fmt.Errorf("Class not found with ID: %s", classID)
}
