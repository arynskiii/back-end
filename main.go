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
		log.Fatal("cannot connect to db:", err)
	}
	store := db.NewStore(conn)

	server := api.NewServer(store)
	if err := server.Start(cfg.ServerAddress); err != nil {
		log.Fatalf("cannot start server: %v", err)
	}
}
