package repositories

import (
	"context"

	"github.com/tranthanhsang2k3/healthmate-backend/auth-service/internal/models/user"
	"gorm.io/gorm"
)

type UserRepository interface {
	Login(ctx context.Context, email string)(*user.Users, error)
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