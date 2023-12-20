package middleware

import (
	"fmt"

	"github.com/adarsh-jaiss/GO-Hotel-reservation/types"
	"github.com/gofiber/fiber/v2"
)

func AdminAuth(c *fiber.Ctx) error  {
	user,ok := c.Context().UserValue("user").(*types.User)
	if !ok{
		return fmt.Errorf("not authorised")
	}

	if !user.IsAdmin{
		return fmt.Errorf("not authorised")
	}

	return c.Next()

}