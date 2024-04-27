package helper

import (
	"errors"
	"fmt"
	"room-meet/internal/config"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var JWT_KEY = []byte("your-secret-key")

func ParseJWTToken(tokenString string) (*config.JWTClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &config.JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		return config.JWT_KEY, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*config.JWTClaim)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("Invalid token")
	}

	return claims, nil
}
func CheckToken(tokenString string) error {
	// Periksa apakah token kosong
	if tokenString == "" {
		return errors.New("token is empty")
	}

	// Parse token
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Kembalikan kunci rahasia untuk verifikasi
		return JWT_KEY, nil
	})
	if err != nil {
		return errors.New("invalid token")
	}

	// Periksa apakah token valid dan klaim tersedia
	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		// Periksa apakah token sudah kadaluarsa
		if claims.ExpiresAt != nil && claims.ExpiresAt.Before(time.Now()) {
			return errors.New("token is expired")
		}
	} else {
		return errors.New("invalid claims or token")
	}

	// Token valid dan belum kadaluarsa
	return nil
}
