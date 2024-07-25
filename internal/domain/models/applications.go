package models

import (
	"errors"
	"time"
)

type Application struct {
	ID        int
	Date      time.Time
	UserID    int
	FirstName string
	LastName  string
	Email     string
	Role      string
}

const (
	AdminRole   string = "admin"
	TeacherRole string = "teacher"
)

const (
	Approve string = "approve"
	Deny    string = "deny"
)

type Applications []*Application

type ApplicationRepository interface {
	Store(*Application) error
	Delete(*Application) error
	All() (Applications, error)
	GetApplicationByID(id int) (*Application, error)
	GetApplicationsByUserID(userID int) (Applications, error)
}

var ErrInvalidRole = errors.New("invalid role")

var ErrInvalidAction = errors.New("invalid action")

func NewApplication(userID int, firstName, lastName, email string, role string) (*Application, error) {
	if role != AdminRole && role != TeacherRole {
		return nil, ErrInvalidRole
	}
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
