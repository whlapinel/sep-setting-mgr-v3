package database

import (
	"database/sql"
	"log"
	"sep_setting_mgr/internal/domain"

	_ "github.com/go-sql-driver/mysql"
)

type classTableRow struct {
	id         int
	name       string
	block      int
	teacher_id int
}

type classRepo struct {
	db *sql.DB
}

func NewClassesRepo(db *sql.DB) domain.ClassRepository {
	createClassesTable(db)
	return &classRepo{db: db}
}

func (cr *classRepo) Delete(classID int) error {
	_, err := cr.db.Exec(`DELETE FROM classes WHERE id = ?`, classID)
	if err != nil {
		return err
	}
	return nil
}

func (cr *classRepo) Store(class *domain.Class) (int, error) {
	dbClass := convertToClassTable(class)
	log.Println("Adding class to database")
	_, err := cr.db.Exec(`INSERT INTO classes (name, block, teacher_id) VALUES (?, ?, ?)`, dbClass.name, dbClass.block, dbClass.teacher_id)
	if err != nil {
		return 0, err
	}
	cr.db.QueryRow(`SELECT LAST_INSERT_ID()`).Scan(&dbClass.id)
	return dbClass.id, nil
}

func (classRepo *classRepo) All(teacherID int) ([]*domain.Class, error) {
	var classes []*domain.Class
	var tableRows []classTableRow
	rows, err := classRepo.db.Query(`
	SELECT * FROM classes 
	where teacher_id = ? 
	order by block`, teacherID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var classTable classTableRow
		err := rows.Scan(&classTable.id, &classTable.name, &classTable.block, &classTable.teacher_id)
		if err != nil {
			return nil, err
		}
		tableRows = append(tableRows, classTable)
	}
	for _, tableRow := range tableRows {
		class := convertToClass(tableRow)
		classes = append(classes, class)
	}
	return classes, nil
}

func (classRepo *classRepo) FindByID(classID int) (*domain.Class, error) {
	var dbClass classTableRow

	class := convertToClass(dbClass)
	return class, nil
}

func createClassesTable(db *sql.DB) error {
	log.Println("Creating classes table")
	result, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS 
	classes (
		id int AUTO_INCREMENT PRIMARY KEY, 
		name VARCHAR(255) NOT NULL, 
		block INT NOT NULL, 
		teacher_id int NOT NULL,
		constraint block_teacher_id unique (block, teacher_id)
		)`)
	if err != nil {
		return err
	}
	log.Println(result)

	return nil
}

func convertToClass(dbClass classTableRow) *domain.Class {
	return &domain.Class{
		ID:    dbClass.id,
		Name:  dbClass.name,
		Block: dbClass.block,
	}
}

func convertToClassTable(class *domain.Class) classTableRow {
	var classTable classTableRow
	// if this is a new class then the ID will be 0
	classTable.name = class.Name
	classTable.block = class.Block
	classTable.teacher_id = class.Teacher.ID
	return classTable
}
