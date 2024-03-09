package test

import (
	"bytes"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
	"go.uber.org/mock/gomock"
	"net/http/httptest"
	"server/entity"
	"server/handler"
	"server/service"
	mock_service "server/service/mocks"
	"testing"
)

func TestHandler_signUp(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAuthServiceImpl, user entity.User)

	tests := []struct {
		name                 string
		inputBody            string
		inputUser            entity.User
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "OK",
			inputBody: `{"username":"testName", "password":"testPassword"}`,
			inputUser: entity.User{
				Username: "testName",
				Password: "testPassword",
			},
			mockBehavior: func(r *mock_service.MockAuthServiceImpl, user entity.User) {
				r.EXPECT().SignUp(user).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":1}`,
		},
		{
			name:                 "Invalid input",
			inputBody:            `{"username":"testName"}`,
			inputUser:            entity.User{},
			mockBehavior:         func(r *mock_service.MockAuthServiceImpl, user entity.User) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid input body"}`,
		},
		{
			name:      "Duplicate username value",
			inputBody: `{"username":"testName", "password":"testPassword"}`,
			inputUser: entity.User{
				Username: "testName",
				Password: "testPassword",
			},
			mockBehavior: func(r *mock_service.MockAuthServiceImpl, user entity.User) {
				r.EXPECT().SignUp(user).Return(0, errors.New("Duplicate username value"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"duplicate username value"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockAuthServiceImpl(c)
			test.mockBehavior(repo, test.inputUser)

			services := &service.Service{AuthServiceImpl: repo}
			h := handler.NewHandler(services)

			// Init Endpoint
			r := gin.New()
			r.POST("/auth/signup", h.SignUp)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/auth/signup",
				bytes.NewBufferString(test.inputBody))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}

func TestHandler_logIn(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAuthServiceImpl, user entity.User)

	tests := []struct {
		name                 string
		inputBody            string
		inputUser            entity.User
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "OK",
			inputBody: `{"username":"admin", "password":"123"}`,
			inputUser: entity.User{
				Username: "admin",
				Password: "123",
			},
			mockBehavior: func(r *mock_service.MockAuthServiceImpl, user entity.User) {
				r.EXPECT().GenerateToken(user.Username, user.Password).Return("generatedToken", nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"token":"generatedToken"}`,
		},
		{
			name:                 "Invalid input",
			inputBody:            `{"username":"admin"}`,
			inputUser:            entity.User{},
			mockBehavior:         func(r *mock_service.MockAuthServiceImpl, user entity.User) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid input body"}`,
		},
		{
			name:      "Invalid username or password",
			inputBody: `{"username":"admin", "password":"lol"}`,
			inputUser: entity.User{
				Username: "admin",
				Password: "lol",
			},
			mockBehavior: func(r *mock_service.MockAuthServiceImpl, user entity.User) {
				r.EXPECT().GenerateToken(user.Username, user.Password).Return("", errors.New("cant generate token"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"invalid username or password"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockAuthServiceImpl(c)
			test.mockBehavior(repo, test.inputUser)

			services := &service.Service{AuthServiceImpl: repo}
			h := handler.NewHandler(services)

			r := gin.New()
			r.POST("/auth/login", h.Login)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/auth/login",
				bytes.NewBufferString(test.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}
