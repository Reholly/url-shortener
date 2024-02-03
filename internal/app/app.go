package app

import (
	"urlshortener/internal/config"
	"urlshortener/internal/server"
	"urlshortener/internal/storage"
	"urlshortener/internal/storage/postgres"
)

type App struct {
	UrlRepo storage.UrlRepositoryContract
	Config  config.Config
}

func New() *App {
	cfg := config.LoadConfig()

	database := storage.New(*cfg)

	db, err := database.Open()

	if err != nil {
		panic(err)
	}

	urlRepo := bindUrlRepository(*postgres.NewUrlRepository(db))

	return &App{UrlRepo: urlRepo, Config: *cfg}
}

func (a *App) Run() {
	ser := server.New(a.UrlRepo)

	ser.RunServer(a.Config.Address)
}

func bindUrlRepository(repository postgres.UrlRepository) storage.UrlRepositoryContract {
	return &repository
}
