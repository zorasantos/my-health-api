package utils

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/zorasantos/my-health/config"
)

func GenerateToken(userId uuid.UUID, email string, username string) (string, error) {
	var secretKey = config.GetEnvVars().SecretKey
	claims := jwt.MapClaims{}
	claims["user_id"] = userId
	claims["email"] = email
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

func GenerateTokenEmail() (string, error) {
	var secretKey = config.GetEnvVars().SecretKey

	claims := jwt.MapClaims{}
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	var secretKey = []byte(config.GetEnvVars().SecretKey)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

func GetUserIDFromJWT(r *http.Request) string {
	user_id := r.Header.Get("user_id")
	return user_id
}
