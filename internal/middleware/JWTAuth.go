package middleware

import (
	"net/http"
	"strings"

	"github.com/KingGiveFiveCoins/Expenses/internal/utils"
	"github.com/gorilla/context"
)

func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Brak tokenu autoryzacyjnego", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Nieprawidłowy format tokenu", http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]
		claims, err := utils.ParseToken(tokenString)
		if err != nil {
			http.Error(w, "Nieprawidłowy token", http.StatusUnauthorized)
			return
		}

		// Dodajemy userID do kontekstu
		context.Set(r, "userID", claims.UserID)

		next.ServeHTTP(w, r)
	})
}
