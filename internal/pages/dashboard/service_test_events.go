package dashboard

import (
	"log"
	domain "sep_setting_mgr/internal/domain/models"
	"sep_setting_mgr/internal/util"
)

func (s service) CreateTestEvent(classID int, testName string, testDate string) (*domain.TestEvent, error) {
	log.SetPrefix("Service: ")
	log.Println("Creating test event")
	log.Println("Class ID: ", classID)
	class, err := s.classes.FindByID(classID)
	if err != nil {
		return nil, err
	}
	log.Println("Class ID: ", class.ID)
	parsedDate, err := util.ParseDate(testDate)
	if err != nil {
		return nil, err
	}
	testEvent, err := domain.NewTestEvent(testName, class, parsedDate, class.Block)
	if err != nil {
		return nil, err
	}
	event_id, err := s.testEvents.Store(testEvent)
	if err != nil {
		return nil, err
	}
	testEvent.ID = event_id
	return testEvent, nil
}

func (s service) DeleteTestEvent(testEventID int) error {
	err := s.testEvents.Delete(testEventID)
	if err != nil {
		return err
	}
	return nil
}

func (s service) ListAllTestEvents(classID int) (domain.TestEvents, error) {
	testEvents, err := s.testEvents.FindByClass(classID)
	if err != nil {
		return nil, err
	}
	return testEvents, nil
}
