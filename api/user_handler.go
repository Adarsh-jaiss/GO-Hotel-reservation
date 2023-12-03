package api

import (
	"github.com/adarsh-jaiss/GO-Hotel-reservation/types"
	"github.com/gofiber/fiber/v2"
)

func HandlerUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "Rohan",
		LastName: "Nanda",
	}
	return c.JSON(u)
}

func HandlerUser(c *fiber.Ctx) error {
	return c.JSON("Adarsh")
	
}

