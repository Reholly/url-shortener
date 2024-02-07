package storage

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"urlshortener/internal/config"
	"urlshortener/internal/domain"
)

type Storage struct {
	db               *sqlx.DB
	connectionString string
	UrlRepository    UrlRepositoryContract
}

func NewStorage(cfg config.Config) *Storage {
	return &Storage{
		connectionString: cfg.ConnectionString,
	}
}

type UrlRepositoryContract interface {
	GetByAlias(alias string) (domain.Url, error)
	GetByUrl(url string) (domain.Url, error)
	GetAll() ([]domain.Url, error)
	Add(url domain.Url) error
	Delete(url domain.Url) error
	Update(url domain.Url) error
}

func (s *Storage) Open() (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", s.connectionString)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	s.db = db

	return db, nil
}

func (s *Storage) Close() error {
	err := s.db.Close()
	return err
}
