package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/adarsh-jaiss/GO-Hotel-reservation/db"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	
)

func JWTAuthentication(userstore db.UserStorer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("X-Api-Token")
		if token == "" {
			fmt.Println("Token not present in the header")
			return fiber.ErrUnauthorized
		}

		claims,err := ValidateTokens(token)
		if err != nil {
			return err
		}

		expiresFloat := claims["expires"].(float64)
		expires := int64(expiresFloat)

		// check token expiration
		fmt.Println(expires)

		if time.Now().Unix() > expires {
			return fmt.Errorf("token expired")
		}

		userID := claims["id"].(string)
		user,err := userstore.GetUserByID(c.Context(), userID)
		if err!= nil{
			return fmt.Errorf("unauthorized")
		}

		// set the current authenticated user to the context
		c.Context().SetUserValue("user",user)
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

func ValidateTokens(tokenStr string) (jwt.MapClaims,error) {
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
		fmt.Println("Failed to parse JWT token: ", err)
		return nil, fmt.Errorf("unauthorized")
	}

	if !token.Valid{
		fmt.Println("Invalid token: ", err)
		return nil,fmt.Errorf("unauthorized")
	}


	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok{
		return nil, fmt.Errorf("unauthorized")
	}

	return claims,nil
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