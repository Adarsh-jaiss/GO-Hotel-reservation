package middleware

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthntication(c *fiber.Ctx) error {
	fmt.Println("---- JWT Authenticating-----")

	// Get the value of the X-Api-Token header
	token := c.Get("X-Api-Token")
	if token == "" {
		return fmt.Errorf("unauthorized")
	}

	if err := ParseTokens(token); err != nil {
		return err
	}

	return nil
}

func ParseTokens(tokenStr string) error {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Printf("Invalid signing method: %v\n", token.Header["alg"])
			return nil, fmt.Errorf("unauthorized")
		}

		secret := os.Getenv("JWT_SECRET")
		fmt.Println("NEVER PRINT SECRET",secret)
		return []byte(secret), nil
	})

	if err != nil {
		fmt.Println("Failed to parse JWT token", err)
		return fmt.Errorf("unauthorized")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		fmt.Println(claims)
		return nil
	}

	return fmt.Errorf("unauthorized")
}
