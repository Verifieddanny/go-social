package main

import (
	"log"

	_ "github.com/lib/pq"

	"github.com/Verifieddanny/go-social/internal/db"
	"github.com/Verifieddanny/go-social/internal/env"
	"github.com/Verifieddanny/go-social/internal/store"
)

const currentVersion = "0.0.1"

func main() {

	cfg := config{
		addr: env.GetEnv("ADDR", ":8080"),
		env:  env.GetEnv("ENV", "development"),
		db: dbConfig{
			addr:         env.GetEnv("DB_ADDR", "postgres://admin:adminpassword@localhost/social?sslmode=disable"),
			maxOpenConns: env.GetEnvAsInt("DM_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetEnvAsInt("DM_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetEnv("DM_MAX_IDLE_TIME", "15m"),
		},
	}

	db, err := db.New(cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime)

	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	log.Println("Database connection pool established")

	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))

}

// docker compose up
