package main

import (
	"context"

	"flag"

	"log"

	"github.com/adarsh-jaiss/GO-Hotel-reservation/api"
	"github.com/adarsh-jaiss/GO-Hotel-reservation/db"
	"github.com/adarsh-jaiss/GO-Hotel-reservation/middleware"
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
	
	hotelStore = db.NewMongoHotelStore(client)
	roomStore = db.NewMongoRoomStore(client,db.DBNAME,hotelStore)
	userStore = db.NewMongoUserStore(client)
	
	store = &db.Store{
		User: userStore,
		Hotel: hotelStore,
		Room: roomStore,
	}
	hotelHandler = api.NewHotelHandler(store)
	userHandler = api.NewUserhandler(userStore)
	authHandler = api.NewAuthHandler(userStore)

	// Defining Routes
	app	= fiber.New(config)
	auth =app.Group("api")
	appV1 = app.Group("api/v1",middleware.JWTAuthntication)
	)
	
	// Auth Handlers
	auth.Post("/auth",authHandler.HandleAuthenticate )
	
	// Versioned API routes
	// This is user handlers
	appV1.Get("/user",userHandler.HandlerUsers)
	appV1.Get("/user/:id",userHandler.HandlerUser)
	appV1.Post("/user",userHandler.HandlePostUser)
	appV1.Delete("/user/:id",userHandler.HandleDeleteUser)
	appV1.Put("/user/:id",userHandler.HandlePutUser)

	// This is hotel handlers
	appV1.Get("/hotel", hotelHandler.HandleGetHotels )
	appV1.Get("/hotel/:id/rooms", hotelHandler.HandleGetRooms )
	appV1.Get("/hotel/:id", hotelHandler.HandleGetHotel )

	listerAddr:= flag.String("listenAddr",":3000","This is the listen Address of the API Server")
	app.Listen(*listerAddr)

	
}


var config = fiber.Config{
    ErrorHandler: func(c *fiber.Ctx, err error) error {
        return c.JSON(map[string]string{"error":err.Error()})
    },
}






