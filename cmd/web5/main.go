package main

import (
	"net/http"
	"time"
	"web5/internal/handlers"
	"web5/internal/schedule/doctor"

	"web5/internal/config"

	"github.com/jmoiron/sqlx"

	"web5/internal/schedule"

	"github.com/gorilla/mux"

	_ "github.com/nakagami/firebirdsql"

	"log"
)

func main() {
	cfg := config.MustInitConfig()

	db := sqlx.MustConnect("firebirdsql", cfg.DatabaseDSN)
	defer func() {
		dbErr := db.Close()
		if dbErr != nil {
			log.Printf("db close error: %v ", dbErr)
		}
	}()

	router := mux.NewRouter()
	v1 := router.PathPrefix("/v1").Subrouter()
	v1.Handle("/search", handlers.NewIndexHandler(schedule.NewManager(doctor.NewStore(db)), cfg.TemplateINDEX)).Methods("GET")
	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
		// Таймауты сервера! (рекомандация задавать из мануалов gorilla/mux
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())

}
