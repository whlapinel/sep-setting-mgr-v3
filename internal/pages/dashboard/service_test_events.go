package dashboard

import (
	"log"
	"sep_setting_mgr/internal/domain/models"
	"sep_setting_mgr/internal/util"
)

func (s service) CreateTestEvent(classID int, testName string, testDate string) (*models.TestEvent, error) {
	log.SetPrefix("Service: ")
	log.Println("Creating test event")
	log.Println("Class ID: ", classID)
	class, err := s.classes.FindByID(classID)
	if err != nil {
		return nil, err
	}
	log.Println("Class ID: ", class.ID)
	log.Println("Class Block: ", class.Block)
	parsedDate, err := util.ParseDate(testDate)
	if err != nil {
		return nil, err
	}
	testEvent, err := models.NewTestEvent(testName, class, parsedDate, class.Block)
	if err != nil {
		return nil, err
	}
	students, err := s.students.All(classID)
	if err != nil {
		return nil, err
	}
	event_id, err := s.testEvents.Store(testEvent)
	if err != nil {
		return nil, err
	}
	testEvent.ID = event_id
	log.Println("testEvent.Class.Block: ", testEvent.Block)
	log.Println("Test event stored")
	log.Println("Test event ID: ", testEvent.ID)
	for _, student := range students {
		assignment, err := s.CreateAssignment(student, testEvent)
		if err != nil {
			return testEvent, err
		}
		if assignment == nil {
			log.Println("Assignment not created")
			// might as well stop here since subsequent students will not be assigned either
			return testEvent, util.ErrNotAssigned
		} else {
			log.Println("Assignment created")
		}
	}
	return testEvent, nil
}

func (s service) DeleteTestEvent(testEventID int) error {
	err := s.testEvents.Delete(testEventID)
	if err != nil {
		return err
	}
	return nil
}

func (s service) ListAllTestEvents(classID int) (models.TestEvents, error) {
	testEvents, err := s.testEvents.FindByClass(classID)
	if err != nil {
		return nil, err
	}
	return testEvents, nil
}

func (s service) FindTestEventByID(testEventID int) (*models.TestEvent, error) {
	testEvent, err := s.testEvents.FindByID(testEventID)
	if err != nil {
		return nil, err
	}
	return testEvent, nil
}
