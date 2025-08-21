package config

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(cf *Config, log *logrus.Logger) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		cf.DBHost,
		cf.DBUser,
		cf.DBPass,
		cf.DBName,
		cf.DBPort,
		cf.DBTimezone,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.WithError(err).Fatal("Không thể kết nối đến database")
        return

	}

	DB = db
	log.Info("Kết nối database thành công")
}

