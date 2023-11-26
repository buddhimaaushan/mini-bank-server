package db_test

import (
	"log"
	"os"
	"testing"

	"github.com/buddhimaaushan/mini_bank/db"
	"github.com/buddhimaaushan/mini_bank/utils"
	"github.com/jackc/pgx/v5/pgxpool"
)

var store db.Store
var testDB *pgxpool.Pool
var config utils.Config

func TestMain(m *testing.M) {

	var err error

	// Load environment variables
	config, err = utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}
	log.Println("environment variables loaded")
	log.Println(config)

	// Connect to database
	testDB, err = db.ConnectToDb(config.DatabaseURL)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	log.Println("database connection established")

	store = db.NewStore(testDB)
	os.Exit(m.Run())
}
