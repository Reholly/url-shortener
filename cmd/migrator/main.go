package main

import (
	"errors"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"url-shortener/internal/config"
)

const (
	filePath = "file://"
)

func main() {
	cfg := config.LoadConfig()

	m, err := migrate.New(
		filePath+cfg.MigrationsPath,
		cfg.ConnectionString,
	)
	defer func(m *migrate.Migrate) {
		err, _ := m.Close()
		if err != nil {
			panic(err.Error())
		}
	}(m)

	if err != nil {
		panic(err.Error())
	}
	err = m.Up()

	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		panic(err.Error())
	}
}
