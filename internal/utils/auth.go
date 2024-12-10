package utils

import (
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func GetTokenFromRequest(r *http.Request) string {
	// Get token from Authorization header
	tokenAuth := r.Header.Get("Authorization")
	if strings.HasPrefix(tokenAuth, "Bearer ") {
		return strings.TrimPrefix(tokenAuth, "Bearer ")
	}

	// Get token from URL query parameter
	tokenQuery := r.URL.Query().Get("Token")
	if tokenQuery != "" {
		return tokenQuery
	}

	// Return empty string if no token is found
	return ""
}

func ComparePasswords(hashed string, plain []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), plain)
	return err == nil
}
