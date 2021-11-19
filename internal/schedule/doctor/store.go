package doctor

import "github.com/jmoiron/sqlx"

// Store представляет хранилище ...
type Store struct {
	db sqlx.DB
}

// NewStore конструктор хранилища
func NewStore(db sqlx.DB) *Store {
	return &Store{db: db}
}
