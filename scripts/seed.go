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
	userStore db.UserStorer
)

func Seeduser(fName, lName,email string)  {
	user, err := types.NewUserFromParams(types.CreateUserParams{
		FirstName: fName,
		LastName: lName,
		Email: email,
		Password: "supersecurepassword",

	})

	if err!= nil{
		log.Fatal(err)
	}

	_,err = userStore.InsertUsers(context.TODO(),user)
	if err!= nil{
		log.Fatal(err)
	}
}

func SeedHotel(name,location string, rating int) {
	hotel := types.Hotel{
		Name: name,
		Location: location,
		Rooms: []primitive.ObjectID{},
		Rating: rating,
	}

	rooms := []types.Room{
		// Type: types.DoubleRoomType,
		// BasePrice: 99.9,
		{
			Size: "small",
			Price: 99.9,
		},
		
		{
			Size: "normal",
			Price: 150.9,
		},
		{
			Size: "kingsize",
			Price: 200.9,
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
	// SeedHotel("royal palace", "bhopal", 4)
	// SeedHotel("neena palace", "bhopal", 3)
	// SeedHotel("radision bhopal", "bhopal", 5)
	Seeduser("Adarsh","jaiswal","adarsh@gmail.com")
	
}

func Init() {
    var err error

    client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
    if err != nil {
        log.Fatal(err)
    }

    hotelStore = db.NewMongoHotelStore(client)
    roomStore = db.NewMongoRoomStore(client, db.DBNAME, hotelStore)  // Add the missing argument (db.DBNAME)
    userStore = db.NewMongoUserStore(client)  // Add the missing argument (db.DBNAME)
}
