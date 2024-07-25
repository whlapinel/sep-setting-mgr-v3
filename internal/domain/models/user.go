package models

type (
	User struct {
		ID        int
		FirstName string
		LastName  string
		Email     string
		Admin     bool // must be approved for role
		Teacher   bool // must be approved for role
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
