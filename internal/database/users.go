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
	_, err := ur.db.Exec(`INSERT INTO users () VALUES ()`, dbUser.username, dbUser.email, dbUser.password, dbUser.admin)
	if err != nil {
		return err
	}
	return nil
}

func (ur *userRepo) All() domain.Users {
	var users domain.Users

	return users
}

func (ur *userRepo) FindByID(classID string) (*domain.User, error) {
	var dbUser userTable

	user := convertFromTable(dbUser)
	return user, nil
}

func createUsersTable(db *sql.DB) error {
	log.Println("Creating users table")
	result, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS 
	users (
		id int AUTO_INCREMENT PRIMARY KEY, 
		first_name VARCHAR(255),
		last_name VARCHAR(255),
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
