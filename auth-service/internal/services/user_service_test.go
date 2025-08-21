package services

import (
	"context"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tranthanhsang2k3/healthmate-backend/auth-service/internal/models/user"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type MockUserRepo struct {
	mock.Mock
}

func(m *MockUserRepo) Login(ctx context.Context, email string) (*user.Users, error) {
	args := m.Called(ctx, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*user.Users), args.Error(1)
}

func TestLoginWithEmailSuccess(t *testing.T) {
	// Setup mock repository and logger
	mockRepo := new(MockUserRepo)
	log := logrus.New()

	// Prepare request
	req := user.AuthRequest{
		Email:    "test@example.com",
		Password: "password",
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	assert.NoError(t, err)

	mockUser := &user.Users{
		UserID:   1,
		Email:    req.Email,
		Password: string(passwordHash),
		Role:     datatypes.JSON([]byte(`["admin"]`)),
		Permission: datatypes.JSON([]byte(`["read"]`)),
	}

	// Setup mock behavior
	mockRepo.On("Login", mock.Anything, req.Email).Return(mockUser, nil)

	// Create service and call method
	svc := NewUserService(mockRepo, log)
	resp, err := svc.LoginWithEmail(context.Background(), req)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	if resp != nil {
		assert.Equal(t, req.Email, resp.Email)
		assert.NotEmpty(t, resp.AccessToken)
	}
}

func TestLoginWithEmail_WrongPassword(t *testing.T){
	passwordHash, err := bcrypt.GenerateFromPassword([]byte("wrongpassword"), bcrypt.DefaultCost)
	assert.NoError(t, err)
	mockUser := &user.Users{
		UserID: 1,
		Email: "test@example.com",
		Password: string(passwordHash),
	}

	
	mockRepo := new(MockUserRepo)
	log := logrus.New()
	mockRepo.On("Login", mock.Anything, mockUser.Email).Return(mockUser, nil)

	// Create service and call method
	svc := NewUserService(mockRepo, log)

	req := user.AuthRequest{Email: "test@example.com", Password: "wrongpass"}
	resp, err := svc.LoginWithEmail(context.Background(), req)

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestLoginWithEmail_UserNotFound(t *testing.T) {
	mockRepo := new(MockUserRepo)
	mockRepo.On("Login", mock.Anything, "notfound@example.com").Return(nil, gorm.ErrRecordNotFound)

	svc := NewUserService(mockRepo, logrus.New())
	req := user.AuthRequest{Email: "notfound@example.com", Password: "123456"}
	resp, err := svc.LoginWithEmail(context.Background(), req)

	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestLoginWithEmail_GenerateJwtError(t *testing.T) {
	mockRepo := new(MockUserRepo)
	mockRepo.On("Login", mock.Anything, "test@example.com").Return(&user.Users{}, nil)
	log := logrus.New()
	svc := NewUserService(mockRepo, log)

	req := user.AuthRequest{Email: "test@example.com", Password: "123456"}
	resp, err := svc.LoginWithEmail(context.Background(), req)

	assert.Error(t, err)
	assert.Nil(t, resp)
}