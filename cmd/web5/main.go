package main

import (
	"net/http"
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
	log.Println("Запуск сервера на :8080")
	err := http.ListenAndServe(":8080", v1)
	if err != nil {
		log.Printf("server close error: %v ", err)
	}

}
