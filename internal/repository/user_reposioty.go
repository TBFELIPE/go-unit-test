package repository

import (
	"errors"
	"go_unit_test/internal/dto"
)

type RepositoryUser struct{}

type UserInTable struct {
	Id    string
	Name  string
	Email string
}

func (r *RepositoryUser) GetUserRepository(id string) (dto.User, error) {

	var userTable = []UserInTable{
		{Id: "1", Name: "João", Email: "jose@gmail.com"},
		{Id: "2", Name: "Maria", Email: "maria@test.com"},
		{Id: "3", Name: "José", Email: "jose@test.com"},
		{Id: "4", Name: "Pedro", Email: "pedro@test.com"},
	}

	if id == "" {
		return dto.User{}, errors.New("id is reequired")
	}

	for _, u := range userTable {

		if id == u.Id {
			return dto.User{
				Name:  u.Name,
				Email: u.Email}, nil
		}
	}

	return dto.User{}, errors.New("User does not exist")
}
