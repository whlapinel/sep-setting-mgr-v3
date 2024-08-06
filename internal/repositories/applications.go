package repositories

import (
	"database/sql"
	"errors"
	"sep_setting_mgr/internal/domain/models"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type applicationTableRow struct {
	id      int
	date    time.Time
	user_id int
	role    string
}

type applicationRepo struct {
	db *sql.DB
}

func NewApplicationRepo(db *sql.DB) models.ApplicationRepository {
	err := createApplicationsTable(db)
	if err != nil {
		panic(err)
	}
	return &applicationRepo{db: db}
}

func createApplicationsTable(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS applications (
			id INT AUTO_INCREMENT PRIMARY KEY,
			date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			user_id INT NOT NULL,
			role VARCHAR(255) NOT NULL
		);
	`
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (r *applicationRepo) DeleteAll() error {
	query := `DELETE FROM applications;`
	_, err := r.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (r *applicationRepo) Store(a *models.Application) error {
	query := `INSERT INTO applications (user_id, role) VALUES (?, ?);`
	_, err := r.db.Exec(query, a.UserID, a.Role)
	if err != nil {
		return err
	}
	return nil
}

func (r *applicationRepo) Update(a *models.Application) error {
	return errors.New("not implemented")
}

func (r *applicationRepo) Delete(id int) error {
	query := `DELETE FROM applications WHERE id = ?;`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *applicationRepo) All() ([]*models.Application, error) {
	query := `
	SELECT a.*, u.first_name, u.last_name, u.email
	FROM applications a
	JOIN users u ON u.id = a.user_id
	ORDER BY a.date ASC
	;`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var apps []*models.Application
	for rows.Next() {
		var row applicationTableRow
		var user userTableRow
		var temp []uint8
		err := rows.Scan(&row.id, &temp, &row.user_id, &row.role, &user.first_name, &user.last_name, &user.email)
		if err != nil {
			return nil, err
		}
		row.date, err = time.Parse("2006-01-02 15:04:05", string(temp))
		if err != nil {
			return nil, err
		}
		app := &models.Application{
			ID:        row.id,
			Date:      row.date,
			UserID:    row.user_id,
			FirstName: user.first_name,
			LastName:  user.last_name,
			Email:     user.email,
			Role:      models.Role(row.role),
		}
		apps = append(apps, app)
	}
	return apps, nil
}

func (r *applicationRepo) FindByUserID(userID int) (models.Applications, error) {
	query := `
	SELECT a.*, u.first_name, u.last_name, u.email
	FROM applications a
	JOIN users u ON u.id = a.user_id
	WHERE user_id = ?
	;`
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var apps []*models.Application
	for rows.Next() {
		var row applicationTableRow
		var user userTableRow
		var temp []uint8
		err := rows.Scan(&row.id, &temp, &row.user_id, &row.role, &user.first_name, &user.last_name, &user.email)
		if err != nil {
			return nil, err
		}
		row.date, err = time.Parse("2006-01-02 15:04:05", string(temp))
		if err != nil {
			return nil, err
		}
		app := &models.Application{
			ID:        row.id,
			Date:      row.date,
			UserID:    row.user_id,
			FirstName: user.first_name,
			LastName:  user.last_name,
			Email:     user.email,
			Role:      models.Role(row.role),
		}
		apps = append(apps, app)
	}
	return apps, nil
}

func (r *applicationRepo) FindByID(id int) (*models.Application, error) {
	query := `
	SELECT a.*, u.first_name, u.last_name, u.email
	FROM applications a
	JOIN users u ON u.id = a.user_id
	WHERE a.id = ?
	;`
	row := r.db.QueryRow(query, id)
	var app models.Application
	var user userTableRow
	var temp []uint8
	err := row.Scan(&app.ID, &temp, &app.UserID, &app.Role, &user.first_name, &user.last_name, &user.email)
	if err != nil {
		return nil, err
	}
	app.Date, err = time.Parse("2006-01-02 15:04:05", string(temp))
	if err != nil {
		return nil, err
	}
	app.FirstName = user.first_name
	app.LastName = user.last_name
	app.Email = user.email
	return &app, nil
}
