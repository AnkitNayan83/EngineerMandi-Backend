package middlewares

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No Authorization header provided"})
			c.Abort()
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header"})
			c.Abort()
			return
		}

		jwtKeyString := os.Getenv("JWT_KEY")
		if jwtKeyString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "JWT_KEY is not set"})
			c.Abort()
			return
		}

		jwtKey := []byte(jwtKeyString)

		token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			var ve *jwt.ValidationError
			log.Print(err)
			if errors.As(err, &ve) {
				switch {
				case ve.Errors&jwt.ValidationErrorMalformed != 0:
					c.JSON(http.StatusBadRequest, gin.H{"error": "That's not even a token"})
				case ve.Errors&jwt.ValidationErrorExpired != 0:
					c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is expired"})
				case ve.Errors&jwt.ValidationErrorNotValidYet != 0:
					c.JSON(http.StatusUnauthorized, gin.H{"error": "Token not valid yet"})
				default:
					c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is invalid 1"})
				}
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is invalid 2"})
			}
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
			c.Set("userID", claims.Subject)
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is invalid 3"})
			c.Abort()
			return
		}

		c.Next()
	}
}
