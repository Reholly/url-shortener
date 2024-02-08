package storage

import (
	"github.com/jmoiron/sqlx"
)

type StorageContract interface {
	Open() (*sqlx.DB, error)
	Close() error
}
