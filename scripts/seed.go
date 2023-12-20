package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/adarsh-jaiss/GO-Hotel-reservation/api"
	"github.com/adarsh-jaiss/GO-Hotel-reservation/db"
	seeding "github.com/adarsh-jaiss/GO-Hotel-reservation/db/Seeding"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main()  {
	fmt.Println("----------- Seeding the Database ------------")
	var err error
	ctx := context.Background()
    client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
    if err != nil {
        log.Fatal(err)
    }

	if err := client.Database(db.DBNAME).Drop(ctx); err!= nil{
		log.Fatal(err)
	}

	store:= &db.Store{
		User: db.NewMongoUserStore(client),
		Booking: db.NewMongoBookingStore(client),
		Hotel: db.NewMongoHotelStore(client),
		Room: db.NewMongoRoomStore(client,db.DBNAME,db.NewMongoHotelStore(client)),
	}

	user := seeding.AddUser(store,"adarsh","jaiswal",false)
	fmt.Println("Adarsh ->",api.CreateTokenFromUser(user))
	admin := seeding.AddUser(store,"aakriti","awasthi",true)
	fmt.Println("aakrti ->",api.CreateTokenFromUser(admin))
	hotel := seeding.AddHotel(store,"royal palace","bhopal",4,nil)
	fmt.Println(hotel)
	room := seeding.AddRooms(store,"medium",true,88.4,hotel.ID)
	fmt.Println(room)
	booking := seeding.AddBooking(store,user.ID,room.ID,time.Now(),time.Now().AddDate(0,0,2))
	fmt.Println("booking ->", booking.ID)	
	fmt.Println("----------Database seeding completed-------")
}

