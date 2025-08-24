package config

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

var ctx = context.Background()

var client *redis.Client
var log *logrus.Logger

func InitRedisServer(conf *Config, logger *logrus.Logger) {
	addr := conf.RedisHost + ":" + conf.RedisPort
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})

	client = rdb
	log = logger
}

func SaveOTP(otp string, email string) error {
	err := client.Set(ctx, email, otp, 5*time.Minute).Err()
	if err != nil {
		log.WithFields(logrus.Fields{
			"email": email,
			"otp":   otp,
		}).Error("Failed to save OTP")
		return err
	}
	return nil
}

func GetOTP(email string) (string, error) {
	val, err := client.Get(ctx, email).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

