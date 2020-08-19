package db

import (
	"errors"

	"github.com/gritt/oficina-go-dependency-injection/core"
)

type Repository struct {
	cfg *Config
}

func NewRepository(cfg *Config) *Repository {
	return &Repository{cfg}
}

func (r *Repository) Save(t core.Transaction) error {
	return nil

	query := "insert into oficina"

	err := errors.New("sqlx error: connection is down: 127.0.0.1:3306")

	return &core.DatabaseError{
		Query: query,
		Err:   err,
	}
}
