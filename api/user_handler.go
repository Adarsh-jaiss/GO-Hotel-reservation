package api

import (
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
	userID, err := h.userstore.GetUserByID(c.Context(),id)
	if err!= nil{
		return err
	}

	return c.JSON(userID)
	
}


func (h *Userhandler) HandlerUsers(c *fiber.Ctx) error {
	user,err := h.userstore.GetUsers(c.Context())
	if err!= nil{
		return err
	}

	
	return c.JSON(user)
}

func (h *Userhandler) HandlePostUser(c *fiber.Ctx) error {
	var params types.CreateUserParams
	if err:= c.BodyParser(&params); err!=nil{
		return err
	}

	if errors:= params.ValidiateUsers(); len(errors) > 0{
		return c.JSON(errors)
	}

	user,err := types.NewUserFromParams(params)
	if err!= nil{
		return err
	}

	InsertedUser, err:= h.userstore.InsertUsers(c.Context(),user)
	if err!=nil{
		return err

	}

	return c.JSON(InsertedUser)
}

