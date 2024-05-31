package middleware

import (
	"net/http"

	"github.com/zorasantos/my-health/utils"
)

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Missing authentication token", http.StatusUnauthorized)
			return
		}

		claims, err := utils.VerifyToken(tokenString)
		if err != nil {
			http.Error(w, "Invalid authentication token", http.StatusUnauthorized)
			return
		}

		r.Header.Set("user_id", claims["user_id"].(string))
		next.ServeHTTP(w, r)
	})
}
