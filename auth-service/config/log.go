package config

import (
	"os"

	"github.com/sirupsen/logrus"
)

func setupDevLogger() *logrus.Logger{
	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
		FullTimestamp: true,
	})
	return logger
}

func setupProdLogger() *logrus.Logger {
    logger := logrus.New()
    logger.SetOutput(os.Stdout)
    logger.SetLevel(logrus.InfoLevel)
    logger.SetFormatter(&logrus.JSONFormatter{
        TimestampFormat: "2006-01-02T15:04:05Z07:00",
    })
    return logger
}

func InitLogger(env string) *logrus.Logger {
    if env == "production" {
        return setupProdLogger()
    }
    return setupDevLogger()
}