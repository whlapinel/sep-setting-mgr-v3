package dashboard

import (
	"sep_setting_mgr/internal/domain"
)

type (
	service struct {
		classes  domain.ClassRepository
		users    domain.UserRepository
		students domain.StudentRepository
	}
)

func NewService(classes domain.ClassRepository, users domain.UserRepository, students domain.StudentRepository) domain.DashboardService {
	return &service{
		classes:  classes,
		users:    users,
		students: students,
	}
}
