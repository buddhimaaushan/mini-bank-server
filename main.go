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
	config, err = utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}
	conn, err = db.ConnectToDb(config.DatabaseURL)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
}

func main() {
	store := db.NewStore(conn)
	server := api.NewServer(store)
	err := server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
