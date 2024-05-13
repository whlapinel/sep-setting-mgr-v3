package database

import (
	"database/sql"
	"log"
	"sep_setting_mgr/internal/domain"

	_ "github.com/go-sql-driver/mysql"
)

type userTable struct {
	id       int
	username string
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
	log.Println("Adding class to database")
	_, err := ur.db.Exec(`INSERT INTO users (username, email, password, admin) VALUES (?, ?, ?, ?)`, dbUser.username, dbUser.email, dbUser.password, dbUser.admin)
	if err != nil {
		return err
	}
	return nil
}

func (ur *userRepo) All() []*domain.User {
	var users []*domain.User

	return users
}

func (ur *userRepo) Find(username string) (*domain.User, error) {
	var dbUser userTable

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
		username VARCHAR(255),
		email VARCHAR(255),
		password VARCHAR(255),
		admin BOOLEAN
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
		Username: dbUser.username,
		Email:    dbUser.email,
		Password: dbUser.password,
		Admin:    dbUser.admin,
	}
}

func convertToTable(user *domain.User) userTable {
	return userTable{
		username: user.Username,
		email:    user.Email,
		password: user.Password,
		admin:    user.Admin,
	}
}
