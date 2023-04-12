// file: main.go
package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/BurntSushi/toml"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type Config struct {
	Username string
	Password string
	Name     string
}

func main() {
	var config Config
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		log.Fatalf("could not read configuration file: %v", err)
	}

	db, err := sql.Open("pgx", fmt.Sprintf("postgresql://%s:%s@localhost:5432/%s", config.Username, config.Password, config.Name))
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	// To verify the connection to our database instance, we can call the `Ping`
	// method. If no error is returned, we can assume a successful connection
	if err := db.Ping(); err != nil {
		log.Fatalf("unable to reach database: %v", err)
	}
	fmt.Println("database is reachable")
}
