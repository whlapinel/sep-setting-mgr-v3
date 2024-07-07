package dashboard

import (
	"sep_setting_mgr/internal/domain/models"
	"sep_setting_mgr/internal/domain/pages"
)

type (
	service struct {
		classes     models.ClassRepository
		users       models.UserRepository
		students    models.StudentRepository
		testEvents  models.TestEventRepository
		rooms       models.RoomRepository
		assignments models.AssignmentRepository
	}
)

func NewService(classes models.ClassRepository, users models.UserRepository, students models.StudentRepository, testEvents models.TestEventRepository, rooms models.RoomRepository, assignments models.AssignmentRepository) pages.DashboardService {
	return &service{
		classes:     classes,
		users:       users,
		students:    students,
		testEvents:  testEvents,
		rooms:       rooms,
		assignments: assignments,
	}
}
