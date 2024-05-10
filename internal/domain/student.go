package domain

type (
	Student struct {
		ID        int
		FirstName string
		LastName  string
		Teacher   User
	}

	Students []*Student
)

func NewStudents() (*Students, error) {
	return &Students{}, nil
}

func NewStudent(firstName string, lastName string, teacher User) *Student {
	return &Student{
		FirstName: firstName,
		LastName:  lastName,
		Teacher:   teacher,
	}
}
