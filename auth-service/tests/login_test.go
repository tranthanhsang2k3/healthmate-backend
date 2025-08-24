package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"gorm.io/datatypes"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/tranthanhsang2k3/healthmate-backend/auth-service/internal/handlers"
	"github.com/tranthanhsang2k3/healthmate-backend/auth-service/internal/models/user"
	"github.com/tranthanhsang2k3/healthmate-backend/auth-service/internal/repositories"
	"github.com/tranthanhsang2k3/healthmate-backend/auth-service/internal/services"
	"golang.org/x/crypto/bcrypt"
)

func TestLoginWithEmail_Integration(t *testing.T) {
	gin.SetMode(gin.TestMode)

	db := setupTestDB()
	if db == nil {
		t.Fatal("Failed to connect to test database")
	}

	// Seed dữ liệu user
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	db.Create(&user.Users{
		Email:        "test@example.com",
		Password:     string(passwordHash),
		IsActive:     false,
		Role:         datatypes.JSON([]byte(`["admin"]`)),
		Permission:   datatypes.JSON([]byte(`["read"]`)),
		RefreshToken: "",
	})

	// Khởi tạo repo, service, handler thật
	repo := repositories.NewUserRepository(db)
	service := services.NewUserService(repo, logrus.New())
	handler := handlers.NewUserHandler(service)

	// Setup router
	router := gin.New()
	router.POST("/login", handler.LoginWithEmail())

	// Gửi request thật
	reqBody := `{"email":"test@example.com","password":"123456"}`
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	// Kiểm tra kết quả
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Login with email successfully")
}
