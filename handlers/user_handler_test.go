package handlers

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	customErrors "github.com/pamateus-henrique/infinitepay-firewatchers-api/errors"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/middlewares"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/models"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/validators"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserService is a mock implementation of the UserService interface
type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) Register(user *models.Register) error {
	args := m.Called(user)
	return args.Error(0)

}

func (m *MockUserService) Login(login *models.Login) (*models.User, error) {
	args := m.Called(login)
	return args.Get(0).(*models.User), args.Error(1)
}

func TestRegister(t *testing.T) {
	app := fiber.New(fiber.Config{
		ErrorHandler: middlewares.ErrorHandler,
	})
	mockService := new(MockUserService)
	handler := NewUserHandler(mockService)

	app.Post("/register", handler.Register)

	tests := []struct {
		name           string
		inputJSON      string
		mockBehavior   func()
		expectedStatus int
		expectedBody   string
	}{
		{
			name:      "Valid Registration",
			inputJSON: `{"name":"John Doe","email":"john@example.com","password":"password123"}`,
			mockBehavior: func() {
				mockService.On("Register", mock.Anything).Return(nil).Once()
			},
			expectedStatus: http.StatusCreated,
			expectedBody:   `{"message":"User registered successfully"}`,
		},
		{
			name:      "Invalid JSON Input",
			inputJSON: `{"name":"John Doe"`,  // Malformed JSON
			mockBehavior: func() {
				// No mock behavior needed
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":true,"message":"Invalid input format"}`,
		},
		{
			name:      "Validation Error",
			inputJSON: `{"name":"John Doe","email":"invalid-email","password":"short"}`,
			mockBehavior: func() {
				mockService.On("Register", mock.Anything).Return(&validators.ValidationError{Err: errors.New("validation failed")}).Once()
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":true,"message":"Validation failed","details":["validation failed"]}`,
		},
		{
			name:      "User Already Exists",
			inputJSON: `{"name":"John Doe","email":"john@example.com","password":"password123"}`,
			mockBehavior: func() {
				mockService.On("Register", mock.Anything).Return(errors.New("user already exists")).Once()
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"error":true,"message":"Internal Server Error"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset the mock between tests
			mockService.ExpectedCalls = nil
			mockService.Calls = nil

			tt.mockBehavior()

			req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBufferString(tt.inputJSON))
			req.Header.Set("Content-Type", "application/json")

			resp, err := app.Test(req)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			body, _ := io.ReadAll(resp.Body)
			assert.JSONEq(t, tt.expectedBody, string(body))

			mockService.AssertExpectations(t)
		})
	}
}

func TestLogin(t *testing.T) {
	app := fiber.New(fiber.Config{
		ErrorHandler: middlewares.ErrorHandler,
	})
	mockService := new(MockUserService)
	handler := NewUserHandler(mockService)

	app.Post("/login", handler.Login)

	tests := []struct {
		name           string
		inputJSON      string
		mockBehavior   func()
		expectedStatus int
		expectedBody   string
	}{
		{
			name:      "Valid Login",
			inputJSON: `{"email":"john@example.com","password":"password123"}`,
			mockBehavior: func() {
				mockService.On("Login", mock.AnythingOfType("*models.Login")).Return(&models.User{Name: "John Doe"}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"error":false,"msg":"Login successful"}`,
		},
		{
			name:      "Missing Password",
			inputJSON: `{"email":"john@example.com"}`,
			mockBehavior: func() {
				mockService.On("Login", mock.AnythingOfType("*models.Login")).Return((*models.User)(nil), &validators.ValidationError{Err: errors.New("Password is required")})
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":true,"message":"Validation failed","details":["Password is required"]}`,
		},
		{
			name:      "Invalid Credentials",
			inputJSON: `{"email":"john@example.com","password":"wrongpassword"}`,
			mockBehavior: func() {
				mockService.On("Login", mock.AnythingOfType("*models.Login")).Return((*models.User)(nil), &customErrors.AuthenticationError{Msg: "Invalid Email or password"})
			},
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   `{"error":true,"message":"Invalid Email or password"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset mock
			mockService.ExpectedCalls = nil
			mockService.Calls = nil

			// Set up mock behavior
			tt.mockBehavior()

			// Create request
			req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(tt.inputJSON))
			req.Header.Set("Content-Type", "application/json")

			// Perform request
			resp, err := app.Test(req)
			assert.NoError(t, err)

			// Check status code
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			// Check response body
			body, _ := io.ReadAll(resp.Body)
			assert.JSONEq(t, tt.expectedBody, string(body))

			// Assert that mock expectations were met
			mockService.AssertExpectations(t)
		})
	}
}