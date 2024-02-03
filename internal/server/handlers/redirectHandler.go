package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"urlshortener/internal/storage"
)

type redirectOnUrlRequest struct {
	Alias string `json:"alias" bson:"alias"`
}

type RedirectOnUrlHandler struct {
	UrlRepo storage.UrlRepositoryContract
}

func NewRedirectOnUrlHandler(repo storage.UrlRepositoryContract) *RedirectOnUrlHandler {
	return &RedirectOnUrlHandler{UrlRepo: repo}
}

func (h *RedirectOnUrlHandler) RedirectOnUrl(c echo.Context) error {
	req := new(redirectOnUrlRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	url, err := h.UrlRepo.GetByAlias(req.Alias)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.Redirect(http.StatusPermanentRedirect, url.Url)
}
