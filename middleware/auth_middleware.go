package middleware

import (
	"net/http"
	"strings"

	"github.com/zorasantos/my-health/utils"
)

func AuthenticationMiddleware() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Missing authentication token", http.StatusUnauthorized)
			return
		}

		tokenParts := strings.Split(tokenString, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			http.Error(w, "Invalid authentication token", http.StatusUnauthorized)
			return
		}

		tokenString = tokenParts[1]

		claims, err := utils.VerifyToken(tokenString)
		if err != nil {
			http.Error(w, "Invalid authentication token", http.StatusUnauthorized)
			return
		}

		r.Header.Set("user_id", claims["user_id"].(string))
	})
	// return func(c *gin.Context) {
	// 	tokenString := c.GetHeader("Authorization")
	// 	if tokenString == "" {
	// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authentication token"})
	// 		c.Abort()
	// 		return
	// 	}

	// 	tokenParts := strings.Split(tokenString, " ")
	// 	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
	// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authentication token"})
	// 		c.Abort()
	// 		return
	// 	}

	// 	tokenString = tokenParts[1]

	// 	claims, err := utils.VerifyToken(tokenString)
	// 	if err != nil {
	// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authentication token"})
	// 		c.Abort()
	// 		return
	// 	}

	// 	c.Set("user_id", claims["user_id"])
	// 	c.Next()
	// }
}
