package domain

type (
	User struct {
		ID       int
		Username string
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
		Find(username string) (*User, error)
		GetClasses(user *User) ([]*Class, error)
		GetStudents(user *User) ([]*Student, error)
	}
)

func NewUser(username string, email string, password string, admin bool) (*User, error) {
	return &User{
		Username: username,
		Email:    email,
		Password: password,
		Admin:    admin,
	}, nil
}
