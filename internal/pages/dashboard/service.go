package dashboard

import (
	"sep_setting_mgr/internal/domain"
)

type (
	service struct {
		classes    domain.ClassRepository
		users      domain.UserRepository
		students   domain.StudentRepository
		testEvents domain.TestEventRepository
	}
)

func NewService(classes domain.ClassRepository, users domain.UserRepository, students domain.StudentRepository, testEvents domain.TestEventRepository) domain.DashboardService {
	return &service{
		classes:    classes,
		users:      users,
		students:   students,
		testEvents: testEvents,
	}
}
