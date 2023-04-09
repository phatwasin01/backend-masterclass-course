package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/ticketx?sslmode=disable"
)

var testQueries *Queries

var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		// log.Println(err)
		log.Fatal("cannot connect to db:", err)
	}
	// require.NoError(m, err)
	testQueries = New(testDB)
	os.Exit(m.Run())
}
