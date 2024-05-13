package database

import (
	"database/sql"
	"log"
	"sep_setting_mgr/internal/domain"

	_ "github.com/go-sql-driver/mysql"
)

type classTable struct {
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

func (cr *classRepo) Store(class *domain.Class) error {
	dbClass := convertToClassTable(class)
	log.Println("Adding class to database")
	_, err := cr.db.Exec(`INSERT INTO classes (id, name, block) VALUES (?, ?, ?)`, dbClass.id, dbClass.name, dbClass.block)
	if err != nil {
		return err
	}
	return nil
}

func (classRepo *classRepo) All() []*domain.Class {
	var classes []*domain.Class

	return classes
}

func (classRepo *classRepo) FindByID(classID string) (*domain.Class, error) {
	var dbClass classTable

	class := convertToClass(dbClass)
	return class, nil
}

func createClassesTable(db *sql.DB) error {
	log.Println("Creating classes table")
	result, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS 
	classes (
		id int AUTO_INCREMENT PRIMARY KEY, 
		name VARCHAR(255), 
		block INT, 
		teacher_id int
		)`)
	if err != nil {
		return err
	}
	log.Println(result)

	return nil
}

func convertToClass(dbClass classTable) *domain.Class {
	return &domain.Class{
		ID:    dbClass.id,
		Name:  dbClass.name,
		Block: dbClass.block,
	}
}

func convertToClassTable(class *domain.Class) classTable {
	return classTable{
		name:  class.Name,
		block: class.Block,
	}
}
