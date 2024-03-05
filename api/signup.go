package handler

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type SignupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func setupCORS(w *http.ResponseWriter, req *http.Request) {
    (*w).Header().Set("Access-Control-Allow-Origin", "https://cloudocr.vercel.app")
    (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	setupCORS(&w, r)
	if r.Method == "OPTIONS" {
        return
    }

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	
	var signupReq SignupRequest
	err := json.NewDecoder(r.Body).Decode(&signupReq)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	authData, err := json.Marshal(signupReq)
	if err != nil {
		http.Error(w, "Error encoding signup request", http.StatusInternalServerError)
		return
	}

	authServiceURL := "https://auth-microservice-ocr.vercel.app/api/signup"

	resp, err := http.Post(authServiceURL, "application/json", bytes.NewBuffer(authData))
	if err != nil {
		http.Error(w, "Failed to call auth service", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Inscription réussie"))
	} else {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Failed to read response from auth service", http.StatusInternalServerError)
			return
		}
		http.Error(w, string(body), resp.StatusCode)
	}
}