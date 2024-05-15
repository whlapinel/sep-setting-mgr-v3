package database

import (
	"database/sql"
	"log"
	"sep_setting_mgr/internal/domain"

	_ "github.com/go-sql-driver/mysql"
)

type userTable struct {
	id       int
	email    string
	password string
	admin    bool
}

type userRepo struct {
	db *sql.DB
}

func NewUsersRepo(db *sql.DB) domain.UserRepository {
	createUsersTable(db)
	return &userRepo{db: db}
}

func (ur *userRepo) Store(user *domain.User) error {
	dbUser := convertToTable(user)
	log.Println("Adding user to database")
	_, err := ur.db.Exec(`INSERT INTO users (email, password, admin) VALUES (?, ?, ?)`, dbUser.email, dbUser.password, dbUser.admin)
	if err != nil {
		return err
	}
	return nil
}

func (ur *userRepo) All() []*domain.User {
	var users []*domain.User

	return users
}

func (ur *userRepo) Find(email string) (*domain.User, error) {
	var dbUser userTable
	ur.db.QueryRow(`SELECT * FROM users WHERE email = ?`, email).
		Scan(&dbUser.id, &dbUser.email, &dbUser.password, &dbUser.admin)
	user := convertFromTable(dbUser)
	return user, nil
}

func (ur *userRepo) GetClasses(user *domain.User) ([]*domain.Class, error) {
	var classes []*domain.Class

	return classes, nil
}

func (ur *userRepo) GetStudents(user *domain.User) ([]*domain.Student, error) {
	var classes []*domain.Student

	return classes, nil
}

func createUsersTable(db *sql.DB) error {
	log.Println("Creating users table")
	result, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS 
	users (
		id int AUTO_INCREMENT PRIMARY KEY, 
		email VARCHAR(255) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL,
		admin BOOLEAN NOT NULL
		)`)
	if err != nil {
		return err
	}
	log.Println(result)

	return nil
}

func convertFromTable(dbUser userTable) *domain.User {
	return &domain.User{
		ID:       dbUser.id,
		Email:    dbUser.email,
		Password: dbUser.password,
		Admin:    dbUser.admin,
	}
}

func convertToTable(user *domain.User) userTable {
	return userTable{
		email:    user.Email,
		password: user.Password,
		admin:    user.Admin,
	}
}
