package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"go_unit_test/internal/dto"

	"testing"

	"github.com/gin-gonic/gin"

	"net/http/httptest"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

// Implement a GetUserService, method of UserService interface in controller
func (m *MockUserService) GetUserService(id string) (dto.User, error) {
	args := m.Called(id) //

	return args.Get(0).(dto.User), args.Error(1)
}

func TestGetUserHandler(t *testing.T) {

	gin.SetMode(gin.TestMode)

	// testing struct table
	tests := []struct {
		name           string
		requestBody    dto.Request
		mockReturn     dto.User
		mockError      error
		expectedStatus int
	}{
		{
			name:           "Success - user found",
			requestBody:    dto.Request{Id: "1"},
			mockReturn:     dto.User{Name: "João", Email: "joao@gmail.com"},
			mockError:      nil,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Success - user does not found - 404",
			requestBody:    dto.Request{Id: "99"},
			mockReturn:     dto.User{},
			mockError:      errors.New("User does not exist"),
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "invalid body - without id in json - 404",
			requestBody:    dto.Request{Id: ""},
			mockReturn:     dto.User{},
			mockError:      errors.New("id is required"),
			expectedStatus: http.StatusNotFound,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {

			mockSvc := new(MockUserService)

			mockSvc.On("GetUserService", tc.requestBody.Id).Return(tc.mockReturn, tc.mockError)

			handler := NewUserHandler(mockSvc)

			router := gin.New()

			router.POST("/get-user", handler.GetUserHandler)

			bodyBytes, _ := json.Marshal(tc.requestBody)

			req := httptest.NewRequest(
				http.MethodPost,
				"/get-user",
				bytes.NewBuffer(bodyBytes),
			)

			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedStatus, w.Code)

			mockSvc.AssertExpectations(t)
		})
	}

}
