package middleware

import (
	"net/http"
	"strings"
	"webapp-backend/utils"

	"github.com/dgrijalva/jwt-go"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Dobavljanje Authorization header-a
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Autorizacija je obavezna", http.StatusUnauthorized)
			return
		}

		// Parsiranje tokena
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			// Proveravamo da li je koristio pravilan algoritam
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, http.ErrAbortHandler
			}
			return utils.GetJWTKey(), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Neispravan token", http.StatusUnauthorized)
			return
		}

		// Pozivamo sledeći handler
		next.ServeHTTP(w, r)
	})
}
