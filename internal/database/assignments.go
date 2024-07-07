package database

import (
	"database/sql"
	"log"
	"sep_setting_mgr/internal/domain/models"

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

func (ar *assignmentRepo) Update(assignment *models.Assignment) error {
	dbAssignment := convertToAssignmentTable(assignment)
	_, err := ar.db.Exec(`
	UPDATE assignments 
	SET event_id = ?, student_id = ?, room_id = ?
	WHERE id = ?`, dbAssignment.event_id, dbAssignment.student_id, dbAssignment.room_id, dbAssignment.id)
	if err != nil {
		return err
	}
	return nil
}

func (ar *assignmentRepo) GetByEventID(eventID int) (models.Assignments, error) {
	var assignments []*models.Assignment
	var tableRows []assignmentTableRow
	rows, err := ar.db.Query(`
	SELECT * FROM assignments 
	where event_id = ?`, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var assignmentTable assignmentTableRow
		err := rows.Scan(&assignmentTable.id, &assignmentTable.event_id, &assignmentTable.student_id, &assignmentTable.room_id)
		if err != nil {
			return nil, err
		}
		tableRows = append(tableRows, assignmentTable)
	}
	for _, tableRow := range tableRows {
		assignment := convertToAssignment(tableRow)
		assignments = append(assignments, assignment)
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
		FOREIGN KEY (room_id) REFERENCES rooms(id),
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
