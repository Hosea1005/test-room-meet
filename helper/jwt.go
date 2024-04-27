package helper

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"strings"
	"time"
)

func GenerateJWT(signingKey []byte, userID string, email string, role string) (string, error) {
	// Buat klaim untuk token JWT
	expirationTime := time.Now().Add(10 * time.Minute)
	claims := jwt.MapClaims{
		"email": email,                 // Klaim `email`
		"role":  role,                  // Klaim `role`
		"exp":   expirationTime.Unix(), // Waktu kedaluwarsa token dalam format Unix
		"iss":   "room-meet",           // Nama aplikasi atau entitas yang menerbitkan token
		"sub":   email,                 // Subjek token (dapat menggunakan `email` sebagai subjek)
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
//func CheckJWT(tokenString string, signingKey []byte) error {
//	// Hapus prefiks "Bearer " jika ada
//	if strings.HasPrefix(tokenString, "Bearer ") {
//		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
//	}
//
//	// Parse token dengan klaim yang diharapkan (jwt.RegisteredClaims)
//	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
//		// Kembalikan kunci rahasia untuk verifikasi
//		return signingKey, nil
//	})
//
//	if err != nil {
//		return errors.New("invalid token")
//	}
//
//	// Periksa apakah token valid dan klaim tersedia
//	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
//		// Periksa apakah token sudah kedaluwarsa
//		if claims.ExpiresAt != nil && claims.ExpiresAt.Before(time.Now()) {
//			return errors.New("token has expired")
//		}
//	} else {
//		return errors.New("invalid claims or token")
//	}
//
//	// Jika token valid dan belum kedaluwarsa, tidak ada error
//	return nil
//}

func CheckJWT(tokenString string, signingKey []byte) (*jwt.MapClaims, error) {
	// Periksa apakah token kosong
	if tokenString == "" {
		return nil, errors.New("token is empty")
	}

	// Hapus prefiks "Bearer " jika ada
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	// Parse token
	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Kembalikan kunci rahasia untuk verifikasi
		return signingKey, nil
	})
	if err != nil {
		return nil, errors.New("invalid token")
	}

	// Periksa apakah token valid dan klaim tersedia
	if claims, ok := token.Claims.(*jwt.MapClaims); ok && token.Valid {
		// Periksa apakah token sudah kedaluwarsa
		if exp, ok := (*claims)["exp"].(int64); ok {
			if time.Now().Unix() > exp {
				return nil, errors.New("token is expired")
			}
		}
		return claims, nil
	} else {
		return nil, errors.New("invalid claims or token")
	}
}
