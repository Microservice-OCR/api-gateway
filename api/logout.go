package handler

import (
	"bytes"
	"net/http"
)

func GatewayLogoutHandler(w http.ResponseWriter, r *http.Request) {
	// TODO : SUPPRIMER POUR PROD
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*") // or specify your domain
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	
	tokenHeader := r.Header.Get("Authorization")
	if tokenHeader == "" {
		http.Error(w, "No token provided", http.StatusBadRequest)
		return
	}

	authServiceURL := "https://auth-theta-opal.vercel.app/api/logout"
	client := &http.Client{}
	req, err := http.NewRequest("POST", authServiceURL, bytes.NewBufferString(""))
	if err != nil {
		http.Error(w, "Failed to create request to auth service", http.StatusInternalServerError)
		return
	}
	req.Header.Add("Authorization", tokenHeader)

	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Failed to request auth service", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Failed to logout", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "https://cloudocr.vercel.app/", http.StatusSeeOther)
}
