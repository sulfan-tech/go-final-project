package jwt

import (
	"context"
	"fmt"
	usermodel "go-final-project/internal/domain/user/model"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Key string

const (
	UserKey  Key = "user"
	EmailKey Key = "email"
)

// GenerateJWT generates a JWT for the provided user with a custom expiration time.
func GenerateJWT(user *usermodel.User, secret string, expiration time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user,
		"exp": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWT(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Validate the JWT token
		authToken := c.GetHeader("Authorization")
		splitToken := strings.Split(authToken, "Bearer ")

		if len(splitToken) < 2 {
			c.JSON(http.StatusUnauthorized, "User Unauthorized")
			c.Abort()
			return
		}

		token := splitToken[1]
		t, err := validateToken(token)

		if err != nil {
			c.JSON(http.StatusUnauthorized, "User Unauthorized: "+err.Error())
			c.Abort()
			return
		}

		claims, ok := t.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusInternalServerError, "Failed to get user information from token")
			c.Abort()
			return
		}

		subClaim, ok := claims["sub"].(map[string]interface{})
		if !ok {
			c.JSON(http.StatusInternalServerError, "Failed to get sub claim")
			c.Abort()
			return
		}

		emailFromSubClaim, ok := subClaim["email"].(string)
		if !ok {
			c.JSON(http.StatusInternalServerError, "Failed to get email from sub claim")
			c.Abort()
			return
		}

		ctx := context.WithValue(c.Request.Context(), UserKey, subClaim)
		ctx = context.WithValue(ctx, EmailKey, emailFromSubClaim)

		c.Request = c.Request.WithContext(ctx)
	}
}

// validateToken validates the JWT token.
func validateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		secret := os.Getenv("JWT_SECRET")
		return []byte(secret), nil
	})
}
