package api

import (
	"errors"
	"fmt"

	"github.com/adarsh-jaiss/GO-Hotel-reservation/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
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

func (h *AuthHandler) HandleAuthenticate(c *fiber.Ctx) error {
	var Params AuthParams
	if err := c.BodyParser(&Params); err != nil {
		return err
	}

	user, err := h.userStore.GetUserByEmail(c.Context(),Params.Email)
	if err!= nil{
		if errors.Is(err, mongo.ErrNoDocuments){
			return fmt.Errorf("Invalid credentials")
		}
	}
	fmt.Println(user) // if email found it will print the user details

	if err = bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword),[]byte(Params.Password)); err!= nil{
		return fmt.Errorf("Invalid credentials")
	}

	fmt.Println("authenticated ->",user)



	return nil
}
