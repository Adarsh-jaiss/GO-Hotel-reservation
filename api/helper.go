package api

import (
	"fmt"

	"github.com/adarsh-jaiss/GO-Hotel-reservation/types"
	"github.com/gofiber/fiber/v2"
)

func GetAuthUser(c fiber.Ctx) (*types.User,error) {
	user,ok := c.Context().UserValue("user").(*types.User)
	if !ok {
		return nil, fmt.Errorf("unauthorized")
	}
	return user, nil
}