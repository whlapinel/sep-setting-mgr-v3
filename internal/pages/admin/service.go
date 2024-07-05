package admin

import (
	domain "sep_setting_mgr/internal/domain/models"
	pages "sep_setting_mgr/internal/domain/pages"
)

type service struct {
	users domain.UserRepository
}

func NewService(users domain.UserRepository) pages.AdminService {
	return &service{users: users}
}

func (s service) ListUsers() ([]*domain.User, error) {
	return s.users.All()
}
