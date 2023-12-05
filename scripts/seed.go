package main

import (
	"context"
	"fmt"
	"log"

	"github.com/adarsh-jaiss/GO-Hotel-reservation/db"
	"github.com/adarsh-jaiss/GO-Hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
	roomStore db.RoomStore
	hotelStore *db.MongoHotelStore
	ctx = context.Background()
)

func SeedHotel(name,location string) {
	hotel := types.Hotel{
		Name: name,
		Location: location,
		Rooms: []primitive.ObjectID{},
	}

	rooms := []types.Room{
		// Type: types.DoubleRoomType,
		// BasePrice: 99.9,
		{
			Type: types.DeluxeRoomType,
			BasePrice: 100.9,
		},
		
		{
			Type: types.SeaSideRoomType,
			BasePrice: 150.9,
		},
		{
			Type: types.DeluxeRoomType,
			BasePrice: 250.9,
		},
	}

	InsertedHotel,err := hotelStore.InsertHotel(ctx,  &hotel)
	if err!= nil{
		log.Fatal(err)
	}
	
	fmt.Println("Seeding the Database........")
	// fmt.Println(InsertedHotel)

	for i := range rooms {
		rooms[i].HotelID = InsertedHotel.ID
	}

	
	for _, room := range rooms {
		_ , err := roomStore.InsertRoom(ctx, &room)
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println(InsertedRoom)
	}
	fmt.Println("Database Seeding completed :)")
	
}

func main()  {
	Init()
	SeedHotel("royal palace", "bhopal")
	
}

func Init()  {
	var err error

	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}
	
	hotelStore = db.NewMongoHotelStore(client)
	roomStore = db.NewMongoRoomStore(client, db.DBNAME, hotelStore)
}