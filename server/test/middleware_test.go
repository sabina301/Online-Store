package test

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
	"go.uber.org/mock/gomock"
	"net/http/httptest"
	handler "server/handler"
	"server/service"
	"server/service/mocks"
	"testing"
)

func TestHandler_middleware(t *testing.T) {
	type mockBehavior func(r *mock_service.MockAuthServiceImpl, token string)
	testTable := []struct {
		name                 string
		headerName           string
		headerValue          string
		token                string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:        "OK",
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			mockBehavior: func(r *mock_service.MockAuthServiceImpl, token string) {
				r.EXPECT().ParseToken(token).Return(1, "user", nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "1",
		},
		{
			name:                 "Invalid Header Name",
			headerName:           "",
			headerValue:          "Bearer token",
			token:                "token",
			mockBehavior:         func(r *mock_service.MockAuthServiceImpl, token string) {},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"empty auth header name or httpOnlyCookie"}`,
		},
		{
			name:                 "Invalid Header Value",
			headerName:           "Authorization",
			headerValue:          "Bearr token",
			token:                "token",
			mockBehavior:         func(r *mock_service.MockAuthServiceImpl, token string) {},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"invalid header value"}`,
		},
		{
			name:                 "Empty Token",
			headerName:           "Authorization",
			headerValue:          "Bearer ",
			token:                "token",
			mockBehavior:         func(r *mock_service.MockAuthServiceImpl, token string) {},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"token is empty"}`,
		},
		{
			name:        "Parse Error",
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			mockBehavior: func(r *mock_service.MockAuthServiceImpl, token string) {
				r.EXPECT().ParseToken(token).Return(0, "user", errors.New("invalid token"))
			},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"invalid token"}`,
		},
	}
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {

			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockAuthServiceImpl(c)
			test.mockBehavior(repo, test.token)

			services := &service.Service{AuthServiceImpl: repo}
			h := handler.NewHandler(services)

			r := gin.New()
			r.GET("/identity", h.UserIdentity, func(c *gin.Context) {
				id, _ := c.Get("userId")
				c.String(200, "%d", id)
			})

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/identity", nil)
			req.Header.Set(test.headerName, test.headerValue)

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}
