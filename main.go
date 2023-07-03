package main

import (
	"database/sql"
	"go-firebond-assignment/api"
	"go-firebond-assignment/config"
	db "go-firebond-assignment/db/sqlc"
	"go-firebond-assignment/worker"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("error loading configurations %v ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSoure)
	if err != nil {
		log.Fatalf("error connecting to db %v", err)
	}

	defer conn.Close()
	store := db.NewStore(conn)

	ticker := time.NewTicker(5 * time.Minute)
	ws := worker.NewWorkerService(config, store)
	ws.Worker()

	go func() {
		for range ticker.C {
			ws.Worker()
		}
	}()

	runGinServer(config, store)
}

func runGinServer(config config.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal(err)
	}
}
