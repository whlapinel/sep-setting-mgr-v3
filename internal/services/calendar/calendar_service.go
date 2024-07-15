package calendar

import (
	"sep_setting_mgr/internal/domain/models"
	"sep_setting_mgr/internal/domain/services"
)

type CalendarService interface {
	GetAllAssignments() (models.Assignments, error)
}

type service struct {
	assignmentsService services.AssignmentsService
}

func NewService(assignmentsService services.AssignmentsService) CalendarService {
	return &service{
		assignmentsService,
	}
}

func (s service) GetAllAssignments() (models.Assignments, error) {
	return s.assignmentsService.GetAllAssignments()
}
