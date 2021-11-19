package handlers

import (
	"context"
	"web5/internal/schedule/doctor"
)

// Schedule представляет интерфейс доступа к расписанию
type Schedule interface {
	FindSchedule(ctx context.Context) ([]*doctor.Schedule, error)
}
