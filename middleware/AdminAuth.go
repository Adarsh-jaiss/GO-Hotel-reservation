package middleware

import (
	"github.com/adarsh-jaiss/GO-Hotel-reservation/api"
	"github.com/adarsh-jaiss/GO-Hotel-reservation/types"
	"github.com/gofiber/fiber/v2"
)

func AdminAuth(c *fiber.Ctx) error  {
	user,ok := c.Context().UserValue("user").(*types.User)
	if !ok{
		return api.ErrUnAuthorised()
	}

	if !user.IsAdmin{
		return api.ErrUnAuthorised()
	}

	return c.Next()

}