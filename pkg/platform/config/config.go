package config

type Service struct {
}

func New() Service {
	return Service{}
}

func (s Service) Postgres() DatabaseConfiguration {
	host := "localhost"
	user := "postgres"
	dbName := "booking_db"
	password := "1212"
	return DatabaseConfiguration{
		Name:     dbName,
		Username: user,
		Password: password,
		Host:     host,
		Port:     "5432",
	}
}
