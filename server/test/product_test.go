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

func TestHandler_addProduct(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAdminCatalogServiceImpl, product entity.Product)

	tests := []struct {
		name                 string
		inputBody            string
		inputProduct         entity.Product
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "OK",
			inputBody: `{"category":"food", "name":"pizza", "price":"993", "description":"lalala", "color":"yellow"}`,
			inputProduct: entity.Product{
				Category:    "food",
				Name:        "pizza",
				Price:       "993",
				Description: "lalala",
				Color:       "yellow",
			},
			mockBehavior: func(r *mock_service.MockAdminCatalogServiceImpl, product entity.Product) {
				r.EXPECT().AddProduct(product).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":1}`,
		},
		{
			name:      "Invalid input",
			inputBody: `{"category":"food", "price":"993", "description":"lalala", "color":"yellow"}`,
			inputProduct: entity.Product{
				Category:    "food",
				Price:       "993",
				Description: "lalala",
				Color:       "yellow",
			},
			mockBehavior:         func(r *mock_service.MockAdminCatalogServiceImpl, product entity.Product) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"Invalid input"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockAdminCatalogServiceImpl(c)
			test.mockBehavior(repo, test.inputProduct)

			services := &service.Service{AdminCatalogServiceImpl: repo}
			h := handler.NewHandler(services)

			r := gin.New()
			r.POST("/catalog/edit/add", h.AddProduct)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/catalog/edit/add",
				bytes.NewBufferString(test.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}

func TestHandler_getAllProducts(t *testing.T) {
	type mockBehavior func(s *mock_service.MockProductServiceImpl)

	tests := []struct {
		name                 string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "OK",
			mockBehavior: func(r *mock_service.MockProductServiceImpl) {
				r.EXPECT().GetAllProducts().Return([]entity.Product{{1, "food", "tree", "993", "lalala", "white"}, {2, "cloth", "dress", "7", "lalala", "black"}}, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `[{"category":"food","name":"tree","price":"993","description":"lalala","color":"white"},{"category":"cloth","name":"dress","price":"7","description":"lalala","color":"black"}]`,
		},
		{
			name: "Invalid products",
			mockBehavior: func(r *mock_service.MockProductServiceImpl) {
				r.EXPECT().GetAllProducts().Return(nil, errors.New("invalid products"))
			},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid products"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockProductServiceImpl(c)
			test.mockBehavior(repo)

			services := &service.Service{ProductServiceImpl: repo}
			h := handler.NewHandler(services)

			r := gin.New()
			r.GET("/catalog/get/products/all", h.GetAllProducts)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/catalog/get/products/all", nil)

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}
