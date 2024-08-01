package component

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"shellrean.id/golang_auth/internal/config"
)

func GetDatabaseConnection(cnf *config.Config) *sql.DB {
	dns := fmt.Sprintf(
		"host=%s "+
			"port=%s "+
			"dbname=%s "+
			"user=%s "+
			"password=%s "+
			"sslmode=disable",
		cnf.DataBase.Host,
		cnf.DataBase.Port,
		cnf.DataBase.Name,
		cnf.DataBase.User,
		cnf.DataBase.Password,
	)
	connection, err := sql.Open("postgres", dns)
	if err != nil {
		log.Fatal("error when open connection %s", err.Error())
	}

	err = connection.Ping()
	if err != nil {
		log.Fatal("error when open connection %s", err.Error())
	}

	return connection
}
