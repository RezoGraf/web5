package main

import (
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"web5/internal/config"

	"github.com/pressly/goose"
)

func main()  {
	cfg := config.MustInitConfig()

	db := sqlx.MustConnect("pgx", cfg.DatabaseDSN)
	defer func() {
		dbErr := db.Close()
		if dbErr != nil {
			log.Errorf("db close error: %v ", dbErr)
		}
	}()

	if cfg.Environment != "production" && cfg.MigrationsDIR != "" {
		err := goose.Up(db.DB, cfg.MigrationsDIR)
		if err != nil {
			log.Error(err)
			return
		}
	}
}
