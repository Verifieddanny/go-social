package main

import (
	"log"

	_ "github.com/lib/pq"

	"github.com/Verifieddanny/go-social/internal/db"
	"github.com/Verifieddanny/go-social/internal/env"
	"github.com/Verifieddanny/go-social/internal/store"
)

const currentVersion = "0.0.1"

//	@title			Gohpher Social
//	@description	API for GohpherSocial, soial network for gophers
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@BasePath	/v1

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@description

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {

	cfg := config{
		addr:   env.GetEnv("ADDR", ":8080"),
		env:    env.GetEnv("ENV", "development"),
		apiUrl: env.GetEnv("EXTERNAL_URL", "localhost:8080"),
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
