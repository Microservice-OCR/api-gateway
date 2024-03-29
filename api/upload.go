package handler

import (
    "api-gateway/img"
    "api-gateway/middleware"
    "net/http"
)

func ImageUploadHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
    
    if r.Method == "OPTIONS" {
        w.WriteHeader(http.StatusOK)
        return
    }
    
    middleware.JwtMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        

        file, header, err := r.FormFile("image")
        if err != nil {
            http.Error(w, "Failed to get image: "+err.Error(), http.StatusBadRequest)
            return
        }
        defer file.Close()

        userId := r.FormValue("userId")

        imageId, err := img.SendImageToAPI(file, userId, header)
        if err != nil {
            http.Error(w, "Failed to send image to API: "+err.Error(), http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusOK)
        w.Write([]byte(imageId))
    })).ServeHTTP(w, r)
}
