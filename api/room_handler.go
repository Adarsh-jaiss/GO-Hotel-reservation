package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/adarsh-jaiss/GO-Hotel-reservation/db"
	"github.com/adarsh-jaiss/GO-Hotel-reservation/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
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

func (p BookRoomParams) validate() error{
	now := time.Now()
	if now.After(p.FromDate) || now.After(p.TillDate){
		return fmt.Errorf("cannot book a room in the past")
	}
	return nil
}

func(h *RoomHandler) HandleGetRooms(c *fiber.Ctx) error {
	rooms ,err := h.store.Room.GetRooms(c.Context(),bson.M{})
	if err!= nil{
		return err
	}
	return c.JSON(rooms)
	
}

func (h *RoomHandler) HandleBookRoom(c *fiber.Ctx) error{
	var params BookRoomParams
	if err := c.BodyParser(&params); err!= nil{
		return err
	}

	if err := params.validate(); err!= nil{
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



	ok,err = h.isRoomAvailableforBooking(c.Context(),roomID,params)
	if err!= nil{
		return err
	}
	if !ok{
		return c.Status(http.StatusBadRequest).JSON(genericResp{
			Type: "error",
			Msg: fmt.Sprintf("room %s already booked ", c.Params("id")),
		})
	}


	booking := types.Booking{
		UserID: user.ID,
		RoomID: roomID,
		FromDate: params.FromDate,
		TillDate: params.TillDate,
		NumPersons: params.NumPersons,
	}

	inserted,err := h.store.Booking.InsertBooking(c.Context(),&booking)
	if err!= nil{
		return err
	}

	fmt.Printf("%+v\n",booking)

	return c.JSON(inserted)
}

func(h *RoomHandler) isRoomAvailableforBooking(ctx context.Context, roomID primitive.ObjectID, params BookRoomParams) (bool,error) {
	where := bson.M{
		"roomID": roomID,
		"fromDate":bson.M{
			"$gte":params.FromDate,
		},
		"tillDate":bson.M{
			"$gte":params.TillDate,
		},

	}
	bookings,err := h.store.Booking.GetBookings(ctx,where)
	if err!= nil{
		return false,err
	}
	// fmt.Println(bookings)  // for debugging purpose only
	ok := len(bookings) == 0
	return ok,nil

}
