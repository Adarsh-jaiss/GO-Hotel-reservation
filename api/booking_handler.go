package api

import (
	
	"net/http"
	"github.com/adarsh-jaiss/GO-Hotel-reservation/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookingHandler struct{
	store *db.Store

}

func NewBookingHandler(store *db.Store) *BookingHandler {
	return &BookingHandler{
		store:store,
	}
}

// Todo :This needs to be admin authorised
func (h *BookingHandler) HandleGetBookings(c *fiber.Ctx) error {
	booking,err := h.store.Booking.GetBookings(c.Context(),bson.M{})
	if err!=nil{
		return err
	}
	return c.JSON(booking)

}

// Todo : This needs to be user authorised
func (h *BookingHandler) HandleGetBooking(c *fiber.Ctx) error {
	oid,err := primitive.ObjectIDFromHex(c.Params("id"))
	if err!= nil{
		return err
	}

	booking,err := h.store.Booking.GetBookingsByID(c.Context(),oid)
	if err!= nil{
		return err
	}

	// restricting users to see bookings of another user ID
	user, err := GetAuthUser(*c)
	if err!=nil{
		return err
	}
	if booking.UserID != user.ID{
		return c.Status(http.StatusUnauthorized).JSON(genericResp{
			Type: "error",
			Msg: "not-authorised",
		})
	}


	return c.JSON(booking)
}

func(h *BookingHandler) HandleUpdateBooking(c *fiber.Ctx) (error) {
	oid, _ := primitive.ObjectIDFromHex(c.Params("id")) 
	booking,err := h.store.Booking.GetBookingsByID(c.Context(),oid)
	if err!= nil{
		return err
	}

	
	user,err := GetAuthUser(*c)
	if err!= nil{
		return err
	}

	if booking.UserID != user.ID{
		return c.Status(http.StatusUnauthorized).JSON(genericResp{
			Type: "error",
			Msg: "not-authorised",
		})
	}

	if err := h.store.Booking.UpdateBooking(c.Context(),c.Params("id"), bson.M{"cancelled":true}); err!= nil{
		return err
	}

	return c.JSON(map[string]string{"msg":"booking cancelled"})
}

