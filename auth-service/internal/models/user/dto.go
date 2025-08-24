package user

type AuthRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type RegisterRequest struct {
	Email       string   `json:"email" binding:"required,email"`
	Password    string   `json:"password" binding:"required,min=6"`
	Role      	[]string `json:"roles" binding:"required"`
	Permission 	[]string `json:"permissions" binding:"required"`
}

type VerifyOTPRequest struct {
	Email string `json:"email" binding:"required,email"`
	OTP   string `json:"otp" binding:"required"`
}

type LoginResponse struct {
	UserID     int                    `json:"user_id"`
	Email      string                 `json:"email"`
	Role       []string               `json:"role"`
	Permission []string               `json:"permission"`
	AccessToken  string           	  `json:"access_token"`
	RefreshToken string 			  `json:"refresh_token"`
}