package auth

import (
	"encoding/json"
	"net/http"
	"time"
	"api-gateway/img/models"
)

// GetAllTokensFromAuthService appelle le service d'authentification et récupère tous les tokens valides.
func GetAllTokensFromAuthService() ([]models.TokenInfo, error) {
	url := "https://auth-theta-opal.vercel.app/api/token"

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var tokens []models.TokenInfo
	if err := json.NewDecoder(resp.Body).Decode(&tokens); err != nil {
		return nil, err
	}

	return tokens, nil
}
