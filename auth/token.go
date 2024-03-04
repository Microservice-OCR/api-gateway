package auth

import (
	"api-gateway/img/models"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

// GetAllTokensFromAuthService appelle le service d'authentification et récupère tous les tokens valides.
func GetAllTokensFromAuthService() ([]models.TokenInfo, error) {
	AUTH_URI,ok := os.LookupEnv("AUTH_URI")
	if !ok {
		return nil, nil
	}
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(fmt.Sprintf("%s/token",AUTH_URI))
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
