package handlers

import (
	"html/template"
	"net/http"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// IndexHandler представляет обработчик получения расписания
type IndexHandler struct {
	schedule     Schedule
	templatePATH string
}

// NewIndexHandler ...
func NewIndexHandler(schedule Schedule, templatePATHC string) *IndexHandler {
	return &IndexHandler{schedule: schedule, templatePATH: templatePATHC}
}

// ServeHTTP обработчик
func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	schedule, err := h.schedule.FindSchedule(r.Context())
	if err != nil {
		log.Error(errors.Wrap(err, "error find schedule"))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(""))
		return
	}

	tmpl, err := template.ParseFiles(h.templatePATH)
	if err != nil {
		log.Error(errors.Wrap(err, "error find schedule"))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(""))
		return
	}

	if err := tmpl.Execute(w, schedule); err != nil {
		log.Error(errors.Wrap(err, "error find schedule"))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(""))
		return
	}
}
