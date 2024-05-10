package classes

import (
	"sep_setting_mgr/internal/domain"
)

type (
	Service interface {
		Add(name string, block int) (*domain.Class, error)
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

func (s service) Add(name string, block int) (*domain.Class, error) {
	class, err := s.classes.Add(name, block)
	if err != nil {
		return nil, err
	}
	return class, nil
}
