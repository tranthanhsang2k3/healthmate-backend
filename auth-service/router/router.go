package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tranthanhsang2k3/healthmate-backend/auth-service/internal/handlers"
)

func LoginRouter(r *gin.Engine, userHandler *handlers.UserHandler) {
	api := r.Group("/api/v1/auth")
	{
		api.POST("/login", userHandler.LoginWithEmail())
		api.POST("/register", userHandler.RegisterWithEmail())
	}
}