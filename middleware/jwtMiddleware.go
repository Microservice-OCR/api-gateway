package middleware

import (
	"net/http"
	"strings"
	"time"
)

type ApiResponse struct {
	JWT_Token   string `json:"token"`
	ConnectedAt time.Time
}

var sessions = make(map[string]*ApiResponse)

func JwtMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// exécuter le gestionnaire suivant
		next(w, r)
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			// rediriger vers /login frontend
			http.Error(w, "Erreur lors de la lecture de la réponse", http.StatusUnauthorized)
			return
		}
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			// rediriger vers /login frontend
			http.Error(w, "Erreur lors de la lecture de la réponse", http.StatusUnauthorized)
			return
		}
		jwtToken := parts[1]

		// Vérifie si le JWT existe dans sessions
		session, exists := sessions[jwtToken]
		if !exists {
			// rediriger vers /login frontend
			http.Error(w, "Erreur lors de la lecture de la réponse", http.StatusUnauthorized)
			return
		}

		// Vérifie si plus de 10 minutes se sont écoulées depuis la connexion
		if time.Since(session.ConnectedAt) > 10*time.Minute {
			// modifier l'url pour mettre l'url du frontend
			// Si le temps actuel est égale au temps enregistrer dans la session
			http.Error(w, "Erreur lors de la lecture de la réponse", http.StatusUnauthorized)
			return
		}

		
	}
}
