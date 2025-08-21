package user

type AuthRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginResponse struct {
	UserID     int                    `json:"user_id"`
	Email      string                 `json:"email"`
	Role       []string               `json:"role"`
	Permission []string               `json:"permission"`
	AccessToken  string           	  `json:"access_token"`
	RefreshToken string 			  `json:"refresh_token"`
}