package demodata

import (
	"sep_setting_mgr/internal/domain/models"
)

type DemoDataService interface {
	CreateDemoData() error
	DeleteDemoData() error
}

type demoDataService struct {
	demoData
	usersRepo       models.UserRepository
	classesRepo     models.ClassRepository
	studentsRepo    models.StudentRepository
	testEventsRepo  models.TestEventRepository
	roomsRepo       models.RoomRepository
	appsRepo        models.ApplicationRepository
	assignmentsRepo models.AssignmentRepository
}

func NewDemoService(
	usersRepo models.UserRepository,
	classesRepo models.ClassRepository,
	studentsRepo models.StudentRepository,
	testEventsRepo models.TestEventRepository,
	roomsRepo models.RoomRepository,
	appsRepo models.ApplicationRepository,
	assignmentsRepo models.AssignmentRepository,
) DemoDataService {
	return &demoDataService{
		usersRepo:       usersRepo,
		classesRepo:     classesRepo,
		studentsRepo:    studentsRepo,
		testEventsRepo:  testEventsRepo,
		roomsRepo:       roomsRepo,
		appsRepo:        appsRepo,
		assignmentsRepo: assignmentsRepo,
	}
}

type demoData struct {
	users       []*models.User
	classes     []*models.Class
	students    []*models.Student
	testEvents  []*models.TestEvent
	rooms       []*models.Room
	apps        []*models.Application
	assignments []*models.Assignment
}

func (ds *demoDataService) CreateDemoData() error {
	demoUsers, err := ds.createDemoUsers()
	if err != nil {
		return err
	}
	ds.demoData.users = demoUsers
	demoClasses, err := ds.createDemoClasses()
	if err != nil {
		return err
	}
	ds.demoData.classes = demoClasses
	students, err := ds.createDemoStudents()
	if err != nil {
		return err
	}
	ds.demoData.students = students
	testEvents, err := ds.createDemoTestEvents()
	if err != nil {
		return err
	}
	ds.demoData.testEvents = testEvents
	rooms, err := ds.createDemoRooms()
	if err != nil {
		return err
	}
	ds.demoData.rooms = rooms
	err = ds.createDemoAssignments()
	if err != nil {
		return err
	}
	return nil
}

// warning: this deletes all data from the database!
func (ds *demoDataService) DeleteDemoData() error {
	err := ds.assignmentsRepo.DeleteAll()
	if err != nil {
		return err
	}
	err = ds.studentsRepo.DeleteAll()
	if err != nil {
		return err
	}
	err = ds.testEventsRepo.DeleteAll()
	if err != nil {
		return err
	}
	err = ds.roomsRepo.DeleteAll()
	if err != nil {
		return err
	}
	err = ds.appsRepo.DeleteAll()
	if err != nil {
		return err
	}
	err = ds.classesRepo.DeleteAll()
	if err != nil {
		return err
	}
	err = ds.usersRepo.DeleteAll()
	if err != nil {
		return err
	}
	return nil
}
