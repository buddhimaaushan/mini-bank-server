package main

import (
	"log"

	"github.com/buddhimaaushan/mini_bank/api"
	"github.com/buddhimaaushan/mini_bank/db"
	"github.com/buddhimaaushan/mini_bank/utils"
	_ "github.com/golang/mock/mockgen/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

var conn *pgxpool.Pool
var config utils.Config

func init() {
	var err error

	// Load environment variables
	config, err = utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}
	log.Println("environment variables loaded")

	// Connect to database
	conn, err = db.ConnectToDb(config.DatabaseURL)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	log.Println("database connection established")
}

func main() {
	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
