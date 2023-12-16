package api

import (
	"errors"
	"os"
	"time"

	"fmt"

	"github.com/adarsh-jaiss/GO-Hotel-reservation/db"
	"github.com/adarsh-jaiss/GO-Hotel-reservation/types"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthHandler struct {
	userStore db.UserStorer
}

func NewAuthHandler(userstore db.UserStorer) *AuthHandler {
	return &AuthHandler{
		userStore: userstore,
	}
}

type AuthParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct{
	User types.User `json:"user"`
	Token string 	`json:"token"`

}

func (h *AuthHandler) HandleAuthenticate(c *fiber.Ctx) error {
	var Params AuthParams
	if err := c.BodyParser(&Params); err != nil {
		return err
	}

	fmt.Println(Params)

	user, err := h.userStore.GetUserByEmail(c.Context(),Params.Email)
	if err!= nil{
		if errors.Is(err, mongo.ErrNoDocuments){
			return fmt.Errorf("invalid credentials")
		}
	}
	fmt.Println(user) // if email found it will print the user details

	if !types.IsValidPassword(user.EncryptedPassword,Params.Password){
		return fmt.Errorf("invalid credentials")
	}

	// fmt.Println("authenticated ->",user)
	res := AuthResponse{
		User: *user,
		Token: CreateTokenFromUser(user),
	}

	return c.JSON(res)
}

func CreateTokenFromUser(user *types.User) string {
	now := time.Now()
	Expires := now.Add(time.Hour * 4).Unix()
	claims := jwt.MapClaims{
		"id": user.ID,
		"email":user.Email,
		"expires":Expires,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")
	fmt.Println("----",secret)
	tokenStr,err := token.SignedString([]byte(secret))
	if err !=nil{
		fmt.Println("failed to signed token with secret :",err)
	}
	return tokenStr 
}
