package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/adarsh-jaiss/GO-Hotel-reservation/db"
	"github.com/adarsh-jaiss/GO-Hotel-reservation/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RoomHandler struct{
	store *db.Store

}

type BookRoomParams struct{
	FromDate time.Time `json:"fromDate"`
	TillDate time.Time `json:"tillDate"`
	NumPersons int 	   `json:"numPersons"`
}

type genericResp struct {
	Type string `json:"type"`
	Msg  string `json:"msg"`
}

func NewRoomHandler(store *db.Store) *RoomHandler {
	return &RoomHandler{
		store: store,
		
	}
}

func (h *RoomHandler) HandleBookRoom(c *fiber.Ctx) error{
	var params BookRoomParams
	if err := c.BodyParser(&params); err!= nil{
		return err
	}
	roomID, err := primitive.ObjectIDFromHex(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(map[string]string{"error": "invalid room ID"})
    }

	user, ok := c.Context().Value("user").(*types.User)
	if !ok{
		return c.Status(http.StatusInternalServerError).JSON(genericResp{
			Type :"error",
			Msg :"Internal server error",
		})
	}

	booking := types.Booking{
		UserID: user.ID,
		RoomID: roomID,
		FromDate: params.FromDate,
		TillDate: params.TillDate,
		NumPersons: params.NumPersons,
	}

	fmt.Printf("%+v\n",booking)

	return nil
}