package schedule

import (
	"context"
	"web5/internal/schedule/doctor"
)

// Store описывает интерфейс доступа к хранилищц
type Store interface {
	FindSchedule(ctx context.Context) ([]*doctor.Schedule, error)
}

// Manager ...
type Manager struct {
	store Store
}

// NewManager ...
func NewManager(store Store) *Manager {
	return &Manager{store: store}
}

func (m *Manager) FindSchedule(ctx context.Context) ([]*doctor.Schedule, error) {
	return m.store.FindSchedule(ctx)
}
