package models

type AuthResponse struct {
    Message string `json:"message"`
    Error   string `json:"error,omitempty"`
}
