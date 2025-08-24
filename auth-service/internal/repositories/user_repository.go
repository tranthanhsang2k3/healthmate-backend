package repositories

import (
	"context"

	"github.com/tranthanhsang2k3/healthmate-backend/auth-service/internal/models/user"
	"gorm.io/gorm"
)

type UserRepository interface {
	Login(ctx context.Context, email string)(*user.Users, error)
	Register(ctx context.Context, user *user.Users)(error)
}

type UserRepoImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepoImpl{
		db: db,
	}
}

func(r *UserRepoImpl) Login(ctx context.Context, email string)(*user.Users, error){
	var userEntity user.Users

	if err := r.db.WithContext(ctx).Model(&userEntity).Where("email = ?", email).First(&userEntity).Error; err != nil {
		return nil, err
	}

	return &userEntity, nil
}

func(r *UserRepoImpl) Register(ctx context.Context, user *user.Users)(error){
	if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
		return err
	}
	return nil
}