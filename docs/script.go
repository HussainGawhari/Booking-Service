package docs

const (
	dbHost     = "localhost"
	dbPort     = 5432
	dbUser     = "postgres"
	dbPassword = "1212"
	dbName     = "booking_db"
)

dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
	db, err := sql.Open("postgres", dbInfo)
	defer db.Close()
	// fmt.Printf(dbInfo)
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(" \n Database connection established")

	filePath := "SeatPricing.csv" // Replace with the path to your CSV file
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")

		id := line[0]
		searClass := line[1]
		minPrice := line[2]
		normalPrice := line[3]
		maxPrice := line[4]

		stmt, err := db.Prepare("INSERT INTO seat_class VALUES ($1, $2, $3, $4, $5)")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()

		_, err = stmt.Exec(id, searClass, minPrice, normalPrice, maxPrice)
		// fmt.Println(id, searClass+" "+minPrice+" "+normalPrice+" "+maxPrice)
	}