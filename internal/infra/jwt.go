package infra

import (
	"context"
	"fmt"
	"gocommerce/internal/usecase"
	"gocommerce/internal/utils"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const UserKey contextKey = "user_id"

func WithJWTAuth(handlerFunc http.HandlerFunc, usecase usecase.UserUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := utils.GetTokenFromRequest(r)

		token, err := validateJWT(tokenString)
		if err != nil {
			log.Printf("failed to validate token: %v", err)
			permissionDenied(w)
			return
		}

		if !token.Valid {
			log.Println("invalid token")
			permissionDenied(w)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		str := claims["user_id"].(string)

		userId, err := strconv.Atoi(str)
		if err != nil {
			log.Printf("failed to convert user_id to int: %v", err)
			permissionDenied(w)
			return
		}

		u, err := usecase.GetUserByID(r.Context(), userId)
		if err != nil {
			log.Printf("failed to get user by id: %v", err)
			permissionDenied(w)
			return
		}
		// Add the user to the context
		ctx := r.Context()
		ctx = context.WithValue(ctx, UserKey, u.ID)
		r = r.WithContext(ctx)

		// Call the function if the token is valid
		handlerFunc(w, r)
	}
}
func GenerateJWT(secret []byte, userID int, roles string) (string, error) {
	// Calculate the token expiration time
	expiration := time.Now().Add(time.Duration(getEnvAsInt("JWT_EXPIRATION", 3600*5) * int64(time.Second))).Unix()

	// Create the claims
	claims := jwt.MapClaims{
		"user_id": strconv.Itoa(userID), // Ensure userID is a string
		"roles":   roles,                // Add user roles
		"exp":     expiration,           // Expiration timestamp
	}

	// Create a new token with the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenString, nil
}

func validateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})
}

func permissionDenied(w http.ResponseWriter) {
	utils.JSONResponse(w, http.StatusForbidden, map[string]interface{}{"message": "Permission denied"})
}

func GetUserIdFromContext(ctx context.Context) int {
	userId, ok := ctx.Value(UserKey).(int)
	if !ok {
		return -1
	}

	return userId
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}
