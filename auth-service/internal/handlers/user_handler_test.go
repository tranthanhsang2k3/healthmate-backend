package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tranthanhsang2k3/healthmate-backend/auth-service/internal/models/user"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) LoginWithEmail(ctx context.Context, req user.AuthRequest) (*user.LoginResponse, error) {
	args := m.Called(ctx, req)
	if conditions := args.Get(0); conditions == nil {
		return nil, args.Error(1)
		
	}
	return args.Get(0).(*user.LoginResponse), args.Error(1)
}

func TestLoginWithEmail_InvalidJSON(t *testing.T){
	gin.SetMode(gin.TestMode)

	mockService := new(MockUserService)
	h := NewUserHandler(mockService)

	router := gin.New()
	router.POST("/login", h.LoginWithEmail())

	reqBody := `{"email":"test@example.com","password":"password"`
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "message")
}

func TestLoginWithEmail_LoginFailed(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockSvc := new(MockUserService)
	h := NewUserHandler(mockSvc)

	router := gin.New()
	router.POST("/login", h.LoginWithEmail())

	reqData := user.AuthRequest{
		Email:    "test@example.com",
		Password: "wrongpass",
	}
	mockSvc.On("LoginWithEmail", mock.Anything, reqData).Return(nil, assert.AnError)

	body, _ := json.Marshal(reqData)
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "error")
}

func TestLoginWithEmail_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockSvc := new(MockUserService)
	h := NewUserHandler(mockSvc)

	router := gin.New()
	router.POST("/login", h.LoginWithEmail())

	reqData := user.AuthRequest{
		Email:    "test@example.com",
		Password: "123456",
	}
	respData := &user.LoginResponse{
		AccessToken: "jwt-token",
	}

	mockSvc.On("LoginWithEmail", mock.Anything, reqData).Return(respData, nil)

	body, _ := json.Marshal(reqData)
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Login with email successfully")
	assert.Contains(t, w.Body.String(), "jwt-token")
}