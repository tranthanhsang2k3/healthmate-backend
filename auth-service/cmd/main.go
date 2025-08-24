package main

import (
	_ "github.com/tranthanhsang2k3/healthmate-backend/auth-service/docs"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/tranthanhsang2k3/healthmate-backend/auth-service/config"
	"github.com/tranthanhsang2k3/healthmate-backend/auth-service/internal/handlers"
	"github.com/tranthanhsang2k3/healthmate-backend/auth-service/internal/repositories"
	"github.com/tranthanhsang2k3/healthmate-backend/auth-service/internal/services"
	"github.com/tranthanhsang2k3/healthmate-backend/auth-service/router"
)

// @title           Swagger Auth Service API
// @version         1.0
// @description     This is an auth server for cellar service.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.email  tranthanhsang.it.la@gmail.com

// @license.name  GNU
// @license.url   http://www.gnu.org/licenses/gpl-3.0.html

// @host      127.0.0.1:9000
// @BasePath  /api/v1
// @schemes http https
func main() {
	conf := config.LoadConfig()
	log :=  config.InitLogger(conf.AppConfig)

	config.ConnectDatabase(conf, log)
	config.InitRedisServer(conf, log)
	sendGridConfig := services.NewSendGridConfig(conf.SendGridAPIKey, conf.SMTPUsername)
	smtpConfig := services.NewSMTPConfig(conf.SMTPHost, conf.SMTPPort, conf.SMTPUsername, conf.SMTPPassword)
	services.NewSendOTPService(sendGridConfig, smtpConfig, log)

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	createUserHandler(r, log)
	r.Run(conf.GinHost+":"+conf.GinPort)
}

func createUserHandler(r *gin.Engine, log *logrus.Logger) {
	userRepository := repositories.NewUserRepository(config.DB)
	userService := services.NewUserService(userRepository, log)
	userHandler := handlers.NewUserHandler(userService)
	router.LoginRouter(r, userHandler)
}