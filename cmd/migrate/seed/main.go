package main

import (
	"log"

	"github.com/Verifieddanny/go-social/internal/db"
	"github.com/Verifieddanny/go-social/internal/env"
	"github.com/Verifieddanny/go-social/internal/store"
)

func main() {
	addr := env.GetEnv("DB_ADDR", "postgres://admin:adminpassword@localhost/social?sslmode=disable")
	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	store := store.NewStorage(conn)
	db.Seed(store, conn)
}
