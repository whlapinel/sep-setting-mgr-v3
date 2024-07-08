package database

import (
	"database/sql"
	"log"
	"sep_setting_mgr/internal/domain/models"
)

type (
	studentTableRow struct {
		id         int
		first_name string
		last_name  string
		class_id   int
		one_on_one bool
	}

	studentRepo struct {
		db *sql.DB
	}
)

func NewStudentsRepo(db *sql.DB) models.StudentRepository {
	err := createStudentsTable(db)
	if err != nil {
		panic(err)
	}
	return &studentRepo{db: db}
}

func (sr *studentRepo) GetAssignments(studentID int) (models.Assignments, error) {
	return nil, nil
}


func (sr *studentRepo) Store(student *models.Student) (int, error) {
	dbStudent := convertToStudentTable(student)
	_, err := sr.db.Exec(`
	INSERT INTO students (first_name, last_name, class_id, one_on_one) 
	VALUES (?, ?, ?, ?)`, dbStudent.first_name, dbStudent.last_name, dbStudent.class_id, dbStudent.one_on_one)
	if err != nil {
		return 0, err
	}
	sr.db.QueryRow(`SELECT LAST_INSERT_ID()`).Scan(&student.ID)
	return student.ID, nil
}

func (sr *studentRepo) All(classID int) ([]*models.Student, error) {
	var students []*models.Student
	var tableRows []studentTableRow
	rows, err := sr.db.Query(`
	SELECT * FROM students 
	where class_id = ?`, classID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var studentTable studentTableRow
		err := rows.Scan(&studentTable.id, &studentTable.first_name, &studentTable.last_name, &studentTable.class_id, &studentTable.one_on_one)
		if err != nil {
			return nil, err
		}
		tableRows = append(tableRows, studentTable)
	}
	for _, tableRow := range tableRows {
		student := convertToStudent(tableRow)
		students = append(students, student)
	}
	return students, nil
}

func convertToStudentTable(student *models.Student) studentTableRow {
	return studentTableRow{
		id:         student.ID,
		first_name: student.FirstName,
		last_name:  student.LastName,
		class_id:   student.Class.ID,
		one_on_one: student.OneOnOne,
	}
}

func convertToStudent(tableRow studentTableRow) *models.Student {
	return &models.Student{
		ID:        tableRow.id,
		FirstName: tableRow.first_name,
		LastName:  tableRow.last_name,
		Class: models.Class{
			ID: tableRow.class_id,
		},
		OneOnOne: tableRow.one_on_one,
	}
}

func createStudentsTable(db *sql.DB) error {
	log.Println("Creating students table")
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS students (
		id INT AUTO_INCREMENT PRIMARY KEY,
		first_name VARCHAR(100),
		last_name VARCHAR(100),
		class_id INT,
		one_on_one BOOLEAN,
		FOREIGN KEY (class_id) REFERENCES classes(id)
	)`)
	if err != nil {
		return err
	}
	return nil
}

func (sr *studentRepo) Delete(studentID int) error {
	_, err := sr.db.Exec(`
	DELETE FROM students 
	WHERE id = ?`, studentID)
	if err != nil {
		return err
	}
	return nil
}
