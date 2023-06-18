package config

// Server holds the
type Server struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type DatabaseConfiguration struct {
	Name     string `json:"dbName"`
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}
