package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/phatwasin01/ticketx-line-oa/api"
	db "github.com/phatwasin01/ticketx-line-oa/db/sqlc"
	"github.com/phatwasin01/ticketx-line-oa/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load configuration file:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		// log.Println(err)
		log.Fatal("cannot connect to db:", err)
	}
	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot setup server", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server:", err)
	}
}
