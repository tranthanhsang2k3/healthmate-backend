package config

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

var client *redis.Client

func InitRedisServer(conf *Config) {
	addr := conf.RedisHost + ":" + conf.RedisPort
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})

	client = rdb
}

func SaveOTP(otp string, email string) error {
	err := client.Set(ctx, email, otp, 5*time.Minute).Err()
	if err != nil {
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

