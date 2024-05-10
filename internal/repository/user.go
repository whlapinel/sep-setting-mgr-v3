package repository

import (
	"sep_setting_mgr/internal/domain"
	"sync"
)

type (
	UserRepository struct {
		mu    sync.Mutex
		users map[int]domain.User
	}
)
