package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserRepository(t *testing.T) {

	repo := &RepositoryUser{}

	// table-driven test have: true case, wrong case and empty case
	tests := []struct {
		name        string
		inputId     string
		expectError bool
		expectName  string
	}{
		{
			name:        "user exists - id 1",
			inputId:     "1",
			expectError: false,
			expectName:  "João",
		},
		{
			name:        "user does not exists - id 99",
			inputId:     "99",
			expectError: true,
			expectName:  "",
		},
		{
			name:        "empty id",
			inputId:     "",
			expectError: true,
			expectName:  "",
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			user, err := repo.GetUserRepository(tc.inputId) // calling the real GetUserRepository method

			if tc.expectError {
				assert.Error(t, err)

				assert.Empty(t, tc.expectName, user.Name)
			} else {
				assert.NoError(t, err)

				assert.Equal(t, tc.expectName, user.Name)
			}
		})
	}

}
