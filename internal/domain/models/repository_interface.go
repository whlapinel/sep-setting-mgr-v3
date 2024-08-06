package models

type Repository[T any] interface {
	Store(item *T) error
	Delete(int) error
	All() ([]*T, error)
	FindByID(int) (*T, error)
	Update(item *T) error
}

type DeleteAll interface {
	DeleteAll() error
}
