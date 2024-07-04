package database

import (
	"database/sql"
	"log"
	"sep_setting_mgr/internal/domain"
	"time"
)

type (
	testEventTableRow struct {
		id        int
		test_name string
		test_date time.Time
		class_id  int
	}

	testEventsRepo struct {
		db *sql.DB
	}
)

func NewTestEventsRepo(db *sql.DB) domain.TestEventRepository {
	err := createTestEventTable(db)
	if err != nil {
		panic(err)
	}
	return &testEventsRepo{db: db}
}

func (tr *testEventsRepo) Store(testEvent *domain.TestEvent) (int, error) {
	log.SetPrefix("Repository: ")
	log.Println("Storing test event in database")
	log.Println("Class ID: ", testEvent.Class.ID)
	dbTestEvent := convertTestEventToTable(testEvent)
	_, err := tr.db.Exec(`
	INSERT INTO test_events (test_name, test_date, class_id) 
	VALUES (?, ?, ?)`, dbTestEvent.test_name, dbTestEvent.test_date, dbTestEvent.class_id)
	if err != nil {
		return 0, err
	}
	tr.db.QueryRow(`SELECT LAST_INSERT_ID()`).Scan(&testEvent.ID)
	return testEvent.ID, nil
}

func (tr *testEventsRepo) FindByClass(classID int) (*domain.TestEvents, error) {
	var testEvents domain.TestEvents
	var tableRows []testEventTableRow
	rows, err := tr.db.Query(`
	SELECT * FROM test_events 
	where class_id = ?`, classID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var testEventTableRow testEventTableRow
		err := rows.Scan(&testEventTableRow.id, &testEventTableRow.test_name, &testEventTableRow.test_date, &testEventTableRow.class_id)
		if err != nil {
			return nil, err
		}
		tableRows = append(tableRows, testEventTableRow)
	}
	for _, tableRow := range tableRows {
		testEvent := convertToTestEvent(tableRow)
		testEvents = append(testEvents, testEvent)
	}
	return &testEvents, nil
}

func convertTestEventToTable(testEvent *domain.TestEvent) testEventTableRow {
	return testEventTableRow{
		id:        testEvent.ID,
		test_name: testEvent.TestName,
		test_date: *testEvent.TestDate,
		class_id:  testEvent.Class.ID,
	}
}

func convertToTestEvent(tableRow testEventTableRow) *domain.TestEvent {
	return &domain.TestEvent{
		ID:       tableRow.id,
		TestName: tableRow.test_name,
		TestDate: &tableRow.test_date,
		Class: &domain.Class{
			ID: tableRow.class_id,
		},
	}
}

func createTestEventTable(db *sql.DB) error {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS test_events (
		id INT AUTO_INCREMENT PRIMARY KEY,
		test_name VARCHAR(100),
		test_date DATE,
		class_id INT,
		FOREIGN KEY (class_id) REFERENCES classes(id)
	)`)
	if err != nil {
		return err
	}
	return nil
}

func (tr *testEventsRepo) Delete(eventID int) error {
	_, err := tr.db.Exec(`
	DELETE FROM test_events 
	WHERE id = ?`, eventID)
	if err != nil {
		return err
	}
	return nil
}
