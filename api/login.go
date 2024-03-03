package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	// "time"
)

type ApiResponse struct {
    JWT_Token    string    `json:"token"`
    // ConnectedAt time.Time
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	AUTH_URI, ok := os.LookupEnv("AUTH_URI")
	if !ok {
		http.Error(w, "Auth URI not found", http.StatusInternalServerError)
	}

	// méthode POST
	if r.Method != "POST" {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	// Lire et décoder le corps de la requête entrante
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// Gère l'erreur si la lecture échoue
		http.Error(w, "Erreur lors de la lecture du corps de la requête", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	// Encodage des données d'identification au format JSON pour l'API d'authentification
	jsonBody, err := json.Marshal(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Création et envoi de la requête à l'API d'authentification
	client := &http.Client{}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/login",AUTH_URI), bytes.NewBuffer(jsonBody))
	if err != nil {
		http.Error(w, "Erreur lors de création de la requête", http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var result ApiResponse

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Erreur lors de la lecture de la réponse", http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(body, &result); err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de la lecture du corps de la réponse : %s\nresponse : %s",err,body), http.StatusInternalServerError)
	}
}