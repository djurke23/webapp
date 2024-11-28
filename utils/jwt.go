package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Tajni ključ za potpisivanje JWT tokena
var jwtKey = []byte("tajni_kljuc")

// GenerateJWT kreira JWT token za određenog korisnika
func GenerateJWT(nickname string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"nickname": nickname,
		"exp":      time.Now().Add(24 * time.Hour).Unix(), // Token važi 24 sata
	})
	return token.SignedString(jwtKey)
}

// GetJWTKey vraća tajni ključ za potpisivanje JWT tokena
func GetJWTKey() []byte {
	return jwtKey
}
