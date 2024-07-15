package dashboard

import (
	"log"
	"sep_setting_mgr/internal/domain/models"
)

func (s service) CreateAssignment(student *models.Student, testEvent *models.TestEvent) (*models.Assignment, error) {
	var assignment *models.Assignment
	priority := 1
	foundOpenRoom := false
	var max int
	for !foundOpenRoom {
		if priority > 2 {
			log.Println("looping again")
			log.Println("priority: ", priority)
		}
		room, err := s.rooms.FindByPriority(priority)
		if err != nil {
			return nil, err
		}
		if room == nil {
			return nil, nil
		}
		max = room.MaxCapacity
		if student.OneOnOne {
			max = 1
		}
		roomAssignments, err := s.rooms.GetRoomAssignments(room, testEvent.Class.Block, *testEvent.TestDate)
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

func (s service) DeleteAssignments(studentID int) error {
	return s.assignments.DeleteByStudentID(studentID)
}
