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
		Classes  Classes
		Students Students
	}

	Users []*User

	UserService interface {
		RegisterNewUser(username string, email string, password string, admin bool) (*User, error)
	}

	UserRepository interface {
		Store(*User) error
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
