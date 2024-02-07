package app

import (
	"urlshortener/internal/config"
	"urlshortener/internal/server"
	"urlshortener/internal/storage"
	"urlshortener/internal/storage/postgres"
)

type App struct {
	Config config.Config
}

func New() *App {
	cfg := config.LoadConfig()

	return &App{Config: *cfg}
}

func (a *App) Run() {
	database := storage.NewStorage(a.Config)
	db, err := database.Open()
	defer db.Close()

	//di
	urlRepo := bindUrlRepository(*postgres.NewUrlRepository(db))
	ser := server.New(urlRepo)

	if err != nil {
		panic(err)
	}

	ser.RunServer(a.Config.Address)
}

func bindUrlRepository(repository postgres.UrlRepository) storage.UrlRepositoryContract {
	return &repository
}
