package seeding

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/adarsh-jaiss/GO-Hotel-reservation/db"
	"github.com/adarsh-jaiss/GO-Hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddUser(store *db.Store,fName,lName string, admin bool) *types.User {
	user, err := types.NewUserFromParams(types.CreateUserParams{
		FirstName: fName,
		LastName: lName,
		Email: fmt.Sprintf("%s@%s.com",fName,lName),
		Password: fmt.Sprintf("%s_%s",fName,lName),

	})
	if err!= nil{
		log.Fatal(err)
	}
	user.IsAdmin = admin

	InsertedUser,err := store.User.InsertUsers(context.TODO(),user)
	if err!= nil{
		log.Fatal(err)
	}
	return InsertedUser	
}

func AddHotel(store *db.Store, name,location string, rating int,rooms []primitive.ObjectID) *types.Hotel {
	var roomIDS = rooms
	if rooms == nil{
		roomIDS = []primitive.ObjectID{}
	}
	hotel := types.Hotel{
		Name: name,
		Location: location,
		Rooms: roomIDS,
		Rating: rating,
	}

	InsertedHotel,err := store.Hotel.InsertHotel(context.TODO(),  &hotel)
	if err!= nil{
		log.Fatal(err)
	}
	return InsertedHotel
}

func AddRooms(store *db.Store, size string, ss bool,price float64, hotelID primitive.ObjectID) *types.Room {
	rooms := &types.Room{
		Size: size,
		Price: price,
		HotelID: hotelID,
		Seaside: ss,
	}

	InsertedRoom , err := store.Room.InsertRoom(context.TODO(), rooms)
	if err != nil {
		log.Fatal(err)
	}
	return InsertedRoom
}

func AddBooking(store *db.Store, uid, rid primitive.ObjectID, from, till time.Time) *types.Booking {
	booking := &types.Booking{
		UserID:   uid,
		RoomID:   rid,
		FromDate: from,
		TillDate: till,
	}
	insertedBooking, err := store.Booking.InsertBooking(context.Background(), booking)
	if err != nil {
		log.Fatal(err)
	}
	return insertedBooking
}
