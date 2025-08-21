package utils

import (
	"math/rand"
)

const digits = "0123456789"

func RandomOTP() string {
	var otp string
	for i := 0; i < 6; i++ {
		otp += string(digits[rand.Intn(len(digits))])
	}
	return otp
}