package app

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	if err := godotenv.Load(); err != nil {
		panic(err.Error())
	}
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
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           "method=${method}, uri=${uri}, status=${status}\n",
		CustomTimeFormat: "15:00:00 01-01-2024",
	}))

	e.POST("/add", handlers.AddUrl(urlRepo))
	e.GET("/get", handlers.GetUrlByAlias(urlRepo))
	e.GET("/:alias", handlers.RedirectOnUrl(urlRepo))
	e.POST("/remove", handlers.RemoveUrl(urlRepo))

	err := e.Start(a.Config.Address)

	if err != nil {
		panic(err.Error())
	}
}
