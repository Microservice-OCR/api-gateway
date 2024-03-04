package handler

import (
	"net/http"
	"io/ioutil"
)

// TokenResponse représente la structure de la réponse attendue contenant les tokens
// type TokenResponse struct {
// 	Tokens []models.Output `json:"tokens"`
// }

func GetAllTokensFromAuthService(w http.ResponseWriter, r *http.Request) {
	authServiceURL := "https://auth-theta-opal.vercel.app/api/token"

	// Création de la requête HTTP pour appeler le service d'authentification
	resp, err := http.Get(authServiceURL)
	if err != nil {
		http.Error(w, "Failed to connect to auth service", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response from auth service", http.StatusInternalServerError)
		return
	}

	// Décommentez et ajustez en fonction de la structure de réponse de votre service d'authentification
	// var tokenResponse TokenResponse
	// err = json.Unmarshal(body, &tokenResponse)
	// if err != nil {
	//     http.Error(w, "Error parsing JSON response from auth service", http.StatusInternalServerError)
	//     return
	// }

	// Réponse au client avec les tokens récupérés
	// Utilisez `tokenResponse` comme corps de la réponse si vous utilisez la structure de décodage JSON ci-dessus
	w.Header().Set("Content-Type", "application/json")
	w.Write(body) // Ou `jsonData` après le décodage et le re-codage en JSON si nécessaire
}
