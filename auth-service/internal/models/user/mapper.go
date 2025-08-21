package user

func UsersToEntity(userDTO AuthRequest) *Users{
	return &Users{
		Email:   userDTO.Email,
		Password: userDTO.Password,
	}
}


