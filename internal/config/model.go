package config

type Config struct {
	Server   Server
	DataBase DataBase
}

type Server struct {
	Host string
	Port string
}

type DataBase struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}
