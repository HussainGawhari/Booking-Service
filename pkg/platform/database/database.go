package database

import "database/sql"

func Connection() (*sql.DB, error) {

	// host := "localhost"
	// user := "postgres"
	// dbName := "booking_db"
	// password := "1212"
	// Port:     "5432",
	connectionString := "host=localhost port=5432 user=postgres password=1212 dbname=booking_db sslmode=disable"

	// Open the database connection
	db, err := sql.Open("postgres", connectionString)

	return db, err

}
