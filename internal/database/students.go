package database

import (
	"database/sql"
	"sep_setting_mgr/internal/domain"
)

type (
	studentTableRow struct {
		id         int
		first_name string
		last_name  string
		class_id   int
	}

	studentRepo struct {
		db *sql.DB
	}
)

func NewStudentsRepo(db *sql.DB) domain.StudentRepository {
	err := createStudentsTable(db)
	if err != nil {
		panic(err)
	}
	return &studentRepo{db: db}
}

func (sr *studentRepo) Store(student *domain.Student) (int, error) {
	dbStudent := convertToStudentTable(student)
	_, err := sr.db.Exec(`
	INSERT INTO students (first_name, last_name, class_id) 
	VALUES (?, ?, ?)`, dbStudent.first_name, dbStudent.last_name, dbStudent.class_id)
	if err != nil {
		return 0, err
	}
	sr.db.QueryRow(`SELECT LAST_INSERT_ID()`).Scan(&student.ID)
	return student.ID, nil
}

func (sr *studentRepo) All(classID int) ([]*domain.Student, error) {
	var students []*domain.Student
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
		err := rows.Scan(&studentTable.id, &studentTable.first_name, &studentTable.last_name, &studentTable.class_id)
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

func convertToStudentTable(student *domain.Student) studentTableRow {
	return studentTableRow{
		id:         student.ID,
		first_name: student.FirstName,
		last_name:  student.LastName,
		class_id:   student.Class.ID,
	}
}

func convertToStudent(tableRow studentTableRow) *domain.Student {
	return &domain.Student{
		ID:        tableRow.id,
		FirstName: tableRow.first_name,
		LastName:  tableRow.last_name,
		Class: domain.Class{
			ID: tableRow.class_id,
		},
	}
}

func createStudentsTable(db *sql.DB) error {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS students (
		id INT AUTO_INCREMENT PRIMARY KEY,
		first_name VARCHAR(100),
		last_name VARCHAR(100),
		class_id INT,
		FOREIGN KEY (class_id) REFERENCES classes(id)
	)`)
	if err != nil {
		return err
	}
	return nil
}
