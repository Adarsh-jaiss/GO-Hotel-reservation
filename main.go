package main

import (
	"flag"
	"fmt"

	"github.com/adarsh-jaiss/GO-Hotel-reservation/api"
	"github.com/gofiber/fiber/v2"
)

func main()  {
	fmt.Println("Welcome to Hotel reservation system")
	
	app:= fiber.New()
	
	//  creating an api group for versioning
	appV1 := app.Group("api/v1")
	appV1.Get("user",api.HandlerUsers)
	appV1.Get("user/:id",api.HandlerUser)

	listerAddr:= flag.String("listenAddr",":3000","This is the listen Address of the API Server")
	app.Listen(*listerAddr)

}


