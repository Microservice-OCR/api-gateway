package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	// "api-gateway/middleware"
	// "time"
)

type AuthInput struct {
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// TODO : SUPPRIMER POUR PROD
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*") // or specify your domain
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	if r.Method == "OPTIONS" {
        w.WriteHeader(http.StatusOK)
        return
    }
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
	
	input := AuthInput{}

	// Décodage du JSON dans la structure IInput
	err = json.Unmarshal(body, &input)
	if err != nil {
		http.Error(w, "Erreur lors de la conversion du JSON", http.StatusInternalServerError)
		return
	}

	// Convertir les données du corps en JSON
	jsonData, err := json.Marshal(input)
	if err != nil {
		http.Error(w, "Erreur lors de la conversion du JSON", http.StatusInternalServerError)
		return 
	}

	// Création et envoi de la requête à l'API d'authentification
	client := &http.Client{}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/login",AUTH_URI), bytes.NewBuffer(jsonData))
	if err != nil {
		http.Error(w, "Erreur lors de création de la requête", http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Erreur lors de la requête", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Erreur lors de la lecture de la réponse", http.StatusInternalServerError)
		return
	}
	
	// var result middleware.ApiResponse
	// if err := json.Unmarshal(body, &result); err != nil {
	// 	http.Error(w, fmt.Sprintf("Erreur lors de la lecture du corps de la réponse : %s\nresponse : %s",err,body), http.StatusInternalServerError)
	// }
	w.Write(body)
}