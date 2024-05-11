package database

import (
	"sep_setting_mgr/internal/domain"

	"github.com/google/uuid"
)

type database struct {
}

type classTable struct {
	ID        uuid.UUID
	Name      string
	Block     int
	TeacherID uuid.UUID
}

func NewDatabase() domain.ClassRepository {
	return &database{}
}

func (db *database) Add(name string, block int) (*domain.Class, error) {
	var class *domain.Class

	return class, nil
}

func (db *database) All() domain.Classes {
	var classes domain.Classes

	return classes
}

func (db *database) FindByID(classID string) (*domain.Class, error) {
	var class *domain.Class

	return class, nil
}
