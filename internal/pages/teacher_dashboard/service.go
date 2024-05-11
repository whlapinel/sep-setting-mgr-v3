package teacher_dashboard

import (
	"sep_setting_mgr/internal/domain"
)

type (
	Service interface {
		// List returns a copy of the todos list
		List() (domain.Classes, error)
	}

	service struct {
		classes domain.ClassRepository
	}
)

func NewService(classes domain.ClassRepository) Service {
	return &service{
		classes: classes,
	}
}

func (s service) List() (domain.Classes, error) {
	return s.classes.All(), nil
}
