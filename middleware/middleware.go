package middleware

import (
	"net/http"
	"strings"
	"time"

	"go-crud/config"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	// Define public routes that allow access only for POST method
	publicRoutes := map[string]bool{
		"/login": true,
		"/users": true,
		// Add more public routes as needed
	}

	return func(c *gin.Context) {
		// Get the request path
		path := c.Request.URL.Path

		// Check if the request method is POST and if it's a public route
		if c.Request.Method == http.MethodPost && publicRoutes[path] {
			// If it's a public route and method is POST, proceed to the next middleware or handler
			c.Next()
			return
		}

		// For other routes or methods, perform JWT authentication
		const BearerSchema = "Bearer "
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, BearerSchema)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.GetEnv("JWT_SECRET")), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			exp := int64(claims["exp"].(float64))
			if time.Unix(exp, 0).Before(time.Now()) {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
				c.Abort()
				return
			}
			c.Set("user_id", claims["user_id"])
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
		}
	}
}
