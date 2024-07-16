package repositories

import (
	"database/sql"
	"log"
	"sep_setting_mgr/internal/domain/models"
	"time"
)

type roomsTableRow struct {
	id           int
	name         string
	number       string
	max_capacity int
	priority     int
}

type roomRepo struct {
	db *sql.DB
}

func NewRoomsRepo(db *sql.DB) models.RoomRepository {
	err := createRoomsTable(db)
	if err != nil {
		panic(err)
	}
	return &roomRepo{db: db}
}

func (rr *roomRepo) Store(room *models.Room) (int, error) {
	dbRoom := convertToRoomTable(room)
	_, err := rr.db.Exec(`
	INSERT INTO rooms (name, number, max_capacity, priority) 
	VALUES (?, ?, ?, ?)`, dbRoom.name, dbRoom.number, dbRoom.max_capacity, dbRoom.priority)
	if err != nil {
		return 0, err
	}
	rr.db.QueryRow(`SELECT LAST_INSERT_ID()`).Scan(&room.ID)
	return room.ID, nil
}

func (rr *roomRepo) All() ([]*models.Room, error) {
	log.SetPrefix("Rooms Repo: All()")
	var rooms []*models.Room
	var tableRows []roomsTableRow
	rows, err := rr.db.Query(`
	SELECT * FROM rooms`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var roomTableRow roomsTableRow
		err := rows.Scan(&roomTableRow.id, &roomTableRow.name, &roomTableRow.number, &roomTableRow.max_capacity, &roomTableRow.priority)
		if err != nil {
			return nil, err
		}
		tableRows = append(tableRows, roomTableRow)
	}
	for _, tableRow := range tableRows {
		room := convertToRoom(tableRow)
		rooms = append(rooms, room)
	}
	return rooms, nil
}

func (rr *roomRepo) GetRoomAssignments(room *models.Room, block int, date time.Time) (models.Assignments, error) {
	var assignments models.Assignments
	var tableRows []assignmentTableRow
	log.Println("Getting assignments for room: ", room.ID, " block: ", block, " date: ", date)
	rows, err := rr.db.Query(`
	SELECT a.*
	FROM assignments a
	LEFT JOIN test_events te ON a.event_id = te.id
	LEFT JOIN classes c ON te.class_id = c.id
	WHERE a.room_id = ? AND c.block = ? AND te.test_date = ?
	`, room.ID, block, date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var tableRow assignmentTableRow
		err := rows.Scan(
			&tableRow.id,
			&tableRow.student_id,
			&tableRow.room_id,
			&tableRow.event_id,
		)
		if err != nil {
			return nil, err
		}
		tableRows = append(tableRows, tableRow)
	}
	for _, tableRow := range tableRows {
		assignment := convertToAssignment(tableRow)
		assignments = append(assignments, assignment)
	}
	return assignments, nil
}

func (rr *roomRepo) FindByID(id int) (*models.Room, error) {
	var roomTableRow roomsTableRow
	rr.db.QueryRow(`SELECT * FROM rooms WHERE id = ?`, id).
		Scan(&roomTableRow.id, &roomTableRow.name, &roomTableRow.number, &roomTableRow.max_capacity, &roomTableRow.priority)
	room := convertToRoom(roomTableRow)
	return room, nil
}

func (rr *roomRepo) FindByPriority(priority int) (*models.Room, error) {
	var roomTableRow roomsTableRow
	err := rr.db.QueryRow(`SELECT * FROM rooms WHERE priority = ?`, priority).
		Scan(&roomTableRow.id, &roomTableRow.name, &roomTableRow.number, &roomTableRow.max_capacity, &roomTableRow.priority)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	room := convertToRoom(roomTableRow)
	return room, nil
}

func createRoomsTable(db *sql.DB) error {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS rooms (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255),
		number VARCHAR(255),
		max_capacity INT,
		priority INT
	)`)
	return err
}

func (rr *roomRepo) Update(room *models.Room) error {
	dbRoom := convertToRoomTable(room)
	_, err := rr.db.Exec(`
	UPDATE rooms 
	SET name = ?, number = ?, max_capacity = ?, priority = ?
	WHERE id = ?`, dbRoom.name, dbRoom.number, dbRoom.max_capacity, dbRoom.priority, dbRoom.id)
	if err != nil {
		return err
	}
	return nil
}

func (rr *roomRepo) Delete(id int) error {
	_, err := rr.db.Exec(`
	DELETE FROM rooms 
	WHERE id = ?`, id)
	if err != nil {
		return err
	}
	return nil
}

func convertToRoomTable(room *models.Room) roomsTableRow {
	return roomsTableRow{
		id:           room.ID,
		name:         room.Name,
		number:       room.Number,
		max_capacity: room.MaxCapacity,
		priority:     room.Priority,
	}
}

func convertToRoom(tableRow roomsTableRow) *models.Room {
	return &models.Room{
		ID:          tableRow.id,
		Name:        tableRow.name,
		Number:      tableRow.number,
		MaxCapacity: tableRow.max_capacity,
		Priority:    tableRow.priority,
	}
}
