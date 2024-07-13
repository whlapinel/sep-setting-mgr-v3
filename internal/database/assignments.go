package database

import (
	"database/sql"
	"log"
	"sep_setting_mgr/internal/domain/models"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type assignmentTableRow struct {
	id         int
	event_id   int
	student_id int
	room_id    int
}

type assignmentRepo struct {
	db *sql.DB
}

func NewAssignmentsRepo(db *sql.DB) models.AssignmentRepository {
	createAssignmentsTable(db)
	return &assignmentRepo{db: db}
}

func (ar *assignmentRepo) Delete(assignmentID int) error {
	_, err := ar.db.Exec(`DELETE FROM assignments WHERE id = ?`, assignmentID)
	if err != nil {
		return err
	}
	return nil
}

func (ar *assignmentRepo) All() (models.Assignments, error) {
	var assignments models.Assignments

	rows, err := ar.db.Query(`
	SELECT a.*, te.*, s.* 
	FROM assignments a
	JOIN test_events te ON a.event_id = te.id
	JOIN students s ON a.student_id = s.id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var assignmentTable assignmentTableRow
		var eventTable testEventTableRow
		var studentTable studentTableRow

		var temp []uint8

		err := rows.Scan(
			&assignmentTable.id,
			&assignmentTable.event_id,
			&assignmentTable.student_id,
			&assignmentTable.room_id,
			&eventTable.id,
			&eventTable.test_name,
			&temp,
			&eventTable.class_id,
			&studentTable.id,
			&studentTable.first_name,
			&studentTable.last_name,
			&studentTable.class_id,
			&studentTable.one_on_one,
		)
		if err != nil {
			return nil, err
		}
		eventTable.test_date, err = time.Parse("2006-01-02", string(temp))
		if err != nil {
			return nil, err
		}
		assignment := convertToAssignment(assignmentTable)
		event := convertToTestEvent(eventTable)
		student := convertToStudent(studentTable)

		assignment.TestEvent = event
		assignment.Student = student

		assignments = append(assignments, assignment)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return assignments, nil
}

func (ar *assignmentRepo) DeleteByStudentID(studentID int) error {
	_, err := ar.db.Exec(`DELETE FROM assignments WHERE student_id = ?`, studentID)
	if err != nil {
		return err
	}
	return nil
}

func (ar *assignmentRepo) GetByEventID(eventID int) (models.Assignments, error) {
	var assignments models.Assignments

	rows, err := ar.db.Query(`
	SELECT a.*, te.*, s.* 
	FROM assignments a
	JOIN test_events te ON a.event_id = te.id
	JOIN students s ON a.student_id = s.id
	WHERE a.event_id = ?
	`, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var assignmentTable assignmentTableRow
		var eventTable testEventTableRow
		var studentTable studentTableRow

		var temp []uint8

		err := rows.Scan(
			&assignmentTable.id,
			&assignmentTable.event_id,
			&assignmentTable.student_id,
			&assignmentTable.room_id,
			&eventTable.id,
			&eventTable.test_name,
			&temp,
			&eventTable.class_id,
			&studentTable.id,
			&studentTable.first_name,
			&studentTable.last_name,
			&studentTable.class_id,
			&studentTable.one_on_one,
		)
		if err != nil {
			return nil, err
		}
		eventTable.test_date, err = time.Parse("2006-01-02", string(temp))
		if err != nil {
			return nil, err
		}
		assignment := convertToAssignment(assignmentTable)
		event := convertToTestEvent(eventTable)
		log.Println("event date: ", event.TestDate.Format("2006-01-02"))
		student := convertToStudent(studentTable)
		log.Println("student first name: ", student.FirstName)

		assignment.TestEvent = event
		assignment.Student = student

		assignments = append(assignments, assignment)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return assignments, nil
}

func (ar *assignmentRepo) Store(assignment *models.Assignment) error {
	dbAssignment := convertToAssignmentTable(assignment)
	log.Println("Adding assignment to database")
	_, err := ar.db.Exec(
		`INSERT INTO assignments (event_id, student_id, room_id) VALUES (?, ?, ?)`,
		dbAssignment.event_id, dbAssignment.student_id, dbAssignment.room_id)
	if err != nil {
		return err
	}
	ar.db.QueryRow(`SELECT LAST_INSERT_ID()`).Scan(&dbAssignment.id)
	return nil
}

func createAssignmentsTable(db *sql.DB) error {
	log.Println("Creating assignments table")
	result, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS 
	assignments (
		id int AUTO_INCREMENT PRIMARY KEY, 
		event_id int,
		student_id int,
		room_id int,
		FOREIGN KEY (event_id) REFERENCES test_events(id),
		FOREIGN KEY (student_id) REFERENCES students(id),
		FOREIGN KEY (room_id) REFERENCES rooms(id)
		)`)
	if err != nil {
		return err
	}
	log.Println(result)

	return nil
}

func convertToAssignment(dbAssignment assignmentTableRow) *models.Assignment {
	return &models.Assignment{
		ID: dbAssignment.id,
		TestEvent: &models.TestEvent{
			ID: dbAssignment.event_id,
		},
		Student: &models.Student{
			ID: dbAssignment.student_id,
		},
		Room: &models.Room{
			ID: dbAssignment.room_id,
		},
	}
}

func convertToAssignmentTable(assignment *models.Assignment) assignmentTableRow {
	var assignmentTable assignmentTableRow
	assignmentTable.event_id = assignment.TestEvent.ID
	assignmentTable.student_id = assignment.Student.ID
	assignmentTable.room_id = assignment.Room.ID
	return assignmentTable
}
