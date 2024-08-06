package repositories

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
	room_id    *int
}

type assignmentRepo struct {
	db *sql.DB
}

func NewAssignmentsRepo(db *sql.DB) models.AssignmentRepository {
	err := createAssignmentsTable(db)
	if err != nil {
		panic(err)
	}
	return &assignmentRepo{db: db}
}

func (ar *assignmentRepo) DeleteAll() error {
	_, err := ar.db.Exec(`DELETE FROM assignments`)
	if err != nil {
		return err
	}
	return nil
}

func (ar *assignmentRepo) Update(as *models.Assignment) error {
	log.Println("Updating assignment. ID: ", as.ID)
	if as.Room.ID < 0 {
		log.Printf("Nullifying room for assignment %d", as.ID)
		err := ar.nullifyRoomIDForAssignment(as)
		if err != nil {
			return err
		}
		return nil
	}
	_, err := ar.db.Exec(`UPDATE assignments SET room_id = ? WHERE id = ?`, as.Room.ID, as.ID)
	if err != nil {
		return err
	}
	return nil
}

func (ar *assignmentRepo) FindByID(id int) (*models.Assignment, error) {
	var classRow classTableRow
	var userRow userTableRow
	var studentRow studentTableRow
	var eventRow testEventTableRow
	var roomRow roomsTableRow
	var tempDate []uint8

	row := ar.db.QueryRow(
		`SELECT c.block, c.name, u.email, s.first_name, s.last_name, s.one_on_one, e.test_name, e.test_date, r.name, r.number, r.max_capacity
		FROM assignments a
		JOIN test_events e ON e.id = a.event_id
		JOIN students s ON a.student_id = s.id
		JOIN classes c ON c.id = s.class_id
		JOIN users u ON u.id = c.teacher_id
		LEFT JOIN rooms r ON r.id = a.room_id
		WHERE a.id = ?
		`, id)
	err := row.Scan(
		&classRow.block,
		&classRow.name,
		&userRow.email,
		&studentRow.first_name,
		&studentRow.last_name,
		&studentRow.one_on_one,
		&eventRow.test_name,
		&tempDate,
		&roomRow.name,
		&roomRow.number,
		&roomRow.max_capacity,
	)
	if err != nil {
		return nil, err
	}
	eventRow.test_date, err = time.Parse("2006-01-02", string(tempDate))
	if err != nil {
		return nil, err
	}
	class := convertToClass(classRow)
	user := convertFromTable(userRow)
	student := convertToStudent(studentRow)
	event := convertToTestEvent(eventRow)
	room := convertToRoom(roomRow)

	student.Class = *class
	student.Teacher = *user

	assignment := models.NewAssignment(student, room, event)
	assignment.ID = id

	return assignment, nil
}

func (ar *assignmentRepo) Delete(assignmentID int) error {
	_, err := ar.db.Exec(`DELETE FROM assignments WHERE id = ?`, assignmentID)
	if err != nil {
		return err
	}
	return nil
}

// NullifyRoomID sets the room_id of all assignments with the given roomID to NULL
func (ar *assignmentRepo) NullifyRoomID(roomID int) error {
	_, err := ar.db.Exec(`UPDATE assignments SET room_id = NULL WHERE room_id = ?`, roomID)
	if err != nil {
		return err
	}
	return nil
}

func (ar *assignmentRepo) nullifyRoomIDForAssignment(as *models.Assignment) error {
	_, err := ar.db.Exec(`UPDATE assignments SET room_id = NULL WHERE id = ?`, as.ID)
	if err != nil {
		return err
	}
	return nil
}

func (ar *assignmentRepo) DeleteByEventID(eventID int) error {
	_, err := ar.db.Exec(`DELETE FROM assignments WHERE event_id = ?`, eventID)
	if err != nil {
		return err
	}
	return nil
}

func (ar *assignmentRepo) GetByTeacherID(teacherID int) (models.Assignments, error) {
	log.SetPrefix("Assignments Repo: All()")
	var assignments models.Assignments
	rows, err := ar.db.Query(`
	SELECT a.*, te.*, s.*, c.block, r.*
	FROM assignments a
	JOIN test_events te ON a.event_id = te.id
	JOIN students s ON a.student_id = s.id
	JOIN classes c ON te.class_id = c.id
	LEFT JOIN rooms r ON a.room_id = r.id
	WHERE c.teacher_id = ?
	`, teacherID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var assignmentTable assignmentTableRow
		var eventTable testEventTableRow
		var studentTable studentTableRow
		var classTable classTableRow
		var roomsTable roomsTableRow

		var temp []uint8

		err := rows.Scan(
			// assignment
			&assignmentTable.id,
			&assignmentTable.event_id,
			&assignmentTable.student_id,
			&assignmentTable.room_id,
			// event
			&eventTable.id,
			&eventTable.test_name,
			&temp,
			&eventTable.class_id,
			// student
			&studentTable.id,
			&studentTable.first_name,
			&studentTable.last_name,
			&studentTable.class_id,
			&studentTable.one_on_one,
			// classes
			&classTable.block,
			// room
			&roomsTable.id,
			&roomsTable.name,
			&roomsTable.number,
			&roomsTable.max_capacity,
			&roomsTable.priority,
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
		room := convertToRoom(roomsTable)

		assignment.TestEvent = event
		assignment.Student = student
		assignment.Room = room
		assignment.Block = classTable.block

		assignments = append(assignments, assignment)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return assignments, nil
}

func (ar *assignmentRepo) All() ([]*models.Assignment, error) {
	log.SetPrefix("Assignments Repo: All()")
	var assignments models.Assignments
	rows, err := ar.db.Query(`
	SELECT a.*, te.*, s.*, c.block, r.*
	FROM assignments a
	JOIN test_events te ON a.event_id = te.id
	JOIN students s ON a.student_id = s.id
	JOIN classes c ON te.class_id = c.id
	LEFT JOIN rooms r ON a.room_id = r.id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var assignmentTable assignmentTableRow
		var eventTable testEventTableRow
		var studentTable studentTableRow
		var classTable classTableRow
		var roomsTable roomsTableRow

		var temp []uint8

		err := rows.Scan(
			// assignment
			&assignmentTable.id,
			&assignmentTable.event_id,
			&assignmentTable.student_id,
			&assignmentTable.room_id,
			// event
			&eventTable.id,
			&eventTable.test_name,
			&temp,
			&eventTable.class_id,
			// student
			&studentTable.id,
			&studentTable.first_name,
			&studentTable.last_name,
			&studentTable.class_id,
			&studentTable.one_on_one,
			// classes
			&classTable.block,
			// room
			&roomsTable.id,
			&roomsTable.name,
			&roomsTable.number,
			&roomsTable.max_capacity,
			&roomsTable.priority,
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
		room := convertToRoom(roomsTable)

		assignment.TestEvent = event
		assignment.Student = student
		assignment.Room = room
		assignment.Block = classTable.block

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

func (ar *assignmentRepo) FindByEventID(eventID int) (models.Assignments, error) {
	var assignments models.Assignments

	rows, err := ar.db.Query(`
	SELECT a.*, te.*, s.*, r.*
	FROM assignments a
	JOIN test_events te ON a.event_id = te.id
	JOIN students s ON a.student_id = s.id
	JOIN rooms r ON a.room_id = r.id
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
		var roomsTable roomsTableRow

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
			&roomsTable.id,
			&roomsTable.name,
			&roomsTable.number,
			&roomsTable.max_capacity,
			&roomsTable.priority,
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
		room := convertToRoom(roomsTable)

		assignment.TestEvent = event
		assignment.Student = student
		assignment.Room = room

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
	result, err := ar.db.Exec(
		`INSERT INTO assignments (event_id, student_id, room_id) VALUES (?, ?, ?)`,
		dbAssignment.event_id, dbAssignment.student_id, dbAssignment.room_id)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	assignment.ID = int(id)
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
	var room *models.Room
	if dbAssignment.room_id == nil {
		room = nil
	} else {
		room = &models.Room{
			ID: *dbAssignment.room_id,
		}

	}
	return &models.Assignment{
		ID: dbAssignment.id,
		TestEvent: &models.TestEvent{
			ID: dbAssignment.event_id,
		},
		Student: &models.Student{
			ID: dbAssignment.student_id,
		},
		Room: room,
	}
}

func convertToAssignmentTable(assignment *models.Assignment) assignmentTableRow {
	var assignmentTable assignmentTableRow
	assignmentTable.event_id = assignment.TestEvent.ID
	assignmentTable.student_id = assignment.Student.ID
	if assignment.Room == nil || assignment.Room.ID < 0 {
		assignmentTable.room_id = nil
	} else {
		assignmentTable.room_id = &assignment.Room.ID
	}
	return assignmentTable
}
