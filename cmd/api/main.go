package main

import (
	"time"

	_ "github.com/lib/pq"
	"go.uber.org/zap"

	"github.com/Verifieddanny/go-social/internal/db"
	"github.com/Verifieddanny/go-social/internal/env"
	"github.com/Verifieddanny/go-social/internal/mailer"
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
		addr:        env.GetEnv("ADDR", ":8080"),
		env:         env.GetEnv("ENV", "development"),
		apiUrl:      env.GetEnv("EXTERNAL_URL", "localhost:8080"),
		frontendUrl: env.GetEnv("FRONTEND_URL", "http://localhost:3000"),
		db: dbConfig{
			addr:         env.GetEnv("DB_ADDR", ""),
			maxOpenConns: env.GetEnvAsInt("DM_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetEnvAsInt("DM_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetEnv("DM_MAX_IDLE_TIME", "15m"),
		},
		mail: mailConfig{
			exp:       time.Hour * 24 * 3,
			fromEmail: env.GetEnv("FROM_EMAIL", "onboarding@resend.dev"),
			resend: resendConfig{
				apiKey: env.GetEnv("RESEND_API_KEY", ""),
			},
		},
	}

	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	db, err := db.New(cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime)

	if err != nil {
		logger.Fatal(err)
	}

	defer db.Close()
	logger.Info("Database connection pool established")

	store := store.NewStorage(db)
	mailer := mailer.NewResend(cfg.mail.resend.apiKey, cfg.mail.fromEmail)

	app := &application{
		config: cfg,
		store:  store,
		logger: logger,
		mailer: mailer,
	}

	mux := app.mount()

	logger.Fatal(app.run(mux))

}
