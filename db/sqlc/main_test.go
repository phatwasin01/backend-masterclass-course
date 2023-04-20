package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/phatwasin01/ticketx-line-oa/util"
)

var testQueries *Queries

var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load configuration file:", err)
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		// log.Println(err)
		log.Fatal("cannot connect to db:", err)
	}
	//
	// require.NoError(m, err)
	testQueries = New(testDB)
	os.Exit(m.Run())
}
