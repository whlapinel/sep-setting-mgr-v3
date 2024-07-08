package dashboard

import (
	"fmt"
	"log"
	"sep_setting_mgr/internal/domain/models"
	"time"
)

func (s service) CreateAssignment(student *models.Student, testEvent *models.TestEvent) (*models.Assignment, error) {
	var assignment *models.Assignment
	priority := 1
	foundOpenRoom := false
	var max int
	for !foundOpenRoom {
		log.Println("looping again")
		log.Println("priority: ", priority)
		room, err := s.rooms.FindByPriority(priority)
		if err != nil {
			return nil, err
		}
		log.Println("room: ", room)
		log.Println("room name: ", room.Name)
		log.Println("room max capacity: ", room.MaxCapacity)
		if room.Name == "" {
			return nil, fmt.Errorf("all rooms are full")
		}
		max = room.MaxCapacity
		if student.OneOnOne {
			max = 1
		}
		roomAssignments, err := s.rooms.GetRoomAssignments(room, *testEvent.TestDate)
		if err != nil {
			return nil, err
		}
		log.Println("room assignments: ", roomAssignments)
		log.Println("room assignments length: ", len(roomAssignments))
		if len(roomAssignments) < max {
			log.Println("room is not full")
			foundOpenRoom = true
			assignment = models.NewAssignment(student, room, testEvent)
			err := s.assignments.Store(assignment)
			if err != nil {
				return nil, err
			}
		}
		priority++
	}
	return assignment, nil
}

func (s service) GetAssignments(classID int, start, end time.Time) (models.Assignments, error) {
	log.SetPrefix("Service: GetAssignments()")
	var assignments models.Assignments = make(models.Assignments, 0)
	testEvents, err := s.testEvents.FindByClass(classID)
	if err != nil {
		return nil, err
	}
	for _, event := range testEvents {
		if event.TestDate.After(start) && event.TestDate.Before(end) {
			eventAssignments, err := s.assignments.GetByEventID(event.ID)
			log.Println("eventAssignments[0].TestEvent.ID: ", eventAssignments[0].TestEvent.ID)
			log.Println("eventAssignments[0].TestEvent.TestDate: ", eventAssignments[0].TestEvent.TestDate)
			if err != nil {
				return nil, err
			}
			assignments = append(assignments, eventAssignments...)
		}
	}
	testEvent := 
	return assignments, nil
}
