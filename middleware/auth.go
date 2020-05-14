package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const secret = "SECRET"

//Auth is function for Authorization
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		const bearerSchema = "Bearer "
		authHeader := c.GetHeader("Authorization")

		if len(authHeader) < 1 {
			log.Println("No Authorization header")
			c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("You are not authorized for this action"))
			return
		}

		tokenString := authHeader[len(bearerSchema):]

		token, err := validateToken(tokenString)

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claims[authorized]: ", claims["authorized"])
			log.Println("Claims[userId]: ", claims["user_id"])
			log.Println("Claims[expiry]: ", claims["expiry"])
		} else {
			log.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

func validateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Signing method validation
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// Return the secret signing key
		return []byte(os.Getenv(secret)), nil
	})
}
