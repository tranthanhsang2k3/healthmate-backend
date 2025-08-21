package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tranthanhsang2k3/healthmate-backend/auth-service/internal/models/user"
	"github.com/tranthanhsang2k3/healthmate-backend/auth-service/internal/services"
	"github.com/tranthanhsang2k3/healthmate-backend/auth-service/utils"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// LoginWithEmail godoc
// @Summary Login with email
// @Description Login with email
// @Tags auth
// @Accept json
// @Produce json
// @Param request body user.AuthRequest true "Login request"
// @Success 200 {object} utils.Response{data=user.LoginResponse} "Login successful"
// @Failure 400 {object} utils.ErrorResponse "Invalid request"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /auth/login [post]
func (h *UserHandler) LoginWithEmail() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req user.AuthRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, utils.ErrorResponseFull(false, "Invalid request body: "+ err.Error()))
			return
		}

		resp, err := h.userService.LoginWithEmail(c.Request.Context(), req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, utils.ErrorResponseFull(false, "Login failed: "+ err.Error()))
			return
		}

		c.JSON(http.StatusOK, utils.ResponseFull(
			true,
			resp,
			"Login with email successfully",
		))
	}
}