package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"urlshortener/internal/storage"
)

type RedirectOnUrlHandler struct {
	UrlRepo storage.UrlRepositoryContract
}

func NewRedirectOnUrlHandler(repo storage.UrlRepositoryContract) *RedirectOnUrlHandler {
	return &RedirectOnUrlHandler{UrlRepo: repo}
}

func (h *RedirectOnUrlHandler) RedirectOnUrl(c echo.Context) error {
	alias := c.QueryParam("alias")

	url, err := h.UrlRepo.GetByAlias(alias)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.Redirect(http.StatusPermanentRedirect, url.Url)
}
