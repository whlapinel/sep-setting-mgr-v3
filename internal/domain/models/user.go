package models

type (
	User struct {
		ID        int
		FirstName string
		LastName  string
		Email     string
		Admin     bool
	}

	Teacher struct {
		User
		Classes  []*Class
		Students []*Student
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

func NewUser(firstName, lastName, email string) (*User, error) {
	return &User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}, nil
}
