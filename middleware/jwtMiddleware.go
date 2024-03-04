package middleware

import (
	"net/http"
	"strings"
	"api-gateway/auth"
)

func JwtMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
        // Votre logique de gestion d'upload
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		tokenHeader := r.Header.Get("Authorization")
		if tokenHeader == "" {
			http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(tokenHeader, "Bearer ")
		if token == "" {
			http.Error(w, "Token is missing", http.StatusUnauthorized)
			return
		}

		validTokens, err := auth.GetAllTokensFromAuthService()
		if err != nil {
			http.Error(w, "Failed to get valid tokens", http.StatusInternalServerError)
			return
		}

		tokenIsValid := false
		for _, tokenInfo := range validTokens {
			if token == tokenInfo.Token {
				tokenIsValid = true
				break
			}
		}

		if !tokenIsValid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}