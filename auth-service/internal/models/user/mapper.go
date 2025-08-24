package user

import (
	"encoding/json"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/datatypes"
)

func UsersToEntity(userDTO AuthRequest) *Users {
	return &Users{
		Email:    userDTO.Email,
		Password: userDTO.Password,
	}
}

func RegisterUserToEntity(registerDTO RegisterRequest) *Users {
	password, _ := bcrypt.GenerateFromPassword([]byte(registerDTO.Password), bcrypt.DefaultCost)
	rolesBytes, _ := json.Marshal(registerDTO.Role)
	permissionsBytes, _ := json.Marshal(registerDTO.Permission)
	roles := datatypes.JSON(rolesBytes)
	permissions := datatypes.JSON(permissionsBytes)
	return &Users{
		Email:      registerDTO.Email,
		Password:   string(password),
		Role:       roles,
		Permission: permissions,
		IsActive:  false,
	}
}
