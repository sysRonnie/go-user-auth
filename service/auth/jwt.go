package auth

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
)

// CreateJWT generates a JWT for a given user ID with a secret key
func CreateJWT(secret []byte, userID int) (string, error) {
	expiration := time.Second * 60
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    strconv.Itoa(userID),
		"expiredAt": time.Now().Add(expiration).Unix(),
	})
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// SetJWTTokenCookie sets the JWT token in an HttpOnly secure cookie
func SetJWTTokenCookie(c echo.Context, token string, expiration time.Duration) {
	cookie := &http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,  // JavaScript cannot access the cookie
		Secure:   true,  // Ensure this is sent over HTTPS only
		Expires:  time.Now().Add(expiration),
	}
	c.SetCookie(cookie)
}


// JWTMiddleware provides the JWT middleware using echo-jwt
func JWTMiddleware(secret string) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(secret),
		TokenLookup: "cookie:token",  // Explicitly look for the token in the cookie named 'token'
	})
}

// GetUserFromToken extracts user information from JWT token
func GetUserFromToken(c echo.Context) (int, error) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userIDStr, ok := claims["userID"].(string)
	if !ok {
		return 0, fmt.Errorf("userID not found in token claims")
	}
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return 0, fmt.Errorf("failed to convert userID to int: %v", err)
	}
	return userID, nil
}

