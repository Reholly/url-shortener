package storage

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"urlshortener/internal/storage/postgres"
)

type Storage struct {
	db               *sqlx.DB
	connectionString string
	UrlRepository    postgres.UrlRepositoryContract
}

func New(connectionString string) *Storage {
	return &Storage{
		connectionString: connectionString,
	}
}

func (s *Storage) Open() error {
	db, err := sqlx.Open("postgres", s.connectionString)

	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

func (s *Storage) Close() error {
	err := s.db.Close()
	return err
}
