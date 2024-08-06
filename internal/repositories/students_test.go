package repositories

import (
	"log"
	"sep_setting_mgr/internal/domain/models"
	"sep_setting_mgr/internal/testutils"
	"testing"
)

func TestNewStudentsRepo(t *testing.T) {
	testutils.LoadEnvironment()
	db, err := InitializeDB(false)
	if err != nil {
		t.Errorf("InitializeDB() error = %v; want nil", err)
	}

	t.Run("should return a new instance of StudentsRepo", func(t *testing.T) {
		got := NewStudentsRepo(db)
		if got == nil {
			t.Error("NewStudentsRepo() = nil; want a new instance of StudentsRepo")
		}
	})
	t.Cleanup(func() {
		result, err := db.Exec("DROP TABLE students")
		if err != nil {
			t.Errorf("Drop table error = %v; want nil", err)
		}
		if result != nil {
			count, err := result.RowsAffected()
			if err != nil {
				t.Errorf("RowsAffected() error = %v; want nil", err)
			}
			log.Println(count)
		}
		db.Close()
	})
}

func TestStudentsRepo_Store(t *testing.T) {
	testutils.LoadEnvironment()
	db, err := InitializeDB(false)
	if err != nil {
		t.Errorf("InitializeDB() error = %v; want nil", err)
	}
	sr := NewStudentsRepo(db)
	student, err := models.NewStudent("John", "Doe", 1, false)
	if err != nil {
		t.Errorf("NewStudent() error = %v; want nil", err)
	}
	t.Run("should store a student in the database", func(t *testing.T) {
		err := sr.Store(student)
		if err != nil {
			t.Errorf("Store() error = %v; want nil", err)
		}
		if student.ID == 0 {
			t.Errorf("Store() id = %v; want non-zero", student.ID)
		}
	})
	t.Cleanup(func() {
		result, err := db.Exec("DROP TABLE students")
		if err != nil {
			t.Errorf("Drop table error = %v; want nil", err)
		}
		if result != nil {
			count, err := result.RowsAffected()
			if err != nil {
				t.Errorf("RowsAffected() error = %v; want nil", err)
			}
			log.Println(count)
		}
		db.Close()
	})
}

func TestStudentsRepo_All(t *testing.T) {
	testutils.LoadEnvironment()
	db, err := InitializeDB(false)
	if err != nil {
		t.Errorf("InitializeDB() error = %v; want nil", err)
	}
	sr := NewStudentsRepo(db)
	student, err := models.NewStudent("John", "Doe", 1, false)
	if err != nil {
		t.Errorf("NewStudent() error = %v; want nil", err)
	}
	err = sr.Store(student)
	if err != nil {
		t.Errorf("Store() error = %v; want nil", err)
	}
	t.Run("should return all students in the database", func(t *testing.T) {
		students, err := sr.AllInClass(1)
		if err != nil {
			t.Errorf("All() error = %v; want nil", err)
		}
		if len(students) == 0 {
			t.Errorf("All() students = %v; want non-zero", students)
		}
	})
	t.Cleanup(func() {
		result, err := db.Exec("DROP TABLE students")
		if err != nil {
			t.Errorf("Drop table error = %v; want nil", err)
		}
		if result != nil {
			count, err := result.RowsAffected()
			if err != nil {
				t.Errorf("RowsAffected() error = %v; want nil", err)
			}
			log.Println(count)
		}
		db.Close()
	})
}
