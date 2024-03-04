package middleware

import (
	"api-gateway/auth"
	"fmt"
	"net/http"
	"strings"
)

func JwtMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
			http.Error(w, fmt.Sprintf("Failed to get valid tokens : %s",err), http.StatusInternalServerError)
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