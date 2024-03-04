package handler

import (
	"net/http"
	"io/ioutil"
)

func GetAllTokensFromAuthService(w http.ResponseWriter, r *http.Request) {
	authServiceURL := "https://auth-theta-opal.vercel.app/api/token"

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

	w.Header().Set("Content-Type", "application/json")
	w.Write(body) 
}
