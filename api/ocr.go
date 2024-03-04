package handler

import (
	"api-gateway/middleware"
	"api-gateway/ocr"
	"api-gateway/ocr/models"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func OcrHandler(w http.ResponseWriter, r *http.Request) {
	// TODO : SUPPRIMER POUR PROD
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*") // or specify your domain
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
    if r.Method == "OPTIONS" {
        w.WriteHeader(http.StatusOK)
        return
    }
	middleware.JwtMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Logique pour OCR
		if r.Method == "GET" {
			// Récupération de l'id de l'image
			imageId := r.URL.Query().Get("id")
			// Récupération du nom de l'image
			imageName := r.URL.Query().Get("image")

			if imageId != "" {
				// Appel de l'API OCR
				ocrData, err := ocr.GetOCRFromId(imageId)
				if err != nil {
					http.Error(w, "Erreur lors de l'appel de l'API OCR", http.StatusInternalServerError)
					return
				}

				jsonData, err := json.Marshal(ocrData)
				if err != nil {
					http.Error(w, "Erreur lors de la conversion en JSON", http.StatusInternalServerError)
					return
				}

				w.Write(jsonData)
			} else if imageName != "" {
				// Appel de l'API OCR
				ocrData, err := ocr.GetOCR(imageName)
				if err != nil {
					http.Error(w, "Erreur lors de l'appel de l'API OCR", http.StatusInternalServerError)
					return
				}

				jsonData, err := json.Marshal(ocrData)
				if err != nil {
					http.Error(w, "Erreur lors de la conversion en JSON", http.StatusInternalServerError)
					return
				}

				w.Write(jsonData)
			} else {
				http.Error(w, "Erreur lors de l'appel de l'API OCR", http.StatusInternalServerError)
				return
			}
		} else if r.Method == "POST" {
			// Récupération de l'id de l'image
			imageId := r.URL.Query().Get("id")
			// Récupération du nom de l'image
			input := models.IInput{}

			// Lecture du corps de la requête
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Erreur lors de la lecture du corps de la requête", http.StatusInternalServerError)
				return
			}

			// Décodage du JSON dans la structure IInput
			err = json.Unmarshal(body, &input)
			if err != nil {
				http.Error(w, "Erreur lors de la conversion du JSON", http.StatusInternalServerError)
				return
			}

			if imageId != "" {
				// Appel de l'API OCR
				ocrData, err := ocr.PostOCRFromId(imageId, input)
				if err != nil {
					http.Error(w, "Erreur lors de l'appel de l'API OCR", http.StatusInternalServerError)
					return
				}

				jsonData, err := json.Marshal(ocrData)
				if err != nil {
					http.Error(w, "Erreur lors de la conversion en JSON", http.StatusInternalServerError)
					return
				}

				w.Write(jsonData)
			} else {
				http.Error(w, "Erreur lors de l'appel de l'API OCR", http.StatusInternalServerError)
				return
			}
		} else {
			http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
			return
		}
	})).ServeHTTP(w, r)
}
