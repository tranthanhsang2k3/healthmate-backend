package services

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/sirupsen/logrus"
	"github.com/tranthanhsang2k3/healthmate-backend/auth-service/internal/models/user"
	"github.com/tranthanhsang2k3/healthmate-backend/auth-service/internal/repositories"
	"github.com/tranthanhsang2k3/healthmate-backend/auth-service/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterWithEmail(ctx context.Context, req user.RegisterRequest) error
	LoginWithEmail(ctx context.Context, req user.AuthRequest) (*user.LoginResponse, error)
}

type UserServiceImpl struct {
	userRepo repositories.UserRepository
	log      *logrus.Logger
}

func NewUserService(userRepo repositories.UserRepository, log *logrus.Logger) UserService {
	return &UserServiceImpl{
		log: 	log,
		userRepo: userRepo,
	}
}

func(s *UserServiceImpl) LoginWithEmail(ctx context.Context, req user.AuthRequest) (*user.LoginResponse, error){
	userEntity, err := s.userRepo.Login(ctx, req.Email)
	if err != nil {
		s.log.Error("Failed to login user: ", err)
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userEntity.Password), []byte(req.Password)); err != nil {
		s.log.Error("Password mismatch: ", err)
		return nil, err
	}

	var roles []string
	if err := json.Unmarshal(userEntity.Role, &roles); err != nil {
		s.log.Errorf("Failed to convert roles: %v", err)
		return nil, errors.New("invalid roles format")
	}

	var permissions []string
	if err := json.Unmarshal(userEntity.Permission, &permissions); err != nil {
		s.log.Error("Failed to convert permissions: ", err)
		return nil, err
	}

	accessToken, refreshToken, err := utils.GenerateJwtToken(userEntity.UserID, roles, permissions)
	if err != nil {
		s.log.Error("Failed to generate JWT tokens: ", err)
		return nil, err
	}

	return &user.LoginResponse{
		UserID: 	userEntity.UserID,
		Email:   	userEntity.Email,
		Role:   	roles,
		Permission: permissions,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func(s *UserServiceImpl) RegisterWithEmail(ctx context.Context, req user.RegisterRequest) error{
	if len(req.Role) == 0 || len(req.Permission) == 0 {
		s.log.Error("Roles and permissions cannot be empty")
		return utils.ErrorEmptyRoleOrPermission
	}

	userExists, _ := s.userRepo.Login(ctx, req.Email);
	if userExists != nil {
		s.log.Error("User already exists")
		return utils.ErrorUserAlreadyExists
	}

	userEntity := user.RegisterUserToEntity(req)
	if err := s.userRepo.Register(ctx, userEntity); err != nil {
		s.log.Error("Failed to register user: ", err)
		return err
	}

	return nil
}