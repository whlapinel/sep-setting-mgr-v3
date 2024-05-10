package repository

import (
	"errors"
	"fmt"
	"log"
	"sync"

	"sep_setting_mgr/internal/domain"
)

// MemoryTestEventRepository is an in-memory storage for TestEvents.
type MemoryTestEventRepository struct {
	mu     sync.Mutex
	events map[int]domain.TestEvent
}

// NewMemoryTestEventRepository creates a new instance of a memory-based TestEvent repository.
func NewMemoryTestEventRepository() *MemoryTestEventRepository {
	return &MemoryTestEventRepository{
		events: make(map[int]domain.TestEvent),
	}
}

// Store implements the TestEventRepository interface for MemoryTestEventRepository.
func (repo *MemoryTestEventRepository) Store(event *domain.TestEvent) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.assignID(event)
	if _, exists := repo.events[event.ID]; exists {
		return errors.New("event already exists")
	}

	repo.events[event.ID] = *event
	return nil
}

func (repo *MemoryTestEventRepository) FindAll() (*domain.TestEvents, error) {
	var testEventSlice domain.TestEvents
	for _, testEvent := range repo.events {
		testEventSlice = append(testEventSlice, &testEvent)
	}
	fmt.Println(testEventSlice)
	return &testEventSlice, nil
}

func (repo *MemoryTestEventRepository) assignID(event *domain.TestEvent) {
	// get max ID
	getMax := func(eventMap map[int]domain.TestEvent) int {
		max := 0
		for key := range eventMap {
			log.Println("current key: ", key)
			if key > max {
				max = key
			}
		}
		log.Println("number of keys: ", len(eventMap))
		log.Println("max: ", max)
		return max
	}
	maxID := getMax(repo.events)

	// assign number to event.ID
	event.ID = maxID + 1

}
