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
	dburi = "mongodb+srv://adarsh_jaiss:baburaokijai@baburao.dg1eflt.mongodb.net/"
	dbname = "Hotel-reservation"

)
func main() {
	// connecting with Mongo DB
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}
	//  handlers initialization
	userHandler := api.NewUserhandler(db.NewMongoUserStore(client,dbname))

	
	// Defining Routes
	app:= fiber.New(config)
	
	//  creating an api group for versioning
	appV1 := app.Group("api/v1")
	appV1.Get("/user",userHandler.HandlerUsers)
	appV1.Get("/user/:id",userHandler.HandlerUser)
	appV1.Post("/user",userHandler.HandlePostUser)
	appV1.Delete("/user/:id",userHandler.HandleDeleteUser)
	appV1.Put("/user/:id",userHandler.HandlePutUser)

	listerAddr:= flag.String("listenAddr",":3000","This is the listen Address of the API Server")
	app.Listen(*listerAddr)

	
}


var config = fiber.Config{
    ErrorHandler: func(c *fiber.Ctx, err error) error {
        return c.JSON(map[string]string{"error":err.Error()})
    },
}






