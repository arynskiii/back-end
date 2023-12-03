package main

import (
	"back-end/api"
	db "back-end/db/sqlc"
	"back-end/util"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	cfg, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}

	conn, err := sql.Open(cfg.DBDriver, cfg.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: %v", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(cfg, store)
	if err != nil {
		log.Fatal("cannot create server: %v", err)
	}

	if err := server.Start(cfg.ServerAddress); err != nil {
		log.Fatalf("cannot start server: %v", err)
	}
}
