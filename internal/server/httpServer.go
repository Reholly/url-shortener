package server

import (
	"github.com/labstack/echo/v4"
	"urlshortener/internal/server/handlers"
	"urlshortener/internal/storage"
)

type Server struct {
	urlRepository storage.UrlRepositoryContract
}

func New(repo storage.UrlRepositoryContract) *Server {
	return &Server{urlRepository: repo}
}

func (s *Server) RunServer(address string) {
	e := echo.New()

	addUrlHandler := handlers.NewAddUrlHandler(s.urlRepository)
	getByAliasHandler := handlers.NewGetUrlByAliasHandler(s.urlRepository)
	redirectHandler := handlers.NewRedirectOnUrlHandler(s.urlRepository)
	removeUrlHandler := handlers.NewRemoveUrlHandler(s.urlRepository)
	generateAgainHandler := handlers.NewGenerateAgainHandler(s.urlRepository)

	e.POST("/add", addUrlHandler.AddUrl)
	e.GET("/getByAlias", getByAliasHandler.GetUrlByAlias)
	e.GET("/", redirectHandler.RedirectOnUrl)
	e.POST("/remove", removeUrlHandler.RemoveUrl)
	e.POST("/generateAgain", generateAgainHandler.GenerateAgain)

	err := e.Start(address)

	if err != nil {
		panic(err.Error())
	}
}
