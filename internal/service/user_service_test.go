package service

import (
	"errors"
	"go_unit_test/internal/dto"

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetUserRepository(id string) (dto.User, error) {
	args := m.Called(id)

	return args.Get(0).(dto.User), args.Error(1)
}

func TestGetUserService(t *testing.T) {

	tests := []struct {
		name        string
		inputId     string
		mockReturn  dto.User
		mockError   error
		expectError bool
		expectName  string
		expectEmail string
	}{
		{
			name:        "success - user found",
			inputId:     "1",
			mockReturn:  dto.User{Name: "João", Email: "joao@gmail.com"},
			mockError:   nil,
			expectError: false,
			expectName:  "João",
			expectEmail: "joao@gmail.com",
		},
		{
			name:        "empty id - service validation error",
			inputId:     "",
			mockReturn:  dto.User{},
			mockError:   nil,
			expectError: true,
			expectName:  "",
			expectEmail: "",
		},
		{
			name:        "repository return error - service validation error",
			inputId:     "99",
			mockReturn:  dto.User{},
			mockError:   errors.New("User does not exist"),
			expectError: true,
			expectName:  "",
			expectEmail: "",
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(MockUserRepository)

			if tc.inputId != "" {
				mockRepo.On("GetUserRepository", tc.inputId).Return(tc.mockReturn, tc.mockError)
			}

			svc := NewUserService(mockRepo)

			user, err := svc.GetUserService(tc.inputId)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectName, user.Name)
				assert.Equal(t, tc.expectEmail, user.Email)
			}

			mockRepo.AssertExpectations(t)
		})

	}

}
