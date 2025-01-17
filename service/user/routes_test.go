package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sakeththota/habit-tracker-go/types"
)

func TestUserServiceHandlers(t *testing.T) {
	t.Run("valid register user payload should succeed", func(t *testing.T) {
		userStore := &mockUserStore{userExists: false}
		handler := NewHandler(userStore)

		payload := types.RegisterUserPayload{
			Username: "testuser",
			Email:    "valid@email.com",
			Password: "testpassword",
		}
		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := gin.Default()
		router.POST("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)
		if rr.Code != http.StatusCreated {
			t.Errorf("expected status code %d, got %d", http.StatusCreated, rr.Code)
		}
	})

	t.Run("invalid register user payload should fail", func(t *testing.T) {
		userStore := &mockUserStore{userExists: false}
		handler := NewHandler(userStore)

		payload := types.RegisterUserPayload{
			Username: "testuser",
			Email:    "invalid",
			Password: "testpassword",
		}
		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := gin.Default()
		router.POST("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)
		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("user already exists should fail", func(t *testing.T) {
		userStore := &mockUserStore{userExists: true}
		handler := NewHandler(userStore)

		payload := types.RegisterUserPayload{
			Username: "testuser",
			Email:    "valid@email.com",
			Password: "testpassword",
		}
		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := gin.Default()
		router.POST("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)
		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusCreated, rr.Code)
		}
	})

	t.Run("invalid login user payload should fail", func(t *testing.T) {
		userStore := &mockUserStore{userExists: true}
		handler := NewHandler(userStore)

		payload := types.LoginUserPayload{
			Email:    "valid@email.com",
			Password: "short",
		}
		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := gin.Default()
		router.POST("/login", handler.handleLogin)
		router.ServeHTTP(rr, req)
		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("valid email and password should successfully log in", func(t *testing.T) {
		userStore := &mockUserStore{userExists: true}
		handler := NewHandler(userStore)

		payload := types.LoginUserPayload{
			Email:    "valid@email.com",
			Password: "testpassword",
		}
		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := gin.Default()
		router.POST("/login", handler.handleLogin)
		router.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusCreated, rr.Code)
		}
	})

	t.Run("invalid email should fail log in", func(t *testing.T) {
		userStore := &mockUserStore{userExists: false}
		handler := NewHandler(userStore)

		payload := types.LoginUserPayload{
			Email:    "invalid@email.com",
			Password: "testpassword",
		}
		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := gin.Default()
		router.POST("/login", handler.handleLogin)
		router.ServeHTTP(rr, req)
		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("invalid password should fail log in", func(t *testing.T) {
		userStore := &mockUserStore{userExists: true}
		handler := NewHandler(userStore)

		payload := types.LoginUserPayload{
			Email:    "invalid@email.com",
			Password: "wrongpassword",
		}
		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := gin.Default()
		router.POST("/login", handler.handleLogin)
		router.ServeHTTP(rr, req)
		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})
}

type mockUserStore struct {
	userExists bool
}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	if m.userExists {
		return &types.User{
			Username:     "testuser",
			Email:        "valid@email.com",
			PasswordHash: "$2a$10$zeo9dDF63T2huzccHZ7DSe2r6AcnlP491YyAnS0mvE3pXrcLyTP5C",
		}, nil
	}
	return nil, fmt.Errorf("User not found")
}

func (m *mockUserStore) GetUserById(id int) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) CreateUser(user types.User) error {
	return nil
}
