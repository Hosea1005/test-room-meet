package helper

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"strings"
	"time"
)

var JWT_KEY = []byte("your-secret-key")

//
//func ParseJWTToken(tokenString string) (*config.JWTClaim, error) {
//	token, err := jwt.ParseWithClaims(tokenString, &config.JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
//		return config.JWT_KEY, nil
//	})
//
//	if err != nil {
//		return nil, err
//	}
//
//	claims, ok := token.Claims.(*config.JWTClaim)
//	if !ok || !token.Valid {
//		return nil, fmt.Errorf("Invalid token")
//	}
//
//	return claims, nil
//}
//func CheckToken(tokenString string) error {
//	// Periksa apakah token kosong
//	if tokenString == "" {
//		return errors.New("token is empty")
//	}
//
//	// Periksa apakah tokenString mengandung prefiks "Bearer " dan hapus prefiks tersebut
//	const bearerPrefix = "Bearer "
//	if len(tokenString) > len(bearerPrefix) && tokenString[:len(bearerPrefix)] == bearerPrefix {
//		tokenString = tokenString[len(bearerPrefix):]
//	}
//
//	// Parse token dengan klaim yang diharapkan
//	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
//		// Kembalikan kunci rahasia untuk verifikasi
//		return JWT_KEY, nil
//	})
//	if err != nil {
//		log.Println("Error parsing token:", err)
//		return errors.New("invalid token")
//	}
//
//	// Periksa apakah token valid dan klaim tersedia
//	claims, ok := token.Claims.(*jwt.RegisteredClaims)
//	if !ok || !token.Valid {
//		log.Println("Invalid claims or token")
//		return errors.New("invalid claims or token")
//	}
//
//	// Periksa apakah token sudah kadaluarsa
//	if claims.ExpiresAt != nil && claims.ExpiresAt.Before(time.Now()) {
//		log.Println("Token has expired")
//		return errors.New("token is expired")
//	}
//
//	// Token valid dan belum kadaluarsa
//	return nil
//}

func GenerateJWT(signingKey []byte, userID string, email string) (string, error) {
	// Buat klaim untuk token JWT
	claims := &jwt.RegisteredClaims{
		Issuer:    "room-meet",                                          // Ganti dengan nama aplikasi Anda
		Subject:   userID,                                               // Sesuaikan sesuai kebutuhan Anda
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 10)), // Token kedaluwarsa dalam 10 menit
	}

	// Buat token JWT dengan klaim dan metode tanda tangan HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Tandatangani token dengan kunci rahasia
	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Fungsi untuk memvalidasi token JWT
func CheckJWT(tokenString string, signingKey []byte) error {
	// Hapus prefiks "Bearer " jika ada
	if strings.HasPrefix(tokenString, "Bearer ") {
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	}

	// Parse token dengan klaim yang diharapkan (jwt.RegisteredClaims)
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Kembalikan kunci rahasia untuk verifikasi
		return signingKey, nil
	})

	if err != nil {
		return errors.New("invalid token")
	}

	// Periksa apakah token valid dan klaim tersedia
	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		// Periksa apakah token sudah kedaluwarsa
		if claims.ExpiresAt != nil && claims.ExpiresAt.Before(time.Now()) {
			return errors.New("token has expired")
		}
	} else {
		return errors.New("invalid claims or token")
	}

	// Jika token valid dan belum kedaluwarsa, tidak ada error
	return nil
}
