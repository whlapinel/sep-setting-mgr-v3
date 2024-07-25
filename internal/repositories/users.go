package repositories

import (
	"database/sql"
	"log"
	"sep_setting_mgr/internal/domain/models"

	_ "github.com/go-sql-driver/mysql"
)

type userTableRow struct {
	id         int
	first_name string
	last_name  string
	email      string
	admin      bool
}

type userRepo struct {
	db *sql.DB
}

func NewUsersRepo(db *sql.DB) models.UserRepository {
	createUsersTable(db)
	return &userRepo{db: db}
}

func (ur *userRepo) Store(user *models.User) error {
	dbUser := convertToTable(user)
	log.Println("Adding user to database")
	_, err := ur.db.Exec(`INSERT INTO users (email, first_name, last_name, admin) VALUES (?, ?, ?, ?)`, dbUser.email, dbUser.first_name, dbUser.last_name, dbUser.admin)
	if err != nil {
		return err
	}
	return nil
}

func (ur *userRepo) Update(user *models.User) error {
	dbUser := convertToTable(user)
	log.Println("Updating user in database")
	_, err := ur.db.Exec(`
	UPDATE users 
	SET first_name = ?, last_name = ?, email = ?, admin = ? 
	WHERE id = ?`, dbUser.first_name, dbUser.last_name, dbUser.email, dbUser.admin, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (ur *userRepo) All() ([]*models.User, error) {
	var users []*models.User
	rows, err := ur.db.Query(`SELECT * FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var dbUser userTableRow
		err := rows.Scan(&dbUser.id, &dbUser.email, &dbUser.admin)
		if err != nil {
			return nil, err
		}
		user := convertFromTable(dbUser)
		users = append(users, user)
	}
	return users, nil
}

func (ur *userRepo) Find(email string) (*models.User, error) {
	var dbUser userTableRow
	ur.db.QueryRow(`SELECT * FROM users WHERE email = ?`, email).
		Scan(&dbUser.id, &dbUser.first_name, &dbUser.last_name, &dbUser.email, &dbUser.admin)
	user := convertFromTable(dbUser)
	return user, nil
}

func (ur *userRepo) FindByID(id int) (*models.User, error) {
	var dbUser userTableRow
	ur.db.QueryRow(`SELECT * FROM users WHERE id = ?`, id).
		Scan(&dbUser.id, &dbUser.email, &dbUser.admin)
	user := convertFromTable(dbUser)
	return user, nil
}

func (ur *userRepo) GetClasses(user *models.User) ([]*models.Class, error) {
	var classes []*models.Class

	return classes, nil
}

func (ur *userRepo) GetStudents(user *models.User) ([]*models.Student, error) {
	var classes []*models.Student

	return classes, nil
}

func createUsersTable(db *sql.DB) error {
	log.Println("Creating users table")
	result, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS 
	users (
		id int AUTO_INCREMENT PRIMARY KEY, 
		first_name VARCHAR(255) NOT NULL,
		last_name VARCHAR(255) NOT NULL,
		email VARCHAR(255) UNIQUE NOT NULL,
		admin BOOLEAN NOT NULL
		)`)
	if err != nil {
		return err
	}
	log.Println(result)

	return nil
}

func convertFromTable(dbUser userTableRow) *models.User {
	return &models.User{
		ID:        dbUser.id,
		FirstName: dbUser.first_name,
		LastName:  dbUser.last_name,
		Email:     dbUser.email,
		Admin:     dbUser.admin,
	}
}

func convertToTable(user *models.User) userTableRow {
	return userTableRow{
		first_name: user.FirstName,
		last_name:  user.LastName,
		email:      user.Email,
		admin:      user.Admin,
	}
}
