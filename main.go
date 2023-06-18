package main

import (
	"Booking_Service/api"
	"Booking_Service/pkg/platform/database"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// configs := config.New()
	db, err := database.Connection()

	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(" connection established")

	api := api.RegisterRoutes(db)
	//start server
	api.Run(":8000")

}
