package models

import (
	"errors"
	"time"
)

type role string

type Application struct {
	ID        int
	Date      time.Time
	UserID    int
	FirstName string
	LastName  string
	Email     string
	Role      role
}

const (
	AdminRole   role = "admin"
	TeacherRole role = "teacher"
)

type Applications []*Application

type ApplicationRepository interface {
	Store(*Application) error
	Delete(*Application) error
	All() ([]*Application, error)
}

var ErrInvalidRole = errors.New("invalid role")

func GetRole(roleString string) (role, error) {
	var newRole role = role(roleString)
	switch newRole {
	case AdminRole:
		return AdminRole, nil
	case TeacherRole:
		return TeacherRole, nil
	default:
		return "", ErrInvalidRole
	}
}

func NewApplication(userID int, firstName, lastName, email string, role role) (*Application, error) {
	return &Application{
		UserID:    userID,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Role:      role,
	}, nil
}

func (a Applications) SortByDate() Applications {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i].Date.After(a[j].Date) {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	return a
}
