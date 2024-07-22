package models

type (
	User struct {
		ID       int
		Email    string
		Admin    bool
		Password string
	}

	Teacher struct {
		User
		Classes  []*Class
		Students []*Student
	}

	UserService interface {
		NewUser(username string, email string, password string, admin bool) (*User, error)
	}

	UserRepository interface {
		Store(*User) error
		Update(*User) error
		Find(email string) (*User, error)
		FindByID(id int) (*User, error)
		GetClasses(user *User) ([]*Class, error)
		GetStudents(user *User) ([]*Student, error)
		All() ([]*User, error)
	}
)

func NewUser(email string, password string) (*User, error) {
	return &User{
		Email:    email,
		Password: password,
		Admin:    false,
	}, nil
}
