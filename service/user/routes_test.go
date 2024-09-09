package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/afallenhope/go-vendor/types"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type mockUserStore struct{}

func TestUserServiceHandlers(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	t.Run("should fail if the user payload is invalid", func(t *testing.T) {
		uuid, _ := uuid.NewUUID()

		payload := types.RegisterUserPayload{
			UUID:     uuid,
			Password: "testpassword",
		}

		marshalled, err := json.Marshal(payload)
		if err != nil {
			t.Errorf("could not marshal data")
		}
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))

		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("should create user account", func(t *testing.T) {
		uuid, _ := uuid.NewUUID()

		payload := types.RegisterUserPayload{
			UUID:     uuid,
			Username: "testuser",
			Password: "testpassword",
		}

		marshalled, err := json.Marshal(payload)
		if err != nil {
			t.Errorf("could not marshal data")
		}
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))

		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("expected status code %d, got %d", http.StatusCreated, rr.Code)
		}
	})
}

// This for mocking responses.
// We don't want to actually touch the DB.
// cause... that would be bad kthx.
func (m *mockUserStore) GetUserByUsername(username string) (*types.User, error) {
	return nil, fmt.Errorf("user not found")
}

func (m *mockUserStore) GetUserByID(id uuid.UUID) (*types.User, error) {
	return nil, fmt.Errorf("user not found")
}

func (m *mockUserStore) GetUserByUUID(id uuid.UUID) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) CreateUser(user types.User) error { return nil }
