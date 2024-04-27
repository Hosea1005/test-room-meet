package config

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

var JWT_KEY = []byte("ashdjqy9283409bsdklkg8hda01")

type JWTClaim struct {
	UserId uuid.UUID
	Email  string
	jwt.RegisteredClaims
}
