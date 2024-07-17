package testevents

import (
	"log"
	"sep_setting_mgr/internal/domain/models"
	"sep_setting_mgr/internal/util"
)

type TestEventsService interface {
	FindByClassID(classID int) (models.TestEvents, error)
	CreateTestEvent(classID int, testName string, testDate string) (*models.TestEvent, error)
	DeleteTestEvent(testEventID int) error
	FindTestEventByID(testEventID int) (*models.TestEvent, error)
	FindByTeacherID(teacherID int) (models.TestEvents, error)
	ListAll() (models.TestEvents, error)
}

type service struct {
	testEvents models.TestEventRepository
	classes    models.ClassRepository
}

func NewService(testEvents models.TestEventRepository,
	classes models.ClassRepository,
) TestEventsService {

	return &service{
		testEvents,
		classes,
	}
}

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
	event_id, err := s.testEvents.Store(testEvent)
	if err != nil {
		return nil, err
	}
	testEvent.ID = event_id
	log.Println("Test event stored")
	return testEvent, nil
}

func (s service) DeleteTestEvent(testEventID int) error {
	err := s.testEvents.Delete(testEventID)
	if err != nil {
		return err
	}
	return nil
}

func (s service) FindByClassID(classID int) (models.TestEvents, error) {
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

func (s service) FindByTeacherID(teacherID int) (models.TestEvents, error) {
	testEvents, err := s.testEvents.FindByTeacherID(teacherID)
	if err != nil {
		return nil, err
	}
	return testEvents, nil
}

func (s service) ListAll() (models.TestEvents, error) {
	testEvents, err := s.testEvents.ListAll()
	if err != nil {
		return nil, err
	}
	return testEvents, nil
}
