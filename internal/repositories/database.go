package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func InitializeDB(production bool) (*sql.DB, error) {

	cfg := mysql.Config{
		User:                 os.Getenv("MARIADB_USER"),
		Passwd:               os.Getenv("MARIADB_PASSWORD"),
		Net:                  "tcp",
		Addr:                 os.Getenv("MARIADB_HOST") + ":3306", // Change this to the service name and port in your docker-compose.yml
		DBName:               os.Getenv("MARIADB_DATABASE"),
		AllowNativePasswords: true,
	}

	var db *sql.DB
	var err error

	for i := 0; i < 5; i++ {
		fmt.Println("Connecting to the database...")
		db, err = sql.Open("mysql", cfg.FormatDSN())
		if err != nil {
			log.Printf("Failed to open database: %v", err)
			time.Sleep(5 * time.Second)
			continue
		}

		err = db.Ping()
		if err != nil {
			log.Printf("Failed to ping database: %v", err)
			db.Close()
			time.Sleep(5 * time.Second)
			continue
		}

		log.Println("Connected to database.")
		return db, nil
	}
	log.Fatal("Failed to connect to the database after 5 attempts. Exiting...")
	return nil, nil
}

func ClearDatabase(db *sql.DB) error {
	log.Println("Clearing database...")
	log.Println("Dropping applications")
	_, err := db.Exec("DROP TABLE IF EXISTS applications")
	if err != nil {
		return err
	}
	log.Println("Dropping assignments")
	_, err = db.Exec("DROP TABLE IF EXISTS assignments")
	if err != nil {
		return err
	}
	log.Println("Dropping students")
	_, err = db.Exec("DROP TABLE IF EXISTS students")
	if err != nil {
		return err
	}
	log.Println("Dropping test_events")
	_, err = db.Exec("DROP TABLE IF EXISTS test_events")
	if err != nil {
		return err
	}
	log.Println("Dropping classes")
	_, err = db.Exec("DROP TABLE IF EXISTS classes")
	if err != nil {
		return err
	}
	log.Println("Dropping rooms")
	_, err = db.Exec("DROP TABLE IF EXISTS rooms")
	if err != nil {
		return err
	}
	log.Println("Dropping users")
	_, err = db.Exec("DROP TABLE IF EXISTS users")
	if err != nil {
		return err
	}
	return nil
}
