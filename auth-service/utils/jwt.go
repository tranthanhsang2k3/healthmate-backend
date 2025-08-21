package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

var jwtSecret []byte

type JWTClaim struct {
	Permission []string `json:"permission"`
	Role       []string `json:"role"`
	UserID     int      `json:"id"`
	jwt.RegisteredClaims
}

func InitJWTSecret(secret string, log *logrus.Logger) {
	if secret == "" {
		log.Fatal("JWT secret is empty")
	}
	
	jwtSecret = []byte(secret)
}

func GenerateJwtToken(
	userID int,
	permissions []string,
	roles []string,
)(string, string, error){
	accessClaims := JWTClaim{
		Permission: permissions,
		Role:      roles,
		UserID:    userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString(jwtSecret)
	if err != nil {
		return "", "", err
	}

	refreshClaims := JWTClaim{
		Permission: permissions,
		Role:      roles,
		UserID:    userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(jwtSecret)
	if err != nil {
		return accessToken, "", err
	}

	return accessToken, refreshToken, nil
}

func ValidateJwtToken(tokenString string) (*JWTClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok || !token.Valid {
		return nil, jwt.ErrTokenInvalidClaims
	}

	return claims, nil
}