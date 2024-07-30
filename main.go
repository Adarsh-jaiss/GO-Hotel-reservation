package main

import (
	"context"
	
	"os"

	"log"
	
	"github.com/adarsh-jaiss/GO-Hotel-reservation/api"
	"github.com/adarsh-jaiss/GO-Hotel-reservation/db"
	"github.com/adarsh-jaiss/GO-Hotel-reservation/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// dbname = "Hotel-reservation"
)

func Init() {
    if err := godotenv.Load(); err != nil {
        log.Fatal(err)
    }

	// for debugging purposes only
    // for _, envVar := range os.Environ() {
    //     fmt.Println(envVar)
    // }
}


func main() {
	// connecting with Mongo DB
	Init()
	mongoEndPoint:= os.Getenv("MONGO_DB_URL")
	

	// for debugging purposes only
	// fmt.Println("MONGO_DB_URL:", mongoEndPoint)  
	// if mongoEndPoint == "" {
	// 	log.Fatal("MONGO_DB_URL is not set in the environment")
	// }


	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoEndPoint))
	if err != nil {
		log.Fatal(err)
	}


	//  handlers initialization
	var (
	
	hotelStore = db.NewMongoHotelStore(client)
	roomStore = db.NewMongoRoomStore(client,hotelStore)
	userStore = db.NewMongoUserStore(client)
	bookingStore  = db.NewMongoBookingStore(client)
	
	store = &db.Store{
		User: userStore,
		Hotel: hotelStore,
		Room: roomStore,
		Booking : bookingStore,
	}
	hotelHandler = api.NewHotelHandler(store)
	roomHandler = api.NewRoomHandler(store)
	userHandler = api.NewUserhandler(userStore)
	// authHandler = api.NewAuthHandler(userStore)
	BookingHandler = api.NewBookingHandler(store)

	// Routes Grouping
	app	= fiber.New(config)
	// auth =app.Group("api")
	// appV1 = app.Group("api/v1",middleware.JWTAuthentication(userStore))
	appV1 = app.Group("api/v1")
	admin = appV1.Group("/admin",middleware.AdminAuth)

)
	
	
	// Auth Handlers
	// auth.Post("/auth",authHandler.HandleAuthenticate )
	
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

	// rooms handlers
	appV1.Get("/room",roomHandler.HandleGetRooms)
	appV1.Post("/room/:id/book",roomHandler.HandleBookRoom )

	// Todo : cancel a booking
	// TODO : GET ALL BOOKINGS BY A USER

	// booking Handlers
	appV1.Get("/booking/:id",BookingHandler.HandleGetBooking)
	appV1.Get("/booking/:id/cancel",BookingHandler.HandleUpdateBooking)

	// admin handlers
	admin.Get("/booking",BookingHandler.HandleGetBookings)
	

	// listerAddr:= flag.String("listenAddr",":3000","This is the listen Address of the API Server")
	listenAddr := os.Getenv("HTTP_LISTEN_ADDRESS")
	app.Listen(listenAddr)

	
}

// central error handling
// var config = fiber.Config{
//     ErrorHandler: func(c *fiber.Ctx, err error) error {
//         return c.JSON(map[string]string{"error":err.Error()})
//     },
// }

var config = fiber.Config{
    ErrorHandler: api.ErrorHandler,
}








