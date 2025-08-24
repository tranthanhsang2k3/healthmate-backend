package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/tranthanhsang2k3/healthmate-backend/auth-service/internal/handlers"
	"github.com/tranthanhsang2k3/healthmate-backend/auth-service/internal/repositories"
	"github.com/tranthanhsang2k3/healthmate-backend/auth-service/internal/services"
)

func TestRegisterWithEmail(t *testing.T){
	gin.SetMode(gin.TestMode)

	db := setupTestDB()
	if db == nil {
		t.Fatal("Failed to connect to test database")
	}

	repo := repositories.NewUserRepository(db)
	service := services.NewUserService(repo, nil)
	handler := handlers.NewUserHandler(service)
	
	router := gin.New()
	router.POST("/register", handler.RegisterWithEmail())

	reqBody := `{"email":"test@example.com","password":"securepass","role":["user"],"permission":["read"]}`
	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBufferString(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "message")
}