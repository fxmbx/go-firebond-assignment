package db

import (
	"database/sql"
	"go-firebond-assignment/config"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDb *sql.DB

func TestMain(m *testing.M) {
	config, err := config.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config files: ", err)
	}

	testDb, err = sql.Open(config.DBDriver, config.DBSoureTest)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	testQueries = New(testDb)
	os.Exit(m.Run())
}
