package main

import (
	"net/http"
	"web5/internal/handlers"
	"web5/internal/schedule/doctor"

	"web5/internal/config"

	"github.com/jmoiron/sqlx"

	"web5/internal/schedule"

	"github.com/gorilla/mux"

	"log"
)

func main() {
	cfg := config.MustInitConfig()

	db := sqlx.MustConnect("firebirdsql", cfg.DatabaseDSN)
	defer func() {
		dbErr := db.Close()
		if dbErr != nil {
			log.Printf("db close error: %v ", dbErr)
			// log.Errorf("db close error: %v ", dbErr)
		}
	}()

	router := mux.NewRouter()
	v1 := router.PathPrefix("/v1").Subrouter()

	v1.Handle("/search", handlers.NewIndexHandler(schedule.NewManager(doctor.NewStore(db)), cfg.TemplateINDEX)).Methods("GET")
	log.Println("Запуск сервера на :4000")
	err := http.ListenAndServe(":4000", v1)
	log.Fatal(err)
	// log.Fatal(http.ListenAndServe(":8000", v1))

}
