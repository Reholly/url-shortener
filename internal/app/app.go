package app

import (
	"github.com/labstack/echo/v4"
	"url-shortener/internal/config"
	"url-shortener/internal/domain/repositories"
	"url-shortener/internal/server/handlers"
	"url-shortener/internal/storage"
	"url-shortener/internal/storage/postgres"
)

type App struct {
	Config config.Config
}

func New() *App {
	cfg := config.LoadConfig()

	return &App{Config: *cfg}
}

func (a *App) Run() {
	s := postgres.NewStorage(a.Config.DriverName, a.Config.ConnectionString)

	db, err := s.Open()
	defer func(s storage.StorageContract) {
		err := s.Close()
		if err != nil {
			panic(err.Error())
		}
	}(s)

	if err != nil {
		panic(err.Error())
	}

	urlRepo := postgres.NewUrlRepository(db)

	a.runServer(urlRepo)
}

func (a *App) runServer(urlRepo repositories.UrlRepositoryContract) {
	e := echo.New()

	e.POST("/add", handlers.AddUrl(urlRepo))
	e.GET("/get", handlers.GetUrlByAlias(urlRepo))
	e.GET("/:alias", handlers.RedirectOnUrl(urlRepo))
	e.POST("/remove", handlers.RemoveUrl(urlRepo))

	err := e.Start(a.Config.Address)

	if err != nil {
		panic(err.Error())
	}
}
