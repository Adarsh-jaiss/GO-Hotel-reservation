package main

import (
	"context"
	
	"flag"


	"log"

	"github.com/adarsh-jaiss/GO-Hotel-reservation/api"
	"github.com/adarsh-jaiss/GO-Hotel-reservation/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	
	// dbname = "Hotel-reservation"

)
func main() {
	// connecting with Mongo DB
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}


	//  handlers initialization
	var (
	userHandler = api.NewUserhandler(db.NewMongoUserStore(client,db.DBNAME))
	hotelStore = db.NewMongoHotelStore(client)
	roomStore = db.NewMongoRoomStore(client,db.DBNAME,hotelStore)
	hotelHandler = api.NewHotelHandler(hotelStore,roomStore)

	// Defining Routes
	app	= fiber.New(config)
	//  creating an api group for versioning
	appV1 = app.Group("api/v1")
	)
	
	
	// This is user handlers
	appV1.Get("/user",userHandler.HandlerUsers)
	appV1.Get("/user/:id",userHandler.HandlerUser)
	appV1.Post("/user",userHandler.HandlePostUser)
	appV1.Delete("/user/:id",userHandler.HandleDeleteUser)
	appV1.Put("/user/:id",userHandler.HandlePutUser)

	// This is hotel handlers
	appV1.Get("/hotel", hotelHandler.HandleGetHotel )

	listerAddr:= flag.String("listenAddr",":3000","This is the listen Address of the API Server")
	app.Listen(*listerAddr)

	
}


var config = fiber.Config{
    ErrorHandler: func(c *fiber.Ctx, err error) error {
        return c.JSON(map[string]string{"error":err.Error()})
    },
}






