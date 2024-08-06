package models

type (
	User struct {
		ID        int
		FirstName string
		LastName  string
		Email     string
		Picture   string
		Admin     bool // must be approved for role
		Teacher   bool // must be approved for role
	}

	Teacher struct {
		User
		Classes  []*Class
		Students []*Student
	}

	UserRepository interface {
		Repository[User]
		DeleteAll
		Find(email string) (*User, error)
		FindByID(id int) (*User, error)
	}
)

func NewUser(firstName, lastName, email, picture string) (*User, error) {
	return &User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Picture:   picture,
	}, nil
}

func (u *User) FullName() string {
	return u.FirstName + " " + u.LastName
}
