package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/adarsh-jaiss/GO-Hotel-reservation/api"
	"github.com/adarsh-jaiss/GO-Hotel-reservation/db"
	seeding "github.com/adarsh-jaiss/GO-Hotel-reservation/db/Seeding"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main()  {

	if err := godotenv.Load(); err != nil {
        log.Fatal(err)
    }
	var (
		err error
		mongoEndPoint = os.Getenv("MONGO_DB_URL")
		mongoDbName = os.Getenv(db.MongoDbEnvName)
	)

	fmt.Println("----------- Seeding the Database ------------")
	ctx := context.Background()
    client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoEndPoint))
    if err != nil {
        log.Fatal(err)
    }

	fmt.Println(mongoDbName)
	if err := client.Database(mongoDbName).Drop(ctx); err!= nil{
		log.Fatal(err)
	}

	store:= &db.Store{
		User: db.NewMongoUserStore(client),
		Booking: db.NewMongoBookingStore(client),
		Hotel: db.NewMongoHotelStore(client),
		Room: db.NewMongoRoomStore(client,db.NewMongoHotelStore(client)),
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

	for i := 0; i < 100; i++ {
		name := fmt.Sprintf("Hotel %d",i)
		location := fmt.Sprintf("location %d",i)
		seeding.AddHotel(store,name, location, 4 ,nil)
	}
}

