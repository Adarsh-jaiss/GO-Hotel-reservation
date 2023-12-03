package api

import (
	"context"

	"github.com/adarsh-jaiss/GO-Hotel-reservation/db"
	"github.com/adarsh-jaiss/GO-Hotel-reservation/types"
	"github.com/gofiber/fiber/v2"
)

type Userhandler struct{
	userstore db.UserStorer
}

// creating some kind of constructor
func NewUserhandler(userstore db.UserStorer) *Userhandler {
	return &Userhandler{
		userstore: userstore,
		
	}
}

func (h *Userhandler) HandlerUser(c *fiber.Ctx) error {
	id := c.Params("id")
	// we used the User handler which implemented the userstore -> now this user store has a property i.e GetuserbyID -> And this property implements the User struct
	// So we can normally access the ID of a user via this logic and store it in a user variable
	ctx := context.Background()
	userID, err := h.userstore.GetUserByID(ctx,id)
	if err!= nil{
		return err
	}

	return c.JSON(userID)
	
}


func (h *Userhandler) HandlerUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "Rohan",
		LastName: "Nanda",
	}
	return c.JSON(u)
}


