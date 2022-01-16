package utils

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"

	"strconv"
	"time"
)

func GetClaim(c *fiber.Ctx, key string) string {
	auth := c.Locals("auth").(*jwt.Token)
	claims := auth.Claims.(jwt.MapClaims)
	value := claims[key].(string)
	return value
}

func GenerateNewAccessToken(id string, secret string) (string, error) {
	// Set secret key from .env file.
	//os.Getenv("JWT_SECRET_KEY")
	fmt.Println(secret)
	// Set expires minutes count for secret key from .env file.
	minutesCount, _ := strconv.Atoi("60")

	// Create a new claims.
	claims := jwt.MapClaims{}

	// Set public claims:
	claims["id"] = id
	claims["expires"] = time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix()

	// Set private token credentials:
	//for _, credential := range credentials {
	//	claims[credential] = true
	//}

	// Create a new JWT access token with claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate token.
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		// Return error, it JWT token generation failed.
		return "", err
	}

	return t, nil
}
