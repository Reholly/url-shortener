package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"url-shortener/internal/storage"
)

type PostgresStorage struct {
	driverName       string
	connectionString string
	db               *sqlx.DB
}

func NewStorage(driverName, connectionString string) storage.StorageContract {
	return &PostgresStorage{driverName: driverName, connectionString: connectionString}
}

func (s *PostgresStorage) Open() (*sqlx.DB, error) {
	db, err := sqlx.Open(s.driverName, s.connectionString)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func (s *PostgresStorage) Close() error {
	err := s.db.Close()
	return err
}
