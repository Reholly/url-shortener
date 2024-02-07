package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"urlshortener/internal/mapping"
	"urlshortener/internal/storage"
)

const (
	aliasQueryParam = "alias"
)

type GetUrlByAliasHandler struct {
	UrlRepo storage.UrlRepositoryContract
}

func NewGetUrlByAliasHandler(repo storage.UrlRepositoryContract) *GetUrlByAliasHandler {
	return &GetUrlByAliasHandler{UrlRepo: repo}
}

func (h *GetUrlByAliasHandler) GetUrlByAlias(c echo.Context) error {
	alias := c.QueryParam(aliasQueryParam)

	url, err := h.UrlRepo.GetByAlias(alias)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, mapping.UrlToUrlWithAliasDto(url))
}
