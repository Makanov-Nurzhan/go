package handlers

import (
	"crud-project/auth"
	"crud-project/models"
	"encoding/json"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type RefreshResponse struct {
	Token string `json:"access_token"`
}

func RefreshHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RefreshRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.RefreshToken == "" {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		var token models.RefreshToken
		if err := db.Where("token = ?", req.RefreshToken).First(&token).Error; err != nil {
			http.Error(w, "Invalid refresh token", http.StatusBadRequest)
			return
		}
		claims := auth.Claims{
			UserID: token.UserID,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			},
		}
		accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		accessTokenStr, err := accessToken.SignedString(auth.JwtSecret)
		if err != nil {
			http.Error(w, "Failed to generate access token", http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(RefreshResponse{Token: accessTokenStr})
	}
}
