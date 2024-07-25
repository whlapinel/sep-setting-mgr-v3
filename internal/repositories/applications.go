package repositories

import (
	"database/sql"
	"log"
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
	createApplicationsTable(db)
	return &applicationRepo{db: db}
}

func createApplicationsTable(db *sql.DB) {
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
		log.Fatalf("could not create applications table: %v", err)
	}
}

func (r *applicationRepo) Store(a *models.Application) error {
	query := `INSERT INTO applications (user_id, role) VALUES (?, ?);`
	_, err := r.db.Exec(query, a.UserID, a.Role)
	if err != nil {
		return err
	}
	return nil
}

func (r *applicationRepo) Delete(a *models.Application) error {
	query := `DELETE FROM applications WHERE id = ?;`
	_, err := r.db.Exec(query, a.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *applicationRepo) All() ([]*models.Application, error) {
	query := `
	SELECT a.*, u.* 
	FROM applications a
	JOIN users u ON u.id = a.user_id
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
		err := rows.Scan(&row.id, &row.date, &row.user_id, &row.role, &user.id, &user.first_name, &user.last_name, &user.email)
		if err != nil {
			return nil, err
		}
		role, err := models.GetRole(row.role)
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
			Role:      role,
		}
		apps = append(apps, app)
	}
	return apps, nil
}
