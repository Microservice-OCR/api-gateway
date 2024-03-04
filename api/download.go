package handler

import (
	"api-gateway/img"
	"api-gateway/middleware"
	"encoding/json"
	"fmt"
	"net/http"
)

func ImageDownloadHandler(w http.ResponseWriter, r *http.Request) {
	middleware.JwtMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO : SUPPRIMER POUR PROD
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*") // or specify your domain
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
	
		if r.Method == "GET" {
			// Récupération de l'id de l'image
			imageId := r.URL.Query().Get("id")
	
			if imageId != "" {
				// Appel de l'API saveIMG
				imgData, err := img.GetImageFromId(imageId)
				if err != nil {
					http.Error(w, fmt.Sprintf("Erreur lors de l'appel de l'API saveIMG : %s",err), http.StatusInternalServerError)
					return
				}
	
				jsonData, err := json.Marshal(imgData)
				if err != nil {
					http.Error(w, "Erreur lors de la conversion en JSON", http.StatusInternalServerError)
					return
				}
	
				w.Write(jsonData)
			} else {
				http.Error(w, "Erreur lors de l'appel de l'API saveIMG", http.StatusInternalServerError)
				return
			}
		}
	})).ServeHTTP(w,r)
}