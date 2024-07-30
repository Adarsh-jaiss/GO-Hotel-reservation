package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/adarsh-jaiss/GO-Hotel-reservation/api"
	"github.com/adarsh-jaiss/GO-Hotel-reservation/db"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthentication(userstore db.UserStorer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("X-Api-Token")
		if token == "" {
			fmt.Println("Token not present in the header")
			return api.ErrUnAuthorised()
		}

		claims, err := ValidateTokens(token)
		if err != nil {
			return err
		}

		expiresFloat := claims["expires"].(float64)
		expires := int64(expiresFloat)

		// check token expiration
		fmt.Println(expires)
		if time.Now().Unix() > expires {
			return api.NewError(http.StatusUnauthorized, "token expired")
		}

		userID := claims["id"].(string)
		user, err := userstore.GetUserByID(c.Context(), userID)
		if err != nil {
			return api.ErrUnAuthorised()
		}

		// set the current authenticated user to the context
		c.Context().SetUserValue("user", user)

		// Set the token in the header for subsequent requests
		c.Set("X-Api-Token", token)

		return c.Next()
	}
}

// func JWTAuthntication(c *fiber.Ctx) error {
// 	fmt.Println("---- JWT Authenticating-----")

// 	// Get the value of the X-Api-Token header
// 	token := c.Get("X-Api-Token")
// 	if token == "" {
// 		return fmt.Errorf("unauthorized")
// 	}

// 	fmt.Println("token : ",token)

// 	if err := ParseTokens(token); err != nil {
// 		return err
// 	}

// 	return nil
// }

func ValidateTokens(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Printf("Invalid signing method: %v\n", token.Header["alg"])
			return nil, api.ErrUnAuthorised()
		}

		secret := os.Getenv("JWT_SECRET")
		// fmt.Println("NEVER PRINT SECRET",secret) //for debugging purposes only
		return []byte(secret), nil
	})

	if err != nil {
		fmt.Println("Failed to parse JWT token: ", err)
		return nil, api.ErrUnAuthorised()
	}

	if !token.Valid {
		fmt.Println("Invalid token: ", err)
		return nil, api.ErrUnAuthorised()
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, api.ErrUnAuthorised()
	}

	return claims, nil
}

// func ParseTokens(tokenStr string) error {
// 	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			fmt.Printf("Invalid signing method: %v\n", token.Header["alg"])
// 			return nil, fmt.Errorf("unauthorized")
// 		}

// 		secret := os.Getenv("JWT_SECRET")
// 		fmt.Println("NEVER PRINT SECRET",secret)
// 		return []byte(secret), nil
// 	})

// 	if err != nil {
// 		fmt.Println("Failed to parse JWT token: ", err)
// 		return fmt.Errorf("unauthorized")
// 	}

// 	if claims, ok := token.Claims.(jwt.MapClaims); ok {
// 		fmt.Println(claims)
// 		return nil
// 	}

// 	return fmt.Errorf("unauthorized")
// }
