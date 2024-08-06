package demodata

import "sep_setting_mgr/internal/domain/models"

func (ds *demoDataService) createDemoAssignments() error {

	var filterStudentsByClass = func(class *models.Class) []*models.Student {
		var students []*models.Student
		for _, student := range ds.demoData.students {
			if student.Class.ID == class.ID {
				students = append(students, student)
			}
		}
		return students
	}
	var filterEventsByClass = func(class *models.Class) []*models.TestEvent {
		var events []*models.TestEvent
		for _, event := range ds.demoData.testEvents {
			if event.Class.ID == class.ID {
				events = append(events, event)
			}
		}
		return events
	}

	for _, class := range ds.demoData.classes {
		for _, student := range filterStudentsByClass(class) {
			for _, testEvent := range filterEventsByClass(class) {

				assignment := models.NewAssignment(student, &models.Unassigned, testEvent)
				err := ds.assignmentsRepo.Store(assignment)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
