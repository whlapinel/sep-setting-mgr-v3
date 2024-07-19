package students

import (
	"log"
	"sep_setting_mgr/internal/domain/models"
	"sep_setting_mgr/internal/domain/services"
)

type StudentsService interface {
	AddStudent(firstName string, lastName string, classID int, oneOneOne bool) (*models.Student, error)
	UpdateStudent(firstName string, lastName string, oneOnOne bool, studentID int) (*models.Student, error)
	FindStudentByID(studentID int) (*models.Student, error)
	DeleteStudent(studentID int) error
	ListStudents(classID int) ([]*models.Student, error)
}

type service struct {
	students          models.StudentRepository
	classes           models.ClassRepository
	assignmentService services.AssignmentsService
}

func NewService(
	students models.StudentRepository,
	classes models.ClassRepository,
	as services.AssignmentsService,
) StudentsService {

	return &service{
		students:          students,
		classes:           classes,
		assignmentService: as,
	}
}

func (s service) AddStudent(firstName string, lastName string, classID int, oneOnOne bool) (*models.Student, error) {
	log.SetPrefix("Student Service: ")
	log.Println("Service: Adding student to database")
	// find teacher id
	student, err := models.NewStudent(firstName, lastName, classID, oneOnOne)
	if err != nil {
		return nil, err
	}
	class, err := s.classes.FindByID(classID)
	if err != nil {
		return nil, err
	}
	student.Class = *class
	log.Println("new student created")
	id, err := s.students.Store(student)
	if err != nil {
		return nil, err
	}
	// create assignments for student, room is null
	s.assignmentService.CreateAssignmentsForStudent(student)
	log.Println("new student stored")
	student.ID = id
	return student, nil
}

func (s service) UpdateStudent(firstName string, lastName string, oneOnOne bool, studentID int) (*models.Student, error) {
	log.Println("Service: Updating student in database")
	student, err := s.students.FindByID(studentID)
	if err != nil {
		return nil, err
	}
	student.FirstName, student.LastName, student.OneOnOne = firstName, lastName, oneOnOne
	err = s.students.Update(student)
	if err != nil {
		log.Println("Failed to update student: ", err)
		return nil, err
	}
	return student, nil
}

func (s service) FindStudentByID(studentID int) (*models.Student, error) {
	student, err := s.students.FindByID(studentID)
	if err != nil {
		return nil, err
	}
	return student, nil
}

func (s service) DeleteStudent(studentID int) error {
	log.Println("Service: Deleting student from database")
	err := s.assignmentService.DeleteAssignmentsForStudent(studentID)
	if err != nil {
		return err
	}
	err = s.students.Delete(studentID)
	if err != nil {
		return err
	}
	return nil
}

func (s service) ListStudents(classID int) ([]*models.Student, error) {
	students, err := s.students.All(classID)
	if err != nil {
		return nil, err
	}
	return students, nil
}
