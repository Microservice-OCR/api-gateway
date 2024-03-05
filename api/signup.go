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

type SignupResponse struct {
	Message string `json:"message"`
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	var req SignupRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	authData, err := json.Marshal(req)
	if err != nil {
		http.Error(w, "Error encoding auth request", http.StatusInternalServerError)
		return
	}

	authServiceURL := "https://auth-microservice-ocr.vercel.app/api/signup"

	resp, err := http.Post(authServiceURL, "application/json", bytes.NewBuffer(authData))
	if err != nil {
		http.Error(w, "Failed to call auth service", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read auth service response", http.StatusInternalServerError)
		return
	}

	var authResp SignupResponse
	err = json.Unmarshal(body, &authResp)
	if err != nil {
		http.Error(w, "Error decoding auth response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(body)
}