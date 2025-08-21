package user

import (
	"time"

	"gorm.io/datatypes"
)

type Users struct {
	UserID     int                    `gorm:"column:id;primaryKey"`
	Email      string                 `gorm:"column:email;unique"`
	Password   string                 `gorm:"column:password_hash"`
	IsActive   bool                   `gorm:"column:is_active"`
	CreatedAt  *time.Time             `gorm:"column:create_at"`
	Role       datatypes.JSON        `gorm:"column:roles;type:jsonb"`
	Permission datatypes.JSON        `gorm:"column:permissions;type:jsonb"`
	RefreshToken string               `gorm:"column:refresh_token"`
}

func(Users) TableName() string{
	return "users"
}