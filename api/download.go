package handler

import (
	"api-gateway/img"
	"encoding/json"
	"log"
	"net/http"
)

func ImageDownloadHandler(w http.ResponseWriter, r *http.Request) {
	// TODO : SUPPRIMER POUR PROD
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*") // or specify your domain
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == "GET" {
		// Récupération de l'id de l'image
		imageId := r.URL.Query().Get("id")

		if imageId != "" {
			// Appel de l'API saveIMG
			imgData, err := img.GetImageFromId(imageId)
			if err != nil {
				log.Panic("Méthode non autorisée")
				http.Error(w, "Erreur lors de l'appel de l'API saveIMG", http.StatusInternalServerError)
				return
			}

			jsonData, err := json.Marshal(imgData)
			if err != nil {
				log.Panic("Méthode non autorisée")
				http.Error(w, "Erreur lors de la conversion en JSON", http.StatusInternalServerError)
				return
			}
			log.Print("Extraction complétée")

			w.Write(jsonData)
		} else {
			log.Panic("Méthode non autorisée")
			http.Error(w, "Erreur lors de l'appel de l'API saveIMG", http.StatusInternalServerError)
			return
		}
	}
}