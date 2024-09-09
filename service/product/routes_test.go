package product

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/afallenhope/go-vendor/types"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type mockProductStore struct{}

func TestProductServiceHandlers(t *testing.T) {
	productStore := &mockProductStore{}
	handler := NewHandler(productStore)

	t.Run("should return a list of products", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/products", nil)

		if err != nil {
			t.Fatal()
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/products", handler.handleGetProducts).Methods(http.MethodGet)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status code %d, received %d", http.StatusOK, rr.Code)
		}
	})

	t.Run("should fail if productID is not a number", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/products/test", nil)
		if err != nil {
			t.Fatal()
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/products/{productID}", handler.handleGetProduct).Methods(http.MethodGet)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, received %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("should succeed the route if productID is a number", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/products/13", nil)
		if err != nil {
			t.Fatal()
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/products/{productID}", handler.handleGetProduct).Methods(http.MethodGet)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status code %d, received %d", http.StatusOK, rr.Code)
		}
	})

	t.Run("should fail if the product payload is invalid", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, "/products", nil)

		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/products", handler.handleCreateProduct).Methods(http.MethodPost)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}

	})

	t.Run("should create product", func(t *testing.T) {
		prodUUId, err := uuid.NewUUID()

		if err != nil {
			t.Fatal("could not create uuid for product.")
		}

		payload := types.CreateProductPayload{
			Name:        "testproduct",
			Price:       100,
			Image:       prodUUId,
			Description: "test image",
			Permissions: 0,
		}

		marshalled, err := json.Marshal(payload)

		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/products", handler.handleCreateProduct).Methods(http.MethodPost)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("expected status code %d, got %d", http.StatusCreated, rr.Code)
		}

	})

}

func (m *mockProductStore) GetProducts() ([]types.Product, error) {
	return nil, nil
}

func (m *mockProductStore) GetProductByID(id int) (*types.Product, error) {
	return nil, nil
}

func (m *mockProductStore) CreateProduct(product types.CreateProductPayload) error {
	return nil
}

func (m *mockProductStore) DeleteProductByID(id int) error {
	return nil
}
