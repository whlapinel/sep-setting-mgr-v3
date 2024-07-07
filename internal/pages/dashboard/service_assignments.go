package dashboard

import (
	"sep_setting_mgr/internal/domain/models"
	"time"
)

func (s service) GetAssignments(classID int, start, end time.Time) (models.Assignments, error) {
	var assignments models.Assignments = make(models.Assignments, 0)
	testEvents, err := s.testEvents.FindByClass(classID)
	if err != nil {
		return nil, err
	}
	for _, event := range testEvents {
		if event.TestDate.After(start) && event.TestDate.Before(end) {
			eventAssignments, err := s.assignments.GetByEventID(event.ID)
			if err != nil {
				return nil, err
			}
			assignments = append(assignments, eventAssignments...)
		}
	}
	return assignments, nil
}
