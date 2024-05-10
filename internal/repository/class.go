package repository

import (
	"errors"
	"fmt"
	"sync"

	"sep_setting_mgr/internal/domain"

	"github.com/google/uuid"
)

// MemoryClassRepository is an in-memory storage for Classs.
type MemoryClassRepository struct {
	mu      sync.Mutex
	classes map[uuid.UUID]domain.Class
}

// NewMemoryClassRepository creates a new instance of a memory-based Class repository.
func NewMemoryClassRepository() *MemoryClassRepository {
	return &MemoryClassRepository{
		classes: make(map[uuid.UUID]domain.Class),
	}
}

// Store implements the ClassRepository interface for MemoryClassRepository.
func (repo *MemoryClassRepository) Store(class *domain.Class) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	if _, exists := repo.classes[class.ID]; exists {
		return errors.New("class already exists")
	}

	repo.classes[class.ID] = *class
	return nil
}

func (repo *MemoryClassRepository) All() (*domain.Classes, error) {
	var classes domain.Classes
	for _, Class := range repo.classes {
		classes = append(classes, &Class)
	}
	fmt.Println(classes)
	return &classes, nil
}
