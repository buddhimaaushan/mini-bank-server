package db_test

import (
	"log"
	"os"
	"testing"

	"github.com/buddhimaaushan/mini_bank/db"
	sqlc "github.com/buddhimaaushan/mini_bank/db/sqlc"
	"github.com/buddhimaaushan/mini_bank/utils"
	"github.com/jackc/pgx/v5/pgxpool"
)

var testQueries *sqlc.Queries
var testDB *pgxpool.Pool
var config utils.Config

func TestMain(m *testing.M) {
	var err error
	config, err = utils.LoadConfig("../../")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}
	testDB, err = db.ConnectToDb(config.DatabaseURL)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	testQueries = sqlc.New(testDB)
	os.Exit(m.Run())
}
