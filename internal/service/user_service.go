package service

import (
	"errors"
	// "fmt"
	"go_unit_test/internal/dto"
)

type UserRepository interface {
	GetUserRepository(id string) (dto.User, error)
}

type UserServ struct {
	repo UserRepository
}

func NewUserService(r UserRepository) *UserServ {
	return &UserServ{repo: r}
}

func (s *UserServ) GetUserService(id string) (dto.User, error) {

	if id == "" {
		return dto.User{}, errors.New("id is required")
	}
	user, err := s.repo.GetUserRepository(id)

	if err != nil {
		return dto.User{}, err
	}

	return user, nil
}
