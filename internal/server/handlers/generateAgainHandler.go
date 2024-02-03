package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"urlshortener/internal/storage"
)

type generateAgainRequest struct {
	Url string `json:"url" bson:"url"`
}

type GenerateAgainHandler struct {
	UrlRepo storage.UrlRepositoryContract
}

func NewGenerateAgainHandler(repo storage.UrlRepositoryContract) *GenerateAgainHandler {
	return &GenerateAgainHandler{UrlRepo: repo}
}

func (h *GenerateAgainHandler) GenerateAgain(c echo.Context) error {
	req := new(generateAgainRequest)

	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	url, err := h.UrlRepo.GetByUrl(req.Url)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	addUrlHandler := NewAddUrlHandler(h.UrlRepo)
	url.Alias = addUrlHandler.CreateUrlAlias()

	err = h.UrlRepo.Update(url)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, url)
}
