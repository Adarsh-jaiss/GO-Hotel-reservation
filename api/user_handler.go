package api

import (
	"errors"
	

	"github.com/adarsh-jaiss/GO-Hotel-reservation/db"
	"github.com/adarsh-jaiss/GO-Hotel-reservation/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
		if errors.Is(err,mongo.ErrNoDocuments){
			return c.JSON(map[string]string{"msg":"user not found"})
		}
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

	if errors := params.ValidateUsers(); len(errors) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
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

func (h *Userhandler) HandleDeleteUser(c *fiber.Ctx) error  {
	userID := c.Params("id")
	if err := h.userstore.DeleteUsers(c.Context(),userID); err!=nil{
		return err
	}
	return c.JSON(map[string]string{"Deleted": userID})
}

func (h *Userhandler) HandlePutUser(c *fiber.Ctx) error {
    var (
        params types.UpdateUserParams
        userID = c.Params("id")
    )

    oid, err := primitive.ObjectIDFromHex(userID)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(map[string]string{"error": "invalid user ID"})
    }

    if err := c.BodyParser(&params); err != nil {
        return err
    }

    // fmt.Printf("Update Params: %+v\n", params)

    filter := bson.M{"_id": oid}
    if err := h.userstore.UpdateUsers(c.Context(), filter, params); err != nil {
        return err
    }

    return c.JSON(map[string]string{"Updated": userID})
}



