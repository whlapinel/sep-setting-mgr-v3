package domain

type (
	Student struct {
		ID        int
		FirstName string
		LastName  string
		Teacher   User
	}
)

func NewStudent(firstName string, lastName string, teacher User) *Student {
	return &Student{
		FirstName: firstName,
		LastName:  lastName,
		Teacher:   teacher,
	}
}
