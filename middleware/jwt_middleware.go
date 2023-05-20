package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"time"
)

func CreateToken(userId int, name string) (string, error) {
	claims := jwt.MapClaims{}
	claims["userId"] = userId
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}

func ValidateToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get the Authorization header value
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing Authorization header")
		}

		// Extract the token from the Authorization header
		tokenString := authHeader[len("Bearer "):]
		if tokenString == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing token")
		}

		// Parse and validate the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Provide the secret key or the public key for token validation
			// You may change this based on your token signing method
			return []byte(os.Getenv("SECRET_KEY")), nil
		})
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		// Check if the token is valid
		if !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		// Extract the claims from the token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token claims")
		}

		// Set the claims in the context for further request handling
		c.Set("claims", claims)

		return next(c)
	}
}
