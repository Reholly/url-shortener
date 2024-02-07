package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"urlshortener/internal/domain"
	"urlshortener/internal/server/dto"
	"urlshortener/internal/storage"
	"urlshortener/internal/utils"
)

type AddUrlHandler struct {
	UrlRepo storage.UrlRepositoryContract
}

func NewAddUrlHandler(repo storage.UrlRepositoryContract) *AddUrlHandler {
	return &AddUrlHandler{UrlRepo: repo}
}

func (h *AddUrlHandler) AddUrl(c echo.Context) error {
	req := new(dto.UrlDto)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	alias := utils.GenerateUrlAlias()

	if err := h.UrlRepo.Add(domain.Url{Alias: alias, Url: req.Url}); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, nil)
}
